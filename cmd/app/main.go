package main

import (
	"GastroReserve/internal/infra/repositories"
	"GastroReserve/internal/infra/web"
	"GastroReserve/internal/usecases"
	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:Gastro@tcp(localhost:3306)/GastroReserve")
	if err != nil {
		panic(err)
	}
	tableRepository := repositories.NewTableRepositoryMysql(db)
	createTable := usecases.NewCreateTableUseCase(tableRepository)
	getAllTable := usecases.NewGetAllTableUseCase(tableRepository)
	getTablePerNumber := usecases.NewGetTablePerNumberUseCase(tableRepository)
	getTablesEmpty := usecases.NewGetTablesEmptyUseCase(tableRepository)
	tableWeb := web.NewTableWeb(createTable, getAllTable, getTablePerNumber, getTablesEmpty)

	reserveRepository := repositories.NewReserveRepositoryMysql(db)
	createReserve := usecases.NewCreateReserveUseCase(reserveRepository)
	getReserves := usecases.NewGetAllReserveUseCase(reserveRepository)
	reserveWeb := web.NewReserveWeb(createReserve, getReserves)

	r := gin.Default()
	Table := r.Group("/table")
	{
		Table.POST("/create", tableWeb.CreateTableWeb)
		Table.GET("/list", tableWeb.GetAllTableWeb)
		Table.GET("/getNumber", tableWeb.GetTablePerNumberWeb)
		Table.POST("/empty", tableWeb.GetTablesEmptyWeb)
	}
	Reserve := r.Group("/reserve")
	{
		Reserve.GET("/list", reserveWeb.GetAllReserveWeb)
		Reserve.POST("/create", reserveWeb.CreateReserveWeb)
	}
	r.Run(":3030")
}
