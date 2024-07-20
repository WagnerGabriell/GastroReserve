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
	db, err := sql.Open("mysql", "root:docker@tcp(localhost:3306)/GastroReserve")
	if err != nil {
		panic(err)
	}
	tableRepository := repositories.NewTableRepositoryMysql(db)
	createTable := usecases.NewCreateTableUseCase(tableRepository)
	tableWeb := web.NewTableWeb(createTable)
	r := gin.Default()
	Table := r.Group("/table")
	{
		Table.POST("/create", tableWeb.CreateTableWeb)
	}
	r.Run(":3030")
}
