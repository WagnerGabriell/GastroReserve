package entities

import "github.com/google/uuid"

type Custumer struct {
	Id          string
	Name        string
	PhoneNumber string
}

func NewCustumer(name string, phoneNumber string) *Custumer {
	return &Custumer{
		Id:          uuid.New().String(),
		Name:        name,
		PhoneNumber: phoneNumber,
	}
}
