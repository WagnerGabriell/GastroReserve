package usecases

import (
	dto "GastroReserve/internal/DTO"
	"GastroReserve/internal/entities"
	"GastroReserve/internal/infra/repositories"
)

type CreateTableUseCase struct {
	RepositoryTable repositories.ITableRepositoryMysql
}

func NewCreateTableUseCase(repositoryTable repositories.ITableRepositoryMysql) *CreateTableUseCase {
	return &CreateTableUseCase{RepositoryTable: repositoryTable}
}

func (u *CreateTableUseCase) Execute(tableInputDTO dto.TableInputDTO) (dto.TableOutputDTO, error) {
	table := entities.NewTable(tableInputDTO.Number, tableInputDTO.Capacity)
	err := u.RepositoryTable.CreateTable(table)
	if err != nil {
		return dto.TableOutputDTO{}, err
	}
	return dto.TableOutputDTO{
		Id:       table.Id,
		Number:   table.Number,
		Capacity: table.Capacity,
	}, nil
}
