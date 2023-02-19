package requests

type ConversionProcessingHeader struct {
	ConvertingDeliveryDocument      *string `json:"ConvertingDeliveryDocument"`
	ConvertedDeliveryDocument       *int    `json:"ConvertedDeliveryDocument"`
	ConvertingBuyer                 *string `json:"ConvertingBuyer"`
	ConvertedBuyer                  *int    `json:"ConvertedBuyer"`
	ConvertingSeller                *string `json:"ConvertingSeller"`
	ConvertedSeller                 *int    `json:"ConvertedSeller"`
	ConvertingDeliverToParty        *string `json:"ConvertingDeliverToParty"`
	ConvertedDeliverToParty         *int    `json:"ConvertedDeliverToParty"`
	ConvertingDeliverFromParty      *string `json:"ConvertingDeliverFromParty"`
	ConvertedDeliverFromParty       *int    `json:"ConvertedDeliverFromParty"`
	ConvertingReferenceDocument     *string `json:"ConvertingReferenceDocument"`
	ConvertedReferenceDocument      *int    `json:"ConvertedReferenceDocument"`
	ConvertingReferenceDocumentItem *string `json:"ConvertingReferenceDocumentItem"`
	ConvertedReferenceDocumentItem  *int    `json:"ConvertedReferenceDocumentItem"`
}
