package usecases

import (
	dto "GastroReserve/internal/DTO"
	"GastroReserve/internal/entities"
	"GastroReserve/internal/infra/repositories"

	"golang.org/x/crypto/bcrypt"
)

type RegisterUserUseCase struct {
	Repository repositories.IUserRepositoryMySql
}

func NewRegisterUserUseCase(repository repositories.IUserRepositoryMySql) *RegisterUserUseCase {
	return &RegisterUserUseCase{Repository: repository}
}
func (u *RegisterUserUseCase) Execute(userInput dto.UserInputDTO) error {
	user := entities.NewUser(userInput.Email, userInput.Name, userInput.PhoneNumber, userInput.Password)
	hash, err := bcrypt.GenerateFromPassword([]byte(userInput.Password), 10)
	if err != nil {
		return err
	}
	user.Password = string(hash)
	err = u.Repository.CreateUser(*user)
	if err != nil {
		return err
	}
	return nil
}
