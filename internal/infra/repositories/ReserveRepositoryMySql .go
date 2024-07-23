package repositories

import (
	"GastroReserve/internal/entities"
	"database/sql"
)

type ReserveRepositoryMysql struct {
	Db *sql.DB
}

func NewReserveRepositoryMysql(db *sql.DB) *ReserveRepositoryMysql {
	return &ReserveRepositoryMysql{Db: db}
}

func (r *ReserveRepositoryMysql) CreateReserve(reserve *entities.Reserve) error {
	_, err := r.Db.Exec("INSERT INTO Reserve (id,name,phoneNumber,tableId,data) value (?,?,?,?,?)", reserve.Id, reserve.Custumer.Name, reserve.Custumer.PhoneNumber, reserve.TableId, reserve.Data)
	if err != nil {
		return err
	}
	return nil
}
