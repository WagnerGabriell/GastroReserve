package usecases

import (
	dto "GastroReserve/internal/DTO"
	"GastroReserve/internal/infra/repositories"
)

type GetTablePerNumberUseCase struct {
	Repository repositories.ITableRepositoryMysql
}

func NewGetTablePerNumberUseCase(repository repositories.ITableRepositoryMysql) *GetTablePerNumberUseCase {
	return &GetTablePerNumberUseCase{Repository: repository}
}
func (u *GetTablePerNumberUseCase) Execute(number int) (*dto.TableOutputDTO, error) {
	table, err := u.Repository.GetTablePerNumber(number)
	if err != nil {
		return &dto.TableOutputDTO{}, err
	}
	return &dto.TableOutputDTO{
		Id:       table.Id,
		Number:   table.Number,
		Capacity: table.Capacity,
	}, nil
}
