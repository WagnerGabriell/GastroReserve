package usecases

import (
	dto "GastroReserve/internal/DTO"
	"GastroReserve/internal/infra/repositories"
)

type GetAllUsersUserCase struct {
	Repository repositories.IUserRepositoryMySql
}

func NewGetAllUsersUserCase(repository repositories.IUserRepositoryMySql) *GetAllUsersUserCase {
	return &GetAllUsersUserCase{Repository: repository}
}
func (u *GetAllUsersUserCase) Execute() ([]dto.UserOutputDTO, error) {
	var sliceOutputDto []dto.UserOutputDTO
	users, err := u.Repository.GetAllUsers()
	if err != nil {
		return nil, err
	}
	for _, user := range users {
		sliceOutputDto = append(sliceOutputDto, dto.UserOutputDTO{
			Id:          user.Id,
			Email:       user.Email,
			Password:    user.Password,
			Name:        user.Name,
			PhoneNumber: user.PhoneNumber,
			IsAdmin:     user.IsAdmin,
		})
	}
	return sliceOutputDto, nil
}
