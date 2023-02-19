package requests

type ConversionProcessingItemScheduleLine struct {
	ConvertingScheduleLine *string `json:"ConvertingScheduleLine"`
	ConvertedScheduleLine  *int    `json:"ConvertedScheduleLine"`
}
