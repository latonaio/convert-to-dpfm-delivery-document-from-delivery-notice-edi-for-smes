package dpfm_api_output_formatter

import (
	dpfm_api_input_reader "convert-to-dpfm-delivery-document-from-delivery-notice-edi-for-smes/DPFM_API_Input_Reader"
	dpfm_api_processing_formatter "convert-to-dpfm-delivery-document-from-delivery-notice-edi-for-smes/DPFM_API_Processing_Formatter"
)

func OutputFormatter(
	sdc *dpfm_api_input_reader.SDC,
	psdc *dpfm_api_processing_formatter.ProcessingFormatterSDC,
	osdc *Output,
) error {
	header := ConvertToHeader(*sdc, *psdc)
	item := ConvertToItem(*sdc, *psdc)
	address := ConvertToAddress(*sdc, *psdc)
	partner := ConvertToPartner(*sdc, *psdc)

	osdc.DataConcatenation = DataConcatenation{
		Header:  header,
		Item:    item,
		Address: address,
		Partner: partner,
	}

	osdc.ServiceLabel = "FUNCTION_DELIVERY_DOCUMENT_DATA_CONCATENATION"
	osdc.APISchema = "DPFMDataConcatenation"
	osdc.APIProcessingResult = getBoolPtr(true)

	return nil
}

func ConvertToHeader(
	sdc dpfm_api_input_reader.SDC,
	psdc dpfm_api_processing_formatter.ProcessingFormatterSDC,
) *Header {
	dataProcessingHeader := psdc.Header
	dataConversionProcessingHeader := psdc.ConversionProcessingHeader

	header := &Header{
		DeliveryDocument:              *dataConversionProcessingHeader.ConvertedDeliveryDocument,
		Buyer:                         dataConversionProcessingHeader.ConvertedBuyer,
		Seller:                        dataConversionProcessingHeader.ConvertedSeller,
		DeliverToParty:                dataConversionProcessingHeader.ConvertedDeliverToParty,
		DeliverFromParty:              dataConversionProcessingHeader.ConvertedDeliverFromParty,
		DeliverFromPlant:              dataProcessingHeader.DeliverFromPlant,
		BillToParty:                   dataConversionProcessingHeader.ConvertedBillToParty,
		BillFromParty:                 dataConversionProcessingHeader.ConvertedBillFromParty,
		Payer:                         dataConversionProcessingHeader.ConvertedPayer,
		Payee:                         dataConversionProcessingHeader.ConvertedPayee,
		ReferenceDocument:             dataConversionProcessingHeader.ConvertedReferenceDocument,
		ReferenceDocumentItem:         dataConversionProcessingHeader.ConvertedReferenceDocumentItem,
		OrderID:                       dataConversionProcessingHeader.ConvertedReferenceDocument,
		OrderItem:                     dataConversionProcessingHeader.ConvertedReferenceDocumentItem,
		DocumentDate:                  dataProcessingHeader.DocumentDate,
		PlannedGoodsIssueDate:         dataProcessingHeader.PlannedGoodsIssueDate,
		PlannedGoodsReceiptDate:       dataProcessingHeader.PlannedGoodsReceiptDate,
		CreationDate:                  dataProcessingHeader.CreationDate,
		CreationTime:                  dataProcessingHeader.CreationTime,
		LastChangeDate:                dataProcessingHeader.LastChangeDate,
		LastChangeTime:                dataProcessingHeader.LastChangeTime,
		GoodsIssueOrReceiptSlipNumber: dataProcessingHeader.GoodsIssueOrReceiptSlipNumber,
		HeaderDeliveryBlockStatus:     dataProcessingHeader.HeaderDeliveryBlockStatus,
		HeaderIssuingBlockStatus:      dataProcessingHeader.HeaderIssuingBlockStatus,
		HeaderReceivingBlockStatus:    dataProcessingHeader.HeaderReceivingBlockStatus,
		IsCancelled:                   dataProcessingHeader.IsCancelled,
		IsMarkedForDeletion:           dataProcessingHeader.IsMarkedForDeletion,
	}

	return header
}

func ConvertToItem(
	sdc dpfm_api_input_reader.SDC,
	psdc dpfm_api_processing_formatter.ProcessingFormatterSDC,
) []*Item {
	dataProcessingHeader := psdc.Header
	dataProcessingItem := psdc.Item
	dataConversionProcessingHeader := psdc.ConversionProcessingHeader
	dataConversionProcessingItem := psdc.ConversionProcessingItem

	items := make([]*Item, 0)
	for i := range dataProcessingItem {
		item := &Item{
			DeliveryDocument:                 *dataConversionProcessingHeader.ConvertedDeliveryDocument,
			DeliveryDocumentItem:             *dataConversionProcessingItem[i].ConvertedDeliveryDocumentItem,
			Buyer:                            dataConversionProcessingHeader.ConvertedBuyer,
			Seller:                           dataConversionProcessingHeader.ConvertedSeller,
			DeliverToParty:                   dataConversionProcessingHeader.ConvertedDeliverToParty,
			DeliverFromParty:                 dataConversionProcessingHeader.ConvertedDeliverFromParty,
			DeliverFromPlant:                 dataProcessingItem[i].DeliverFromPlant,
			BillToParty:                      dataConversionProcessingHeader.ConvertedBillToParty,
			BillFromParty:                    dataConversionProcessingHeader.ConvertedBillFromParty,
			Payer:                            dataConversionProcessingHeader.ConvertedPayer,
			Payee:                            dataConversionProcessingHeader.ConvertedPayee,
			StockConfirmationBusinessPartner: dataConversionProcessingHeader.ConvertedDeliverFromParty,
			StockConfirmationPlant:           dataProcessingHeader.DeliverFromPlant,
			DeliveryDocumentItemText:         dataProcessingItem[i].DeliveryDocumentItemText,
			Product:                          dataProcessingItem[i].Product,
			ProductStandardID:                dataProcessingItem[i].ProductStandardID,
			DeliveryUnit:                     dataProcessingItem[i].DeliveryUnit,
			CreationDate:                     dataProcessingItem[i].CreationDate,
			CreationTime:                     dataProcessingItem[i].CreationTime,
			LastChangeDate:                   dataProcessingItem[i].LastChangeDate,
			LastChangeTime:                   dataProcessingItem[i].LastChangeTime,
			NetAmount:                        dataProcessingItem[i].NetAmount,
			GrossAmount:                      dataProcessingItem[i].GrossAmount,
			OrderID:                          dataConversionProcessingHeader.ConvertedReferenceDocument,
			OrderItem:                        dataConversionProcessingHeader.ConvertedReferenceDocumentItem,
			Project:                          dataConversionProcessingItem[i].ConvertedProject,
			ReferenceDocument:                dataConversionProcessingHeader.ConvertedReferenceDocument,
			ReferenceDocumentItem:            dataConversionProcessingHeader.ConvertedReferenceDocumentItem,
			TransactionTaxClassification:     dataConversionProcessingItem[i].ConvertedTransactionTaxClassification,
			ItemDeliveryBlockStatus:          dataProcessingItem[i].ItemDeliveryBlockStatus,
			ItemIssuingBlockStatus:           dataProcessingItem[i].ItemIssuingBlockStatus,
			ItemReceivingBlockStatus:         dataProcessingItem[i].ItemReceivingBlockStatus,
			ItemBillingBlockStatus:           dataProcessingItem[i].ItemBillingBlockStatus,
			IsCancelled:                      dataProcessingItem[i].IsCancelled,
			IsMarkedForDeletion:              dataProcessingItem[i].IsMarkedForDeletion,
		}

		items = append(items, item)
	}

	return items
}

func ConvertToAddress(
	sdc dpfm_api_input_reader.SDC,
	psdc dpfm_api_processing_formatter.ProcessingFormatterSDC,
) []*Address {
	dataConversionProcessingHeader := psdc.ConversionProcessingHeader
	data := psdc.Address

	addresses := make([]*Address, 0)
	for _, data := range data {
		addresses = append(addresses, &Address{
			DeliveryDocument: *dataConversionProcessingHeader.ConvertedDeliveryDocument,
			PostalCode:       data.PostalCode,
		})
	}

	return addresses
}

func ConvertToPartner(
	sdc dpfm_api_input_reader.SDC,
	psdc dpfm_api_processing_formatter.ProcessingFormatterSDC,
) []*Partner {
	dataProcessingPartner := psdc.Partner
	dataConversionProcessingHeader := psdc.ConversionProcessingHeader

	partners := make([]*Partner, 0)
	for range dataProcessingPartner {
		partners = append(partners, &Partner{
			DeliveryDocument: *dataConversionProcessingHeader.ConvertedDeliveryDocument,
		})
	}

	return partners
}

func getBoolPtr(b bool) *bool {
	return &b
}
