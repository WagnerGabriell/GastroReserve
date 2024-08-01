package repositories

import "GastroReserve/internal/entities"

type IUserRepositoryMySql interface {
	GetAllUsers() ([]*entities.User, error)
	CreateUser(user entities.User) error
	GetUser(id string) (*entities.User, error)
	GetUserPerEmail(email string) (*entities.User, error)
	BecomeAdm(id string) error
}
