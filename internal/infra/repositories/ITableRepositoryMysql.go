package repositories

import "GastroReserve/internal/entities"

type ITableRepositoryMysql interface {
	CreateTable(table *entities.Table) error
	GetTable() ([]*entities.Table, error)
	GetTablePerNumber(number int) (*entities.Table, error)
	GetTablesEmptyData(data string) ([]*entities.Table, error)
}
