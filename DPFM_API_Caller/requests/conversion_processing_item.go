package requests

type ConversionProcessingItem struct {
	ConvertingDeliveryDocumentItem         *string `json:"ConvertingDeliveryDocumentItem"`
	ConvertedDeliveryDocumentItem          *int    `json:"ConvertedDeliveryDocumentItem"`
	ConvertingProject                      *string `json:"ConvertingProject"`
	ConvertedProject                       *string `json:"ConvertedProject"`
	ConvertingTransactionTaxClassification *string `json:"ConvertingTransactionTaxClassification"`
	ConvertedTransactionTaxClassification  *string `json:"ConvertedTransactionTaxClassification"`
}
