package dpfm_api_processing_formatter

import (
	"context"
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

	if bpIDIsNull(sdc) {
		return xerrors.New("business_partner is null")
	}

	wg := sync.WaitGroup{}

	psdc.Header = p.Header(sdc, psdc)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		// Ref: Header
		psdc.ConversionProcessingHeader, e = p.ConversionProcessingHeader(sdc, psdc)
		if e != nil {
			err = e
			return
		}
	}(&wg)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		// Ref: Header
		psdc.Item = p.Item(sdc, psdc)

		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			// Ref: Item
			psdc.ConversionProcessingItem, e = p.ConversionProcessingItem(sdc, psdc)
			if e != nil {
				err = e
				return
			}
		}(wg)
	}(&wg)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		// Ref: Header
		psdc.Address = p.Address(sdc, psdc)
	}(&wg)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		// Ref: Header
		psdc.Partner = p.Partner(sdc, psdc)
	}(&wg)

	wg.Wait()
	if err != nil {
		return err
	}

	p.l.Info(psdc)

	return nil
}

func (p *ProcessingFormatter) Header(sdc *dpfm_api_input_reader.SDC, psdc *ProcessingFormatterSDC) *Header {
	data := sdc.Header
	dataItem := sdc.Header.Item[0]

	systemDate := getSystemDatePtr()
	systemTime := getSystemTimePtr()

	res := Header{
		ConvertingDeliveryDocument:      data.ExchangedDeliveryNoticeDocumentIdentifier,
		ConvertingBuyer:                 data.TradeBuyerIdentifier,
		ConvertingSeller:                data.TradeSellerIdentifier,
		ConvertingDeliverToParty:        data.TradeShipToPartyIdentifier,
		ConvertingDeliverFromParty:      data.TradeShipFromPartyIdentifier,
		DeliverFromPlant:                dataItem.LogisticsLocationIdentification,
		ConvertingBillToParty:           data.TradeBuyerIdentifier,
		ConvertingBillFromParty:         data.TradeSellerIdentifier,
		ConvertingPayer:                 data.TradeBuyerIdentifier,
		ConvertingPayee:                 data.TradeSellerIdentifier,
		ConvertingReferenceDocument:     data.ReferencedOrdersDocumentIssureAssignedIdentifier,
		ConvertingReferenceDocumentItem: dataItem.ReferencedOrdersDocumentItemLineIdentifier,
		ConvertingOrderID:               data.ReferencedOrdersDocumentIssureAssignedIdentifier,
		ConvertingOrderItem:             dataItem.ReferencedOrdersDocumentItemLineIdentifier,
		DocumentDate:                    data.ExchangedDeliveryNoticeDocumentIssueDate,
		PlannedGoodsIssueDate:           systemDate,
		PlannedGoodsReceiptDate:         dataItem.SupplyChainEventRequirementOccurrenceDate,
		CreationDate:                    systemDate,
		CreationTime:                    systemTime,
		LastChangeDate:                  systemDate,
		LastChangeTime:                  systemTime,
		GoodsIssueOrReceiptSlipNumber:   dataItem.SupplyChainDeliveryEventIdentifier,
		HeaderDeliveryBlockStatus:       getBoolPtr(false),
		HeaderIssuingBlockStatus:        getBoolPtr(false),
		HeaderReceivingBlockStatus:      getBoolPtr(false),
		IsCancelled:                     getBoolPtr(false),
		IsMarkedForDeletion:             getBoolPtr(false),
		ConvertingProject:               data.ProjectIdentifier,
	}

	return &res
}

func (p *ProcessingFormatter) ConversionProcessingHeader(sdc *dpfm_api_input_reader.SDC, psdc *ProcessingFormatterSDC) (*ConversionProcessingHeader, error) {
	dataKey := make([]*ConversionProcessingKey, 0)

	p.appendDataKey(&dataKey, sdc, "ExchangedDeliveryNoticeDocumentIdentifier", "DeliveryDocument", psdc.Header.ConvertingDeliveryDocument)
	p.appendDataKey(&dataKey, sdc, "TradeBuyerIdentifier", "Buyer", psdc.Header.ConvertingBuyer)
	p.appendDataKey(&dataKey, sdc, "TradeSellerIdentifier", "Seller", psdc.Header.ConvertingSeller)
	p.appendDataKey(&dataKey, sdc, "TradeShipToPartyIdentifier", "DeliverToParty", psdc.Header.ConvertingDeliverToParty)
	p.appendDataKey(&dataKey, sdc, "TradeShipFromPartyIdentifier", "DeliverFromParty", psdc.Header.ConvertingDeliverFromParty)
	p.appendDataKey(&dataKey, sdc, "TradeBuyerIdentifier", "BillToParty", psdc.Header.ConvertingBillToParty)
	p.appendDataKey(&dataKey, sdc, "TradeSellerIdentifier", "BillFromParty", psdc.Header.ConvertingBillFromParty)
	p.appendDataKey(&dataKey, sdc, "TradeBuyerIdentifier", "Payer", psdc.Header.ConvertingPayer)
	p.appendDataKey(&dataKey, sdc, "TradeSellerIdentifier", "Payee", psdc.Header.ConvertingPayee)
	p.appendDataKey(&dataKey, sdc, "ReferencedOrdersDocumentIssureAssignedIdentifier", "ReferenceDocument", psdc.Header.ConvertingReferenceDocument)
	p.appendDataKey(&dataKey, sdc, "ReferencedOrdersDocumentItemLineIdentifier", "ReferenceDocumentItem", psdc.Header.ConvertingReferenceDocumentItem)

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
			return nil, xerrors.Errorf("Value of %s is not in the database", v.LabelConvertTo)
		}
	}

	res := &ConversionProcessingHeader{}

	if _, ok := data["DeliveryDocument"]; ok {
		res.ConvertingDeliveryDocument = data["DeliveryDocument"].CodeConvertFromString
		res.ConvertedDeliveryDocument = data["DeliveryDocument"].CodeConvertToInt
	}
	if _, ok := data["Buyer"]; ok {
		res.ConvertingBuyer = data["Buyer"].CodeConvertFromString
		res.ConvertedBuyer = data["Buyer"].CodeConvertToInt
	}
	if _, ok := data["Seller"]; ok {
		res.ConvertingSeller = data["Seller"].CodeConvertFromString
		res.ConvertedSeller = data["Seller"].CodeConvertToInt
	}
	if _, ok := data["DeliverToParty"]; ok {
		res.ConvertingDeliverToParty = data["DeliverToParty"].CodeConvertFromString
		res.ConvertedDeliverToParty = data["DeliverToParty"].CodeConvertToInt
	}
	if _, ok := data["DeliverFromParty"]; ok {
		res.ConvertingDeliverFromParty = data["DeliverFromParty"].CodeConvertFromString
		res.ConvertedDeliverFromParty = data["DeliverFromParty"].CodeConvertToInt
	}
	if _, ok := data["BillToParty"]; ok {
		res.ConvertingBillToParty = data["BillToParty"].CodeConvertFromString
		res.ConvertedBillToParty = data["BillToParty"].CodeConvertToInt
	}
	if _, ok := data["BillFromParty"]; ok {
		res.ConvertingBillFromParty = data["BillFromParty"].CodeConvertFromString
		res.ConvertedBillFromParty = data["BillFromParty"].CodeConvertToInt
	}
	if _, ok := data["Payer"]; ok {
		res.ConvertingPayer = data["Payer"].CodeConvertFromString
		res.ConvertedPayer = data["Payer"].CodeConvertToInt
	}
	if _, ok := data["Payee"]; ok {
		res.ConvertingPayee = data["Payee"].CodeConvertFromString
		res.ConvertedPayee = data["Payee"].CodeConvertToInt
	}
	if _, ok := data["ReferenceDocument"]; ok {
		res.ConvertingReferenceDocument = data["ReferenceDocument"].CodeConvertFromString
		res.ConvertedReferenceDocument = data["ReferenceDocument"].CodeConvertToInt
	}
	if _, ok := data["ReferenceDocumentItem"]; ok {
		res.ConvertingReferenceDocumentItem = data["ReferenceDocumentItem"].CodeConvertFromString
		res.ConvertedReferenceDocumentItem = data["ReferenceDocumentItem"].CodeConvertToInt
	}

	return res, nil
}

func (p *ProcessingFormatter) Item(sdc *dpfm_api_input_reader.SDC, psdc *ProcessingFormatterSDC) []*Item {
	res := make([]*Item, 0)
	dataHeader := psdc.Header
	data := sdc.Header.Item

	systemDate := getSystemDatePtr()
	systemTime := getSystemTimePtr()

	for _, data := range data {

		res = append(res, &Item{
			ConvertingDeliveryDocument:                 dataHeader.ConvertingDeliveryDocument,
			ConvertingDeliveryDocumentItem:             data.DeliveryNoticeDocumentItemlineIdentifier,
			ConvertingBuyer:                            dataHeader.ConvertingBuyer,
			ConvertingSeller:                           dataHeader.ConvertingSeller,
			ConvertingDeliverToParty:                   dataHeader.ConvertingDeliverToParty,
			ConvertingDeliverFromParty:                 dataHeader.ConvertingDeliverFromParty,
			DeliverFromPlant:                           dataHeader.DeliverFromPlant,
			ConvertingBillToParty:                      dataHeader.ConvertingBillToParty,
			ConvertingBillFromParty:                    dataHeader.ConvertingBillFromParty,
			ConvertingPayer:                            dataHeader.ConvertingPayer,
			ConvertingPayee:                            dataHeader.ConvertingPayee,
			ConvertingStockConfirmationBusinessPartner: dataHeader.ConvertingDeliverFromParty,
			ConvertingStockConfirmationPlant:           dataHeader.DeliverFromPlant,
			DeliveryDocumentItemText:                   data.NoteDeliveryNoticeItemContentText,
			Product:                                    data.TradeProductIdentifier,
			ProductStandardID:                          data.TradeProductGlobalIdentifier,
			DeliveryUnit:                               data.ReferencedLogisticsPackageQuantityUnitCode,
			CreationDate:                               systemDate,
			CreationTime:                               systemTime,
			LastChangeDate:                             systemDate,
			LastChangeTime:                             systemTime,
			NetAmount:                                  data.ItemTradeDeliveryNoticeSettlementMonetarySummationNetTotalAmount,
			GrossAmount:                                data.ItemTradeDeliveryNoticeSettlementMonetarySummationIncludingTaxesNetTotalAmount,
			ConvertingOrderID:                          dataHeader.ConvertingReferenceDocument,
			ConvertingOrderItem:                        dataHeader.ConvertingReferenceDocument,
			ConvertingProject:                          dataHeader.ConvertingProject,
			ConvertingReferenceDocument:                dataHeader.ConvertingReferenceDocument,
			ConvertingReferenceDocumentItem:            dataHeader.ConvertingReferenceDocumentItem,
			ConvertingTransactionTaxClassification:     data.ItemTradeTaxCategoryCode,
			ItemDeliveryBlockStatus:                    getBoolPtr(false),
			ItemIssuingBlockStatus:                     getBoolPtr(false),
			ItemReceivingBlockStatus:                   getBoolPtr(false),
			ItemBillingBlockStatus:                     getBoolPtr(false),
			IsCancelled:                                getBoolPtr(false),
			IsMarkedForDeletion:                        getBoolPtr(false),
		})
	}

	return res
}

func (p *ProcessingFormatter) ConversionProcessingItem(sdc *dpfm_api_input_reader.SDC, psdc *ProcessingFormatterSDC) ([]*ConversionProcessingItem, error) {
	data := make([]*ConversionProcessingItem, 0)

	for _, item := range psdc.Item {
		dataKey := make([]*ConversionProcessingKey, 0)

		p.appendDataKey(&dataKey, sdc, "DeliveryNoticeDocumentItemlineIdentifier", "DeliveryDocumentItem", item.ConvertingDeliveryDocumentItem)
		p.appendDataKey(&dataKey, sdc, "ProjectIdentifier", "Project", item.ConvertingProject)
		p.appendDataKey(&dataKey, sdc, "ItemTradeTaxCategoryCode", "TransactionTaxClassification", item.ConvertingTransactionTaxClassification)

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

	res := &ConversionProcessingItem{}

	if _, ok := data["DeliveryDocumentItem"]; ok {
		res.ConvertingDeliveryDocumentItem = data["DeliveryDocumentItem"].CodeConvertFromString
		res.ConvertedDeliveryDocumentItem = data["DeliveryDocumentItem"].CodeConvertToInt
	}
	if _, ok := data["Project"]; ok {
		res.ConvertingProject = data["Project"].CodeConvertFromString
		res.ConvertedProject = data["Project"].CodeConvertFromString
	}
	if _, ok := data["TransactionTaxClassification"]; ok {
		res.ConvertingTransactionTaxClassification = data["TransactionTaxClassification"].CodeConvertFromString
		res.ConvertedTransactionTaxClassification = data["TransactionTaxClassification"].CodeConvertFromString
	}

	return res, nil
}

func (p *ProcessingFormatter) Address(sdc *dpfm_api_input_reader.SDC, psdc *ProcessingFormatterSDC) []*Address {
	res := make([]*Address, 0)

	deliverToPartyAddress := deliverToPartyAddress(sdc, psdc)
	if !postalCodeContains(deliverToPartyAddress.PostalCode, res) {
		res = append(res, deliverToPartyAddress)
	}

	deliverFromPartyAddress := deliverFromPartyAddress(sdc, psdc)
	if !postalCodeContains(deliverFromPartyAddress.PostalCode, res) {
		res = append(res, deliverFromPartyAddress)
	}

	buyerAddress := buyerAddress(sdc, psdc)
	if !postalCodeContains(buyerAddress.PostalCode, res) {
		res = append(res, buyerAddress)
	}

	sellerAddress := sellerAddress(sdc, psdc)
	if !postalCodeContains(sellerAddress.PostalCode, res) {
		res = append(res, sellerAddress)
	}

	return res
}

func deliverToPartyAddress(sdc *dpfm_api_input_reader.SDC, psdc *ProcessingFormatterSDC) *Address {
	dataHeader := psdc.Header

	res := &Address{
		ConvertingDeliveryDocument: dataHeader.ConvertingDeliveryDocument,
		PostalCode:                 sdc.Header.SellerAddressPostalCode,
	}

	return res
}

func deliverFromPartyAddress(sdc *dpfm_api_input_reader.SDC, psdc *ProcessingFormatterSDC) *Address {
	dataHeader := psdc.Header

	res := &Address{
		ConvertingDeliveryDocument: dataHeader.ConvertingDeliveryDocument,
		PostalCode:                 sdc.Header.SellerAddressPostalCode,
	}

	return res
}

func buyerAddress(sdc *dpfm_api_input_reader.SDC, psdc *ProcessingFormatterSDC) *Address {
	dataHeader := psdc.Header

	res := &Address{
		ConvertingDeliveryDocument: dataHeader.ConvertingDeliveryDocument,
		PostalCode:                 sdc.Header.BuyerAddressPostalCode,
	}

	return res
}

func sellerAddress(sdc *dpfm_api_input_reader.SDC, psdc *ProcessingFormatterSDC) *Address {
	dataHeader := psdc.Header

	res := &Address{
		ConvertingDeliveryDocument: dataHeader.ConvertingDeliveryDocument,
		PostalCode:                 sdc.Header.SellerAddressPostalCode,
	}

	return res
}

func (p *ProcessingFormatter) Partner(sdc *dpfm_api_input_reader.SDC, psdc *ProcessingFormatterSDC) []*Partner {
	res := make([]*Partner, 0)
	dataHeader := psdc.Header

	res = append(res, &Partner{
		ConvertingDeliveryDocument: dataHeader.ConvertingDeliveryDocument,
	})

	return res
}

func (p *ProcessingFormatter) appendDataKey(dataKey *[]*ConversionProcessingKey, sdc *dpfm_api_input_reader.SDC, labelConvertFrom string, labelConvertTo string, codeConvertFrom any) {
	switch v := codeConvertFrom.(type) {
	case int, float32:
		if v == 0 {
			return
		}
	case string:
		if v == "" {
			return
		}
	case *int, *float32:
		if v == nil {
			return
		}
	case *string:
		if v == nil || *v == "" {
			return
		}
	default:
		return
	}
	*dataKey = append(*dataKey, p.ConversionProcessingKey(sdc, labelConvertFrom, labelConvertTo, codeConvertFrom))
}

func postalCodeContains(postalCode *string, addresses []*Address) bool {
	for _, address := range addresses {
		if address.PostalCode == nil || postalCode == nil {
			return true
		}
		if *address.PostalCode == *postalCode {
			return true
		}
	}

	return false
}

func bpIDIsNull(sdc *dpfm_api_input_reader.SDC) bool {
	return sdc.BusinessPartnerID == nil
}
