package dpfm_api_input_reader

import (
	"convert-to-dpfm-delivery-document-from-delivery-notice-edi-for-smes/DPFM_API_Caller/requests"
)

func (sdc *SDC) ConvertToHeader() *requests.Header {
	data := sdc.Header
	return &requests.Header{
		ExchangedDeliveryNoticeDocumentIdentifier:                      data.ExchangedDeliveryNoticeDocumentIdentifier,
		DeliveryNoticeDocument:                                         data.DeliveryNoticeDocument,
		ExchangedDocumentContextSpecifiedTransactionIdentifier:         data.ExchangedDocumentContextSpecifiedTransactionIdentifier,
		ExchangedDeliveryNoticeDocumentName:                            data.ExchangedDeliveryNoticeDocumentName,
		ExchangeDeliveryNoticeDocumentTypeCode:                         data.ExchangeDeliveryNoticeDocumentTypeCode,
		ExchangedDeliveryNoticeDocumentIssueDate:                       data.ExchangedDeliveryNoticeDocumentIssueDate,
		ExchangedDeliveryNoticeDocumentPurposeCode:                     data.ExchangedDeliveryNoticeDocumentPurposeCode,
		ExchangedDeliveryNoticeDocumentRequestedResponseTypeCode:       data.ExchangedDeliveryNoticeDocumentRequestedResponseTypeCode,
		ExchangedDeliveryNoticeDocumentVersionIdentifier:               data.ExchangedDeliveryNoticeDocumentVersionIdentifier,
		ExchangedDeliveryNoticeDocumentCategoryCode:                    data.ExchangedDeliveryNoticeDocumentCategoryCode,
		ExchangedDeliveryNoticeDocumentSubtypeCode:                     data.ExchangedDeliveryNoticeDocumentSubtypeCode,
		NoteDeliveryNoticeSubjectText:                                  data.NoteDeliveryNoticeSubjectText,
		NoteDeliveryNoticeContentText:                                  data.NoteDeliveryNoticeContentText,
		NoteDeliveryNoticeIdentifier:                                   data.NoteDeliveryNoticeIdentifier,
		SpecifiedBinaryFileIdentifier:                                  data.SpecifiedBinaryFileIdentifier,
		SpecifiedBinaryFileVersionIdentifier:                           data.SpecifiedBinaryFileVersionIdentifier,
		SpecifiedBinaryFileNameText:                                    data.SpecifiedBinaryFileNameText,
		SpecifiedBineryFileURIIdentifier:                               data.SpecifiedBineryFileURIIdentifier,
		SpecifiedBineryFileMIMECode:                                    data.SpecifiedBineryFileMIMECode,
		SpecifiedBineryFileEncodingCode:                                data.SpecifiedBineryFileEncodingCode,
		SpecifiedBineryFileCode:                                        data.SpecifiedBineryFileCode,
		SpecifiedBinaryFileDescriptionText:                             data.SpecifiedBinaryFileDescriptionText,
		TradeSellerIdentifier:                                          data.TradeSellerIdentifier,
		TradeSellerGlobalIdentifier:                                    data.TradeSellerGlobalIdentifier,
		TradeSellerName:                                                data.TradeSellerName,
		TradeBillFromPartyRegisteredIdentifier:                         data.TradeBillFromPartyRegisteredIdentifier,
		TradeContactSellerIdentifier:                                   data.TradeContactSellerIdentifier,
		TradeContactSellerPersonName:                                   data.TradeContactSellerPersonName,
		TradeContactSellerDepartmentName:                               data.TradeContactSellerDepartmentName,
		SellerTelephoneNumber:                                          data.SellerTelephoneNumber,
		SellerFaxNumber:                                                data.SellerFaxNumber,
		SellerEmailAddress:                                             data.SellerEmailAddress,
		SellerAddressPostalCode:                                        data.SellerAddressPostalCode,
		SellerAddress:                                                  data.SellerAddress,
		TradeBuyerIdentifier:                                           data.TradeBuyerIdentifier,
		TradeBuyerGlobalIdentifier:                                     data.TradeBuyerGlobalIdentifier,
		TradeBuyerName:                                                 data.TradeBuyerName,
		TradeContactBuyerIdentifier:                                    data.TradeContactBuyerIdentifier,
		TradeContactBuyerPersonName:                                    data.TradeContactBuyerPersonName,
		TradeContactBuyerDepartmentName:                                data.TradeContactBuyerDepartmentName,
		BuyerTelephoneNumber:                                           data.BuyerTelephoneNumber,
		BuyerFaxNumber:                                                 data.BuyerFaxNumber,
		BuyerEmailAddress:                                              data.BuyerEmailAddress,
		BuyerAddressPostalCode:                                         data.BuyerAddressPostalCode,
		BuyerAddress:                                                   data.BuyerAddress,
		ReferencedOrdersDocumentIssureAssignedIdentifier:               data.ReferencedOrdersDocumentIssureAssignedIdentifier,
		ReferencedOrdersDocumentIssueDate:                              data.ReferencedOrdersDocumentIssueDate,
		ReferencedOrdersDocumentRevisionIdentifier:                     data.ReferencedOrdersDocumentRevisionIdentifier,
		ReferencedOrdersDocumentName:                                   data.ReferencedOrdersDocumentName,
		ReferencedOrdersDocumentInformationText:                        data.ReferencedOrdersDocumentInformationText,
		ReferencedOrdersDocumentInformationPurposeCode:                 data.ReferencedOrdersDocumentInformationPurposeCode,
		ReferencedOrdersDocumentInformationSubtypeCode:                 data.ReferencedOrdersDocumentInformationSubtypeCode,
		ProjectIdentifier:                                              data.ProjectIdentifier,
		ProjectName:                                                    data.ProjectName,
		ReferencedSalesOrderDocumentIssureAddignedIdentifier:           data.ReferencedSalesOrderDocumentIssureAddignedIdentifier,
		ReferencedSalesOrderDocumentIssueDate:                          data.ReferencedSalesOrderDocumentIssueDate,
		ReferencedSalesOrderDocumentRevisionIdentifier:                 data.ReferencedSalesOrderDocumentRevisionIdentifier,
		ReferencedSalesOrderDocumentName:                               data.ReferencedSalesOrderDocumentName,
		ReferencedSalesOrderDocumentInformationText:                    data.ReferencedSalesOrderDocumentInformationText,
		ReferencedSalesOrderDocumentSubtypeCode:                        data.ReferencedSalesOrderDocumentSubtypeCode,
		TradeShipToPartyIdentifier:                                     data.TradeShipToPartyIdentifier,
		TradeShipToPartyGlobalIdentifier:                               data.TradeShipToPartyGlobalIdentifier,
		TradeShipToPartyName:                                           data.TradeShipToPartyName,
		TradeShipToPartyContactIdentifier:                              data.TradeShipToPartyContactIdentifier,
		TradeShipToPartyContactPersonName:                              data.TradeShipToPartyContactPersonName,
		TradeShipToPartyContactDepartmentName:                          data.TradeShipToPartyContactDepartmentName,
		TradeShipToPartyContactPersonIdentifier:                        data.TradeShipToPartyContactPersonIdentifier,
		ShipToPartyTelephoneNumber:                                     data.ShipToPartyTelephoneNumber,
		ShipToPartyFaxNumber:                                           data.ShipToPartyFaxNumber,
		ShipToPartyEmailAddress:                                        data.ShipToPartyEmailAddress,
		ShipToPartyAddressPostalCode:                                   data.ShipToPartyAddressPostalCode,
		ShipToPartyAddress:                                             data.ShipToPartyAddress,
		TradeShipFromPartyIdentifier:                                   data.TradeShipFromPartyIdentifier,
		TradeShipFromPartyName:                                         data.TradeShipFromPartyName,
		TradeLogiName:                                                  data.TradeLogiName,
		TradeLogiContactIdentifier:                                     data.TradeLogiContactIdentifier,
		TradeLogiContactPersonName:                                     data.TradeLogiContactPersonName,
		TradeLogiContactDepartmentName:                                 data.TradeLogiContactDepartmentName,
		TradeLogiContactPersonIdentifier:                               data.TradeLogiContactPersonIdentifier,
		LogiTelephoneNumber:                                            data.LogiTelephoneNumber,
		SupplyChainEventIdentifier:                                     data.SupplyChainEventIdentifier,
		InstructedTemperatureControlCode:                               data.InstructedTemperatureControlCode,
		TradeTaxCalculatedAmount:                                       data.TradeTaxCalculatedAmount,
		TradeTaxTypeCode:                                               data.TradeTaxTypeCode,
		TradeTaxBasisAmount:                                            data.TradeTaxBasisAmount,
		TradeTaxCategoryCode:                                           data.TradeTaxCategoryCode,
		TradeTaxCategoryName:                                           data.TradeTaxCategoryName,
		TradeTaxRatePercent:                                            data.TradeTaxRatePercent,
		TradeTaxGrandTotalAmount:                                       data.TradeTaxGrandTotalAmount,
		TradeTaxCalculationMethod:                                      data.TradeTaxCalculationMethod,
		TradeSettlementMonetarySummationTotalTaxAmount:                 data.TradeSettlementMonetarySummationTotalTaxAmount,
		TradeDeliveryNoticeSettlementMonetarySummationGrandTotalAmount: data.TradeDeliveryNoticeSettlementMonetarySummationGrandTotalAmount,
		TradeDeliveryNoticeSettlementMonetarySummationNetTotalAmount:   data.TradeDeliveryNoticeSettlementMonetarySummationNetTotalAmount,
		TradeDeliveryNoticeMonetarySummationIncludingTaxesTotalAmount:  data.TradeDeliveryNoticeMonetarySummationIncludingTaxesTotalAmount,
	}
}

func (sdc *SDC) ConvertToItem(num int) *requests.Item {
	dataHeader := sdc.Header
	data := sdc.Header.Item[num]

	return &requests.Item{
		ExchangedDeliveryNoticeDocumentIdentifier:                                      dataHeader.ExchangedDeliveryNoticeDocumentIdentifier,
		DeliveryNoticeDocumentItemlineIdentifier:                                       data.DeliveryNoticeDocumentItemlineIdentifier,
		DeliveryNoticeDocumentItemlineStatusCode:                                       data.DeliveryNoticeDocumentItemlineStatusCode,
		DeliveryNoticeDocumentItemlineStatusReasonCode:                                 data.DeliveryNoticeDocumentItemlineStatusReasonCode,
		NoteDeliveryNoticeItemSubjectText:                                              data.NoteDeliveryNoticeItemSubjectText,
		NoteDeliveryNoticeItemContentText:                                              data.NoteDeliveryNoticeItemContentText,
		NoteDeliveryNoticeItemIdentifier:                                               data.NoteDeliveryNoticeItemIdentifier,
		ReferencedSalesOrderDocumentIssuerAssignedIdentifier:                           data.ReferencedSalesOrderDocumentIssuerAssignedIdentifier,
		ReferencedSalesOrderDocumentItemLineIdentifier:                                 data.ReferencedSalesOrderDocumentItemLineIdentifier,
		ReferencedSalesOrderDocumentRevisionIdentifier:                                 data.ReferencedSalesOrderDocumentRevisionIdentifier,
		ReferencedOrdersDocumentIssureAssignedIdentifier:                               data.ReferencedOrdersDocumentIssureAssignedIdentifier,
		ReferencedOrdersDocumentItemLineIdentifier:                                     data.ReferencedOrdersDocumentItemLineIdentifier,
		ReferencedOrdersDocumentRevisionIdentifier:                                     data.ReferencedOrdersDocumentRevisionIdentifier,
		TradePriceTypeCode:                                                             data.TradePriceTypeCode,
		TradeDeliveryPriceChargeAmount:                                                 data.TradeDeliveryPriceChargeAmount,
		TradePriceBasisQuantity:                                                        data.TradePriceBasisQuantity,
		TradePriceBasisUnitCode:                                                        data.TradePriceBasisUnitCode,
		SupplyChainTradeDeliveryPackageQuantity:                                        data.SupplyChainTradeDeliveryPackageQuantity,
		SupplyChainTradeDeliveryProductUnitQuantity:                                    data.SupplyChainTradeDeliveryProductUnitQuantity,
		SupplyChainTradeDeliveryPerPackageUnitQuantity:                                 data.SupplyChainTradeDeliveryPerPackageUnitQuantity,
		SupplyChainTradeDeliveryDespatchedQuantity:                                     data.SupplyChainTradeDeliveryDespatchedQuantity,
		SupplyChainTradeDeliveryRequestedQuantity:                                      data.SupplyChainTradeDeliveryRequestedQuantity,
		SupplyChainTradeDeliveryRemainingRequestedQuantity:                             data.SupplyChainTradeDeliveryRemainingRequestedQuantity,
		ItemTradeDeliverToPartyIdentifier:                                              data.ItemTradeDeliverToPartyIdentifier,
		ItemTradeDeliverToPartyGlobalIdentifier:                                        data.ItemTradeDeliverToPartyGlobalIdentifier,
		ItemTradeDeliverToPartyName:                                                    data.ItemTradeDeliverToPartyName,
		ItemTradeDeliverToPartyContactPersonName:                                       data.ItemTradeDeliverToPartyContactPersonName,
		ItemTradeDeliverToPartyContactDepartmentName:                                   data.ItemTradeDeliverToPartyContactDepartmentName,
		ItemDeliverToPartyPhoneNumber:                                                  data.ItemDeliverToPartyPhoneNumber,
		ItemDeliverToPartyFaxNumber:                                                    data.ItemDeliverToPartyFaxNumber,
		ItemDeliverToPartyEMailAddress:                                                 data.ItemDeliverToPartyEMailAddress,
		ItemDeliverToPartyAddressPostalCode:                                            data.ItemDeliverToPartyAddressPostalCode,
		ItemDeliverToPartyAddress:                                                      data.ItemDeliverToPartyAddress,
		SupplyChainDeliveryEventIdentifier:                                             data.SupplyChainDeliveryEventIdentifier,
		SupplyChainDeliveryEventOccurrenceDate:                                         data.SupplyChainDeliveryEventOccurrenceDate,
		SupplyChainEventTypeCode:                                                       data.SupplyChainEventTypeCode,
		SupplyChainEventRequirementOccurrenceDate:                                      data.SupplyChainEventRequirementOccurrenceDate,
		LogisticsLocationIdentification:                                                data.LogisticsLocationIdentification,
		LogisticsLocationName:                                                          data.LogisticsLocationName,
		SupplyChainEventPlannedOccurrenceDate:                                          data.SupplyChainEventPlannedOccurrenceDate,
		TradeTaxTypeCode:                                                               data.TradeTaxTypeCode,
		ItemTradeTaxBasisAmount:                                                        data.ItemTradeTaxBasisAmount,
		ItemTradeTaxCategoryCode:                                                       data.ItemTradeTaxCategoryCode,
		TradeTaxCategoryName:                                                           data.TradeTaxCategoryName,
		ItemTradeTaxRateApplicablePercent:                                              data.ItemTradeTaxRateApplicablePercent,
		ItemTradeTaxGrandTotalAmount:                                                   data.ItemTradeTaxGrandTotalAmount,
		ItemTradeDeliveryNoticeSettlementMonetarySummationNetTotalAmount:               data.ItemTradeDeliveryNoticeSettlementMonetarySummationNetTotalAmount,
		ItemTradeDeliveryNoticeSettlementMonetarySummationIncludingTaxesNetTotalAmount: data.ItemTradeDeliveryNoticeSettlementMonetarySummationIncludingTaxesNetTotalAmount,
		TradeProductIdentifier:                                                         data.TradeProductIdentifier,
		TradeProductGlobalIdentifier:                                                   data.TradeProductGlobalIdentifier,
		TradeProductSellerAssignedIdentifier:                                           data.TradeProductSellerAssignedIdentifier,
		TradeProductBuyerAssignedIdentifier:                                            data.TradeProductBuyerAssignedIdentifier,
		TradeProductManufacturerAssignedIdentifier:                                     data.TradeProductManufacturerAssignedIdentifier,
		TradeProductName:                                                               data.TradeProductName,
		TradeProductDescription:                                                        data.TradeProductDescription,
		TradeProductTypeCode:                                                           data.TradeProductTypeCode,
		TradeProductEndItemTypeCode:                                                    data.TradeProductEndItemTypeCode,
		TradeProductSizeCode:                                                           data.TradeProductSizeCode,
		TradeProductSizeDescriptionText:                                                data.TradeProductSizeDescriptionText,
		TradeManufacturerIdentifier:                                                    data.TradeManufacturerIdentifier,
		TradeManufacturerName:                                                          data.TradeManufacturerName,
		ReferencedLogisticsPackageUnitQuantity:                                         data.ReferencedLogisticsPackageUnitQuantity,
		ReferencedLogisticsPackageQuantityUnitCode:                                     data.ReferencedLogisticsPackageQuantityUnitCode,
		ReferencedLogisticsPackageTypeCode:                                             data.ReferencedLogisticsPackageTypeCode,
		ReferencedLogisticsPackageIdentifier:                                           data.ReferencedLogisticsPackageIdentifier,
	}
}