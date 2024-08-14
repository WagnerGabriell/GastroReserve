package dto

type ReserveInputDTO struct {
	TableId string `json:"tableId"`
	Data    string `json:"data"`
}

type ReserveOutputDTO struct {
	Id      string
	UserId  string
	TableId string
	Data    string
}
