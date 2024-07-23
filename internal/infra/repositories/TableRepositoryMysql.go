package repositories

import (
	"GastroReserve/internal/entities"
	"database/sql"
)

type TableRepositoryMysql struct {
	Db *sql.DB
}

func NewTableRepositoryMysql(db *sql.DB) *TableRepositoryMysql {
	return &TableRepositoryMysql{Db: db}
}

func (r *TableRepositoryMysql) CreateTable(table *entities.Table) error {
	_, err := r.Db.Exec("INSERT INTO tables (id,number,capacity) values(?,?,?)", table.Id, table.Number, table.Capacity)
	if err != nil {
		return err
	}
	return nil
}
func (r *TableRepositoryMysql) GetTable() ([]*entities.Table, error) {
	var tables []*entities.Table
	row, err := r.Db.Query("SELECT id,number,capacity FROM tables")
	if err != nil {
		return nil, err
	}
	for row.Next() {
		var table entities.Table
		err := row.Scan(&table.Id, &table.Number, &table.Capacity)
		if err != nil {
			return nil, err
		}
		tables = append(tables, &table)
	}
	return tables, nil
}
func (r *TableRepositoryMysql) GetTablePerNumber(number int) (entities.Table, error) {
	var table entities.Table
	row := r.Db.QueryRow("SELECT id,number,capacity FROM tables WHERE number = ?", number)
	err := row.Scan(&table.Id, &table.Number, &table.Capacity)
	if err != nil {
		return entities.Table{}, err
	}
	return table, nil
}
