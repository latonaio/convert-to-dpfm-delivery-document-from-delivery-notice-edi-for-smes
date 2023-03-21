package dpfm_api_processing_formatter

type ProcessingFormatterSDC struct {
	Header                     *Header                     `json:"Header"`
	ConversionProcessingHeader *ConversionProcessingHeader `json:"ConversionProcessingHeader"`
	Item                       []*Item                     `json:"Item"`
	ConversionProcessingItem   []*ConversionProcessingItem `json:"ConversionProcessingItem"`
	Address                    []*Address                  `json:"Address"`
	Partner                    []*Partner                  `json:"Partner"`
}

type ConversionProcessingKey struct {
	SystemConvertTo       string   `json:"SystemConvertTo"`
	SystemConvertFrom     string   `json:"SystemConvertFrom"`
	LabelConvertTo        string   `json:"LabelConvertTo"`
	LabelConvertFrom      string   `json:"LabelConvertFrom"`
	CodeConvertFromInt    *int     `json:"CodeConvertFromInt"`
	CodeConvertFromFloat  *float32 `json:"CodeConvertFromFloat"`
	CodeConvertFromString *string  `json:"CodeConvertFromString"`
	BusinessPartner       int      `json:"BusinessPartner"`
}

type ConversionProcessingCommonQueryGets struct {
	CodeConversionID      int      `json:"CodeConversionID"`
	SystemConvertTo       string   `json:"SystemConvertTo"`
	SystemConvertFrom     string   `json:"SystemConvertFrom"`
	LabelConvertTo        string   `json:"LabelConvertTo"`
	LabelConvertFrom      string   `json:"LabelConvertFrom"`
	CodeConvertFromInt    *int     `json:"CodeConvertFromInt"`
	CodeConvertFromFloat  *float32 `json:"CodeConvertFromFloat"`
	CodeConvertFromString *string  `json:"CodeConvertFromString"`
	CodeConvertToInt      *int     `json:"CodeConvertToInt"`
	CodeConvertToFloat    *float32 `json:"CodeConvertToFloat"`
	CodeConvertToString   *string  `json:"CodeConvertToString"`
	BusinessPartner       int      `json:"BusinessPartner"`
}

type Header struct {
	ConvertingDeliveryDocument      string  `json:"ConvertingDeliveryDocument"`
	ConvertingBuyer                 *string `json:"ConvertingBuyer"`
	ConvertingSeller                *string `json:"ConvertingSeller"`
	ConvertingDeliverToParty        *string `json:"ConvertingDeliverToParty"`
	ConvertingDeliverFromParty      *string `json:"ConvertingDeliverFromParty"`
	DeliverFromPlant                *string `json:"DeliverFromPlant"`
	ConvertingBillToParty           *string `json:"ConvertingBillToParty"`
	ConvertingBillFromParty         *string `json:"ConvertingBillFromParty"`
	ConvertingPayer                 *string `json:"ConvertingPayer"`
	ConvertingPayee                 *string `json:"ConvertingPayee"`
	ConvertingReferenceDocument     *string `json:"ConvertingReferenceDocument"`
	ConvertingReferenceDocumentItem *string `json:"ConvertingReferenceDocumentItem"`
	ConvertingOrderID               *string `json:"ConvertingOrderID"`
	ConvertingOrderItem             *string `json:"ConvertingOrderItem"`
	DocumentDate                    *string `json:"DocumentDate"`
	PlannedGoodsIssueDate           *string `json:"PlannedGoodsIssueDate"`
	PlannedGoodsReceiptDate         *string `json:"PlannedGoodsReceiptDate"`
	CreationDate                    *string `json:"CreationDate"`
	CreationTime                    *string `json:"CreationTime"`
	LastChangeDate                  *string `json:"LastChangeDate"`
	LastChangeTime                  *string `json:"LastChangeTime"`
	GoodsIssueOrReceiptSlipNumber   *string `json:"GoodsIssueOrReceiptSlipNumber"`
	HeaderDeliveryBlockStatus       *bool   `json:"HeaderDeliveryBlockStatus"`
	HeaderIssuingBlockStatus        *bool   `json:"HeaderIssuingBlockStatus"`
	HeaderReceivingBlockStatus      *bool   `json:"HeaderReceivingBlockStatus"`
	IsCancelled                     *bool   `json:"IsCancelled"`
	IsMarkedForDeletion             *bool   `json:"IsMarkedForDeletion"`
	ConvertingProject               *string `json:"ConvertingProject"`
}

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
	ConvertingBillToParty           *string `json:"ConvertingBillToParty"`
	ConvertedBillToParty            *int    `json:"ConvertedBillToParty"`
	ConvertingBillFromParty         *string `json:"ConvertingBillFromParty"`
	ConvertedBillFromParty          *int    `json:"ConvertedBillFromParty"`
	ConvertingPayer                 *string `json:"ConvertingPayer"`
	ConvertedPayer                  *int    `json:"ConvertedPayer"`
	ConvertingPayee                 *string `json:"ConvertingPayee"`
	ConvertedPayee                  *int    `json:"ConvertedPayee"`
	ConvertingReferenceDocument     *string `json:"ConvertingReferenceDocument"`
	ConvertedReferenceDocument      *int    `json:"ConvertedReferenceDocument"`
	ConvertingReferenceDocumentItem *string `json:"ConvertingReferenceDocumentItem"`
	ConvertedReferenceDocumentItem  *int    `json:"ConvertedReferenceDocumentItem"`
}

type Item struct {
	ConvertingDeliveryDocument                 string   `json:"ConvertingDeliveryDocument"`
	ConvertingDeliveryDocumentItem             string   `json:"ConvertingDeliveryDocumentItem"`
	ConvertingBuyer                            *string  `json:"ConvertingBuyer"`
	ConvertingSeller                           *string  `json:"ConvertingSeller"`
	ConvertingDeliverToParty                   *string  `json:"ConvertingDeliverToParty"`
	ConvertingDeliverFromParty                 *string  `json:"ConvertingDeliverFromParty"`
	DeliverFromPlant                           *string  `json:"DeliverFromPlant"`
	ConvertingBillToParty                      *string  `json:"ConvertingBillToParty"`
	ConvertingBillFromParty                    *string  `json:"ConvertingBillFromParty"`
	ConvertingPayer                            *string  `json:"ConvertingPayer"`
	ConvertingPayee                            *string  `json:"ConvertingPayee"`
	ConvertingStockConfirmationBusinessPartner *string  `json:"ConvertingStockConfirmationBusinessPartner"`
	ConvertingStockConfirmationPlant           *string  `json:"ConvertingStockConfirmationPlant"`
	DeliveryDocumentItemText                   *string  `json:"DeliveryDocumentItemText"`
	Product                                    *string  `json:"Product"`
	ProductStandardID                          *string  `json:"ProductStandardID"`
	DeliveryUnit                               *string  `json:"DeliveryUnit"`
	CreationDate                               *string  `json:"CreationDate"`
	CreationTime                               *string  `json:"CreationTime"`
	LastChangeDate                             *string  `json:"LastChangeDate"`
	LastChangeTime                             *string  `json:"LastChangeTime"`
	NetAmount                                  *float32 `json:"NetAmount"`
	GrossAmount                                *float32 `json:"GrossAmount"`
	ConvertingOrderID                          *string  `json:"ConvertingOrderID"`
	ConvertingOrderItem                        *string  `json:"ConvertingOrderItem"`
	ConvertingProject                          *string  `json:"ConvertingProject"`
	ConvertingReferenceDocument                *string  `json:"ConvertingReferenceDocument"`
	ConvertingReferenceDocumentItem            *string  `json:"ConvertingReferenceDocumentItem"`
	ConvertingTransactionTaxClassification     *string  `json:"ConvertingTransactionTaxClassification"`
	ItemDeliveryBlockStatus                    *bool    `json:"ItemDeliveryBlockStatus"`
	ItemIssuingBlockStatus                     *bool    `json:"ItemIssuingBlockStatus"`
	ItemReceivingBlockStatus                   *bool    `json:"ItemReceivingBlockStatus"`
	ItemBillingBlockStatus                     *bool    `json:"ItemBillingBlockStatus"`
	IsCancelled                                *bool    `json:"IsCancelled"`
	IsMarkedForDeletion                        *bool    `json:"IsMarkedForDeletion"`
}

type ConversionProcessingItem struct {
	ConvertingDeliveryDocumentItem         *string `json:"ConvertingDeliveryDocumentItem"`
	ConvertedDeliveryDocumentItem          *int    `json:"ConvertedDeliveryDocumentItem"`
	ConvertingProject                      *string `json:"ConvertingProject"`
	ConvertedProject                       *string `json:"ConvertedProject"`
	ConvertingTransactionTaxClassification *string `json:"ConvertingTransactionTaxClassification"`
	ConvertedTransactionTaxClassification  *string `json:"ConvertedTransactionTaxClassification"`
}

type Address struct {
	ConvertingDeliveryDocument string  `json:"ConvertingDeliveryDocument"`
	PostalCode                 *string `json:"PostalCode"`
}

type Partner struct {
	ConvertingDeliveryDocument string `json:"ConvertingDeliveryDocument"`
}
