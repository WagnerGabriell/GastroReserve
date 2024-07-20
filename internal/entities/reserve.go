package entities

import (
	"github.com/google/uuid"
)

type Reserve struct {
	Id       string
	Custumer Custumer
	TableId  string
	Data     string
}

func NewReserve(tableId string, custumer Custumer, data string) *Reserve {
	return &Reserve{
		Id:       uuid.New().String(),
		Custumer: custumer,
		TableId:  tableId,
		Data:     data,
	}
}
