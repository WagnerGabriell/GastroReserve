package entities

import (
	"github.com/google/uuid"
)

type Reserve struct {
	Id      string
	UserId  string
	TableId string
	Data    string
}

func NewReserve(tableId string, userId string, data string) *Reserve {
	return &Reserve{
		Id:      uuid.New().String(),
		UserId:  userId,
		TableId: tableId,
		Data:    data,
	}
}
