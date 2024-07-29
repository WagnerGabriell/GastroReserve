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
	_, err := r.Db.Exec("INSERT INTO reserve (id,name,phoneNumber,tableId,data) value (?,?,?,?,?)", reserve.Id, reserve.Custumer.Name, reserve.Custumer.PhoneNumber, reserve.TableId, reserve.Data)
	if err != nil {
		return err
	}
	return nil
}
func (r *ReserveRepositoryMysql) GetReserve() ([]*entities.Reserve, error) {
	row, err := r.Db.Query("SELECT id,name,phoneNumber,tableId,data FROM reserve")
	if err != nil {
		return nil, err
	}
	var reserves []*entities.Reserve
	for row.Next() {
		var reserve entities.Reserve
		err := row.Scan(&reserve.Id, &reserve.Custumer.Name, &reserve.Custumer.PhoneNumber, &reserve.TableId, &reserve.Data)
		if err != nil {
			return nil, err
		}
		reserves = append(reserves, &reserve)
	}
	return reserves, nil
}
func (r *ReserveRepositoryMysql) GetReservePerName(name string) (*entities.Reserve, error) {
	var reserve entities.Reserve
	row := r.Db.QueryRow("SELECT id,name,phoneNumber,tableId,data FROM reserve WHERE name = ?", name)
	err := row.Scan(&reserve.Id, &reserve.Custumer.Name, &reserve.Custumer.PhoneNumber, &reserve.TableId, reserve.Data)
	if err != nil {
		return nil, err
	}
	return &reserve, nil
}
