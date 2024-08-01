package usecases

import (
	dto "GastroReserve/internal/DTO"
	"GastroReserve/internal/entities"
	"GastroReserve/internal/infra/repositories"
)

type CreateUserUseCase struct {
	Repository repositories.IUserRepositoryMySql
}

func NewCreateUserUseCase(repository repositories.IUserRepositoryMySql) *CreateUserUseCase {
	return &CreateUserUseCase{Repository: repository}
}
func (u *CreateUserUseCase) Execute(userInput dto.UserInputDTO) (*dto.UserOutputDTO, error) {
	user := entities.NewUser(userInput.Email, userInput.Name, userInput.PhoneNumber, userInput.Password)
	err := u.Repository.CreateUser(*user)
	if err != nil {
		return &dto.UserOutputDTO{}, err
	}

	return &dto.UserOutputDTO{
		Id:          user.Id,
		Email:       user.Email,
		Name:        user.Name,
		Password:    user.Password,
		PhoneNumber: user.PhoneNumber,
		IsAdmin:     user.IsAdmin,
	}, nil
}
