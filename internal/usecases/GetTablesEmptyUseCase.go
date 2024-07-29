package usecases

import (
	dto "GastroReserve/internal/DTO"
	"GastroReserve/internal/infra/repositories"
)

type GetTablesEmptyUseCase struct {
	Repository repositories.ITableRepositoryMysql
}

func NewGetTablesEmptyUseCase(repository repositories.ITableRepositoryMysql) *GetTablesEmptyUseCase {
	return &GetTablesEmptyUseCase{Repository: repository}
}
func (u *GetTablesEmptyUseCase) Execute(data string) ([]*dto.TableOutputDTO, error) {
	var listReserveOutputDTO []*dto.TableOutputDTO
	row, err := u.Repository.GetTablesEmpty(data)
	if err != nil {
		return nil, err
	}
	for _, r := range row {
		listReserveOutputDTO = append(listReserveOutputDTO, &dto.TableOutputDTO{
			Id:       r.Id,
			Number:   r.Number,
			Capacity: r.Capacity,
		})
	}
	return listReserveOutputDTO, nil
}
