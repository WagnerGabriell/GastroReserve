package usecases

import (
	dto "GastroReserve/internal/DTO"
	"GastroReserve/internal/infra/repositories"
)

type GetTablesEmptyDataUseCase struct {
	Repository repositories.ITableRepositoryMysql
}

func NewGetTablesEmptyDataUseCase(repository repositories.ITableRepositoryMysql) *GetTablesEmptyDataUseCase {
	return &GetTablesEmptyDataUseCase{Repository: repository}
}
func (u *GetTablesEmptyDataUseCase) Execute(data string) ([]*dto.TableOutputDTO, error) {
	var listReserveOutputDTO []*dto.TableOutputDTO
	row, err := u.Repository.GetTablesEmptyData(data)
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
