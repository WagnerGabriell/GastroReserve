package dto

type UserInputDTO struct {
	Email       string `json:"email"`
	Name        string `json:"name"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phoneNumber"`
}
type UserOutputDTO struct {
	Id          string
	Email       string
	Name        string
	Password    string
	PhoneNumber string
	IsAdmin     bool
}
type UserLoginInputDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
