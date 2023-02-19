package dpfm_api_processing_formatter

import (
	"context"
	"convert-to-dpfm-delivery-document-from-delivery-notice-edi-for-smes/DPFM_API_Caller/requests"
	dpfm_api_input_reader "convert-to-dpfm-delivery-document-from-delivery-notice-edi-for-smes/DPFM_API_Input_Reader"
	"sync"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	database "github.com/latonaio/golang-mysql-network-connector"
	"golang.org/x/xerrors"
)

type ProcessingFormatter struct {
	ctx context.Context
	db  *database.Mysql
	l   *logger.Logger
}

func NewProcessingFormatter(ctx context.Context, db *database.Mysql, l *logger.Logger) *ProcessingFormatter {
	return &ProcessingFormatter{
		ctx: ctx,
		db:  db,
		l:   l,
	}
}

func (p *ProcessingFormatter) ProcessingFormatter(
	sdc *dpfm_api_input_reader.SDC,
	psdc *ProcessingFormatterSDC,
) error {
	var err error
	var e error

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()

		psdc.Header, e = p.Header(sdc, psdc)
		if e != nil {
			err = e
			return
		}
		psdc.ConversionProcessingHeader, e = p.ConversionProcessingHeader(sdc, psdc)
		if e != nil {
			err = e
			return
		}
		psdc.Item, e = p.Item(sdc, psdc)
		if e != nil {
			err = e
			return
		}
		psdc.ConversionProcessingItem, e = p.ConversionProcessingItem(sdc, psdc)
		if e != nil {
			err = e
			return
		}
		psdc.Address, e = p.Address(sdc, psdc)
		if e != nil {
			err = e
			return
		}
		psdc.Partner, e = p.Partner(sdc, psdc)
		if e != nil {
			err = e
			return
		}

	}(&wg)

	wg.Wait()
	if err != nil {
		return err
	}

	p.l.Info(psdc)

	return nil
}

func (p *ProcessingFormatter) Header(sdc *dpfm_api_input_reader.SDC, psdc *ProcessingFormatterSDC) (*Header, error) {
	data := sdc.Header
	dataItem := sdc.Header.Item

	systemDate := getSystemDatePtr()
	systemTime := getSystemTimePtr()

	res := Header{
		ConvertingDeliveryDocument:      data.ExchangedDeliveryNoticeDocumentIdentifier,
		ConvertingBuyer:                 data.TradeBuyerIdentifier,
		ConvertingSeller:                data.TradeSellerIdentifier,
		ConvertingDeliverToParty:        data.TradeShipToPartyIdentifier,
		ConvertingDeliverFromParty:      data.TradeShipFromPartyIdentifier,
		DeliverFromPlant:                dataItem[0].LogisticsLocationIdentification,
		ConvertingReferenceDocument:     data.ReferencedOrdersDocumentIssureAssignedIdentifier,
		ConvertingReferenceDocumentItem: dataItem[0].ReferencedOrdersDocumentItemLineIdentifier,
		DocumentDate:                    data.ExchangedDeliveryNoticeDocumentIssueDate,
		PlannedGoodsReceiptDate:         dataItem[0].SupplyChainEventRequirementOccurrenceDate,
		CreationDate:                    systemDate,
		CreationTime:                    systemTime,
		LastChangeDate:                  systemDate,
		LastChangeTime:                  systemTime,
		GoodsIssueOrReceiptSlipNumber:   dataItem[0].SupplyChainDeliveryEventIdentifier,
		HeaderDeliveryBlockStatus:       getBoolPtr(false),
		HeaderIssuingBlockStatus:        getBoolPtr(false),
		HeaderReceivingBlockStatus:      getBoolPtr(false),
		IsCancelled:                     getBoolPtr(false),
		IsMarkedForDeletion:             getBoolPtr(false),
	}

	return &res, nil
}

func (p *ProcessingFormatter) ConversionProcessingHeader(sdc *dpfm_api_input_reader.SDC, psdc *ProcessingFormatterSDC) (*ConversionProcessingHeader, error) {
	dataKey := make([]*ConversionProcessingKey, 0)

	dataKey = append(dataKey, p.ConversionProcessingKey(sdc, "ExchangedDeliveryNoticeDocumentIdentifier", "DeliveryDocument", psdc.Header.ConvertingDeliveryDocument))
	dataKey = append(dataKey, p.ConversionProcessingKey(sdc, "TradeBuyerIdentifier", "Buyer", psdc.Header.ConvertingBuyer))
	dataKey = append(dataKey, p.ConversionProcessingKey(sdc, "TradeSellerIdentifier", "Seller", psdc.Header.ConvertingSeller))
	dataKey = append(dataKey, p.ConversionProcessingKey(sdc, "TradeShipToPartyIdentifier", "DeliverToParty", psdc.Header.ConvertingDeliverToParty))
	dataKey = append(dataKey, p.ConversionProcessingKey(sdc, "TradeShipFromPartyIdentifier", "DeliverFromParty", psdc.Header.ConvertingDeliverFromParty))
	dataKey = append(dataKey, p.ConversionProcessingKey(sdc, "ReferencedOrdersDocumentIssureAssignedIdentifier", "ReferenceDocument", psdc.Header.ConvertingReferenceDocument))
	dataKey = append(dataKey, p.ConversionProcessingKey(sdc, "ReferencedOrdersDocumentItemLineIdentifier", "ReferenceDocumentItem", psdc.Header.ConvertingReferenceDocumentItem))

	dataQueryGets, err := p.ConversionProcessingCommonQueryGets(dataKey)
	if err != nil {
		return nil, xerrors.Errorf("ConversionProcessing Error: %w", err)
	}

	data, err := p.ConvertToConversionProcessingHeader(dataKey, dataQueryGets)
	if err != nil {
		return nil, xerrors.Errorf("ConvertToConversionProcessing Error: %w", err)
	}

	return data, nil
}

func (psdc *ProcessingFormatter) ConvertToConversionProcessingHeader(conversionProcessingKey []*ConversionProcessingKey, conversionProcessingCommonQueryGets []*ConversionProcessingCommonQueryGets) (*ConversionProcessingHeader, error) {
	data := make(map[string]*ConversionProcessingCommonQueryGets, len(conversionProcessingCommonQueryGets))
	for _, v := range conversionProcessingCommonQueryGets {
		data[v.LabelConvertTo] = v
	}

	for _, v := range conversionProcessingKey {
		if _, ok := data[v.LabelConvertTo]; !ok {
			return nil, xerrors.Errorf("%s is not in the database", v.LabelConvertTo)
		}
	}

	pm := &requests.ConversionProcessingHeader{}

	pm.ConvertingDeliveryDocument = data["DeliveryDocument"].CodeConvertFromString
	pm.ConvertedDeliveryDocument = data["DeliveryDocument"].CodeConvertToInt
	pm.ConvertingBuyer = data["Buyer"].CodeConvertFromString
	pm.ConvertedBuyer = data["Buyer"].CodeConvertToInt
	pm.ConvertingSeller = data["Seller"].CodeConvertFromString
	pm.ConvertedSeller = data["Seller"].CodeConvertToInt
	pm.ConvertingDeliverToParty = data["DeliverToParty"].CodeConvertFromString
	pm.ConvertedDeliverToParty = data["DeliverToParty"].CodeConvertToInt
	pm.ConvertingDeliverFromParty = data["DeliverFromParty"].CodeConvertFromString
	pm.ConvertedDeliverFromParty = data["DeliverFromParty"].CodeConvertToInt
	pm.ConvertingReferenceDocument = data["ReferenceDocument"].CodeConvertFromString
	pm.ConvertedReferenceDocument = data["ReferenceDocument"].CodeConvertToInt
	pm.ConvertingReferenceDocumentItem = data["ReferenceDocumentItem"].CodeConvertFromString
	pm.ConvertedReferenceDocumentItem = data["ReferenceDocumentItem"].CodeConvertToInt

	res := &ConversionProcessingHeader{
		ConvertingDeliveryDocument:      pm.ConvertingDeliveryDocument,
		ConvertedDeliveryDocument:       pm.ConvertedDeliveryDocument,
		ConvertingBuyer:                 pm.ConvertingBuyer,
		ConvertedBuyer:                  pm.ConvertedBuyer,
		ConvertingSeller:                pm.ConvertingSeller,
		ConvertedSeller:                 pm.ConvertedSeller,
		ConvertingDeliverToParty:        pm.ConvertingDeliverToParty,
		ConvertedDeliverToParty:         pm.ConvertedDeliverToParty,
		ConvertingDeliverFromParty:      pm.ConvertingDeliverFromParty,
		ConvertedDeliverFromParty:       pm.ConvertedDeliverFromParty,
		ConvertingReferenceDocument:     pm.ConvertingReferenceDocument,
		ConvertedReferenceDocument:      pm.ConvertedReferenceDocument,
		ConvertingReferenceDocumentItem: pm.ConvertingReferenceDocumentItem,
		ConvertedReferenceDocumentItem:  pm.ConvertedReferenceDocumentItem,
	}

	return res, nil
}

func (p *ProcessingFormatter) Item(sdc *dpfm_api_input_reader.SDC, psdc *ProcessingFormatterSDC) ([]*Item, error) {
	res := make([]*Item, 0)
	dataHeader := psdc.Header
	data := sdc.Header.Item

	systemDate := getSystemDatePtr()
	systemTime := getSystemTimePtr()

	for _, data := range data {

		res = append(res, &Item{
			ConvertingDeliveryDocument:             dataHeader.ConvertingDeliveryDocument,
			ConvertingDeliveryDocumentItem:         data.DeliveryNoticeDocumentItemlineIdentifier,
			ConvertingBuyer:                        dataHeader.ConvertingBuyer,
			ConvertingSeller:                       dataHeader.ConvertingSeller,
			ConvertingDeliverToParty:               dataHeader.ConvertingDeliverToParty,
			ConvertingDeliverFromParty:             dataHeader.ConvertingDeliverFromParty,
			DeliverFromPlant:                       dataHeader.DeliverFromPlant,
			DeliveryDocumentItemText:               data.NoteDeliveryNoticeItemContentText,
			Product:                                data.TradeProductIdentifier,
			ProductStandardID:                      data.TradeProductGlobalIdentifier,
			DeliveryUnit:                           data.ReferencedLogisticsPackageQuantityUnitCode,
			CreationDate:                           systemDate,
			CreationTime:                           systemTime,
			LastChangeDate:                         systemDate,
			LastChangeTime:                         systemTime,
			NetAmount:                              data.ItemTradeDeliveryNoticeSettlementMonetarySummationNetTotalAmount,
			GrossAmount:                            data.ItemTradeDeliveryNoticeSettlementMonetarySummationIncludingTaxesNetTotalAmount,
			ConvertingOrderID:                      dataHeader.ConvertingReferenceDocument,
			ConvertingOrderItem:                    dataHeader.ConvertingReferenceDocument,
			ConvertingProject:                      sdc.Header.ProjectIdentifier,
			ConvertingReferenceDocument:            dataHeader.ConvertingReferenceDocument,
			ConvertingReferenceDocumentItem:        dataHeader.ConvertingReferenceDocumentItem,
			ConvertingTransactionTaxClassification: data.ItemTradeTaxCategoryCode,
			ItemDeliveryBlockStatus:                getBoolPtr(false),
			ItemIssuingBlockStatus:                 getBoolPtr(false),
			ItemReceivingBlockStatus:               getBoolPtr(false),
			ItemBillingBlockStatus:                 getBoolPtr(false),
			IsCancelled:                            getBoolPtr(false),
			IsMarkedForDeletion:                    getBoolPtr(false),
		})
	}

	return res, nil
}

func (p *ProcessingFormatter) ConversionProcessingItem(sdc *dpfm_api_input_reader.SDC, psdc *ProcessingFormatterSDC) ([]*ConversionProcessingItem, error) {
	data := make([]*ConversionProcessingItem, 0)

	for _, item := range psdc.Item {
		dataKey := make([]*ConversionProcessingKey, 0)

		dataKey = append(dataKey, p.ConversionProcessingKey(sdc, "DeliveryNoticeDocumentItemlineIdentifier", "DeliveryDocumentItem", item.ConvertingDeliveryDocumentItem))
		dataKey = append(dataKey, p.ConversionProcessingKey(sdc, "ProjectIdentifier", "Project", item.ConvertingProject))
		dataKey = append(dataKey, p.ConversionProcessingKey(sdc, "ItemTradeTaxCategoryCode", "TransactionTaxClassification", item.ConvertingTransactionTaxClassification))

		dataQueryGets, err := p.ConversionProcessingCommonQueryGets(dataKey)
		if err != nil {
			return nil, xerrors.Errorf("ConversionProcessing Error: %w", err)
		}

		datum, err := p.ConvertToConversionProcessingItem(dataKey, dataQueryGets)
		if err != nil {
			return nil, xerrors.Errorf("ConvertToConversionProcessing Error: %w", err)
		}

		data = append(data, datum)
	}

	return data, nil
}

func (p *ProcessingFormatter) ConvertToConversionProcessingItem(conversionProcessingKey []*ConversionProcessingKey, conversionProcessingCommonQueryGets []*ConversionProcessingCommonQueryGets) (*ConversionProcessingItem, error) {
	data := make(map[string]*ConversionProcessingCommonQueryGets, len(conversionProcessingCommonQueryGets))
	for _, v := range conversionProcessingCommonQueryGets {
		data[v.LabelConvertTo] = v
	}

	for _, v := range conversionProcessingKey {
		if _, ok := data[v.LabelConvertTo]; !ok {
			return nil, xerrors.Errorf("%s is not in the database", v.LabelConvertTo)
		}
	}

	pm := &requests.ConversionProcessingItem{}

	pm.ConvertingDeliveryDocumentItem = data["DeliveryDocumentItem"].CodeConvertFromString
	pm.ConvertedDeliveryDocumentItem = data["DeliveryDocumentItem"].CodeConvertToInt
	pm.ConvertingProject = data["Project"].CodeConvertFromString
	pm.ConvertedProject = data["Project"].CodeConvertFromString
	pm.ConvertingTransactionTaxClassification = data["TransactionTaxClassification"].CodeConvertFromString
	pm.ConvertedTransactionTaxClassification = data["TransactionTaxClassification"].CodeConvertFromString

	res := &ConversionProcessingItem{
		ConvertingDeliveryDocumentItem:         pm.ConvertingDeliveryDocumentItem,
		ConvertedDeliveryDocumentItem:          pm.ConvertedDeliveryDocumentItem,
		ConvertingProject:                      pm.ConvertingProject,
		ConvertedProject:                       pm.ConvertedProject,
		ConvertingTransactionTaxClassification: pm.ConvertingTransactionTaxClassification,
		ConvertedTransactionTaxClassification:  pm.ConvertedTransactionTaxClassification,
	}

	return res, nil
}

func (p *ProcessingFormatter) Address(sdc *dpfm_api_input_reader.SDC, psdc *ProcessingFormatterSDC) ([]*Address, error) {
	res := make([]*Address, 0)
	dataHeader := psdc.Header

	res = append(res, &Address{
		ConvertingDeliveryDocument: dataHeader.ConvertingDeliveryDocument,
		PostalCode:                 sdc.Header.SellerAddressPostalCode,
	})

	return res, nil
}

func (p *ProcessingFormatter) Partner(sdc *dpfm_api_input_reader.SDC, psdc *ProcessingFormatterSDC) ([]*Partner, error) {
	res := make([]*Partner, 0)
	dataHeader := psdc.Header

	for range psdc.Partner {
		res = append(res, &Partner{
			ConvertingDeliveryDocument: dataHeader.ConvertingDeliveryDocument,
		})
	}
	return res, nil
}
