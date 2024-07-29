package dto

type TableInputDTO struct {
	Number   int `json:"number"`
	Capacity int `json:"capacity"`
}
type TableOutputDTO struct {
	Id       string
	Number   int
	Capacity int
}
type TableNumberInputDTO struct {
	Number int `json:"number"`
}
type TableDataInputDTO struct {
	Data string `json:"data"`
}
