package entities

import "github.com/google/uuid"

type User struct {
	Id          string
	Email       string
	Password    string
	Name        string
	PhoneNumber string
	IsAdmin     bool
}

func NewUser(emal string, name string, phoneNumber string, password string) *User {
	return &User{
		Id:          uuid.New().String(),
		Email:       emal,
		Name:        name,
		PhoneNumber: phoneNumber,
		Password:    password,
		IsAdmin:     false,
	}
}
