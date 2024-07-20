package repositories

import "GastroReserve/internal/entities"

type IReserveRepositoryMysql interface {
	CreateReserve(reserve entities.Reserve) error
	GetReserve() ([]*entities.Reserve, error)
	GetReservePerName(name string)
}
