package entities

import "github.com/google/uuid"

type Table struct {
	Id       string
	Number   int
	Capacity int
}

func NewTable(number int, capacity int) *Table {
	return &Table{
		Id:       uuid.New().String(),
		Number:   number,
		Capacity: capacity,
	}
}
