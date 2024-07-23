package usecases

import (
	dto "GastroReserve/internal/DTO"
	"GastroReserve/internal/infra/repositories"
)

type GetAllTableUseCase struct {
	Repository repositories.ITableRepositoryMysql
}

func NewGetAllTableUseCase(repositoy repositories.ITableRepositoryMysql) *GetAllTableUseCase {
	return &GetAllTableUseCase{Repository: repositoy}
}
func (u *GetAllTableUseCase) Execute() ([]*dto.TableOutputDTO, error) {
	var TableOutputDTO []*dto.TableOutputDTO
	tables, err := u.Repository.GetTable()
	if err != nil {
		return nil, err
	}
	for _, t := range tables {
		TableOutputDTO = append(TableOutputDTO, &dto.TableOutputDTO{
			Id:       t.Id,
			Number:   t.Number,
			Capacity: t.Capacity,
		})
	}
	return TableOutputDTO, nil
}
