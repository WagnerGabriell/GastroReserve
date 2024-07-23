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
	tableWeb := web.NewTableWeb(createTable, getAllTable, getTablePerNumber)

	reserveRepository := repositories.NewReserveRepositoryMysql(db)
	createRepository := usecases.NewCreateReserveUseCase(reserveRepository)
	reserveWeb := web.NewReserveWeb(createRepository)

	r := gin.Default()
	Table := r.Group("/table")
	{
		Table.POST("/create", tableWeb.CreateTableWeb)
		Table.GET("/list", tableWeb.GetAllTableWeb)
		Table.GET("/getNumber", tableWeb.GetTablePerNumberWeb)
	}
	Reserve := r.Group("/reserve")
	{
		Reserve.POST("/create", reserveWeb.CreateReserveWeb)
	}
	r.Run(":3030")
}
