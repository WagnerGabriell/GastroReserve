package main

import (
	"GastroReserve/internal/api/middleware"
	"GastroReserve/internal/infra/repositories"
	"GastroReserve/internal/infra/web"
	"GastroReserve/internal/usecases"
	"database/sql"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	db, err := sql.Open("mysql", os.Getenv("CONECTIONSTRING"))
	if err != nil {
		panic(err)
	}
	tableRepository := repositories.NewTableRepositoryMysql(db)
	createTable := usecases.NewCreateTableUseCase(tableRepository)
	getAllTable := usecases.NewGetAllTableUseCase(tableRepository)
	getTablePerNumber := usecases.NewGetTablePerNumberUseCase(tableRepository)
	GetTablesEmptyData := usecases.NewGetTablesEmptyDataUseCase(tableRepository)
	tableWeb := web.NewTableWeb(createTable, getAllTable, getTablePerNumber, GetTablesEmptyData)

	reserveRepository := repositories.NewReserveRepositoryMysql(db)
	createReserve := usecases.NewCreateReserveUseCase(reserveRepository)
	getReserves := usecases.NewGetAllReserveUseCase(reserveRepository)
	reserveWeb := web.NewReserveWeb(createReserve, getReserves)

	userRepository := repositories.NewUserRepositoryMySql(db)
	loginUser := usecases.NewLoginUseCase(userRepository)

	registroUser := usecases.NewRegisterUserUseCase(userRepository)
	userWeb := web.NewUserWeb(registroUser, loginUser)

	verificationToken := usecases.NewVerifyTokenUseCase()
	authMiddleWare := middleware.NewAuthMiddleWare(verificationToken)
	isAdminMiddleWare := middleware.NewIsAdminMiddleWare()

	r := gin.Default()
	Table := r.Group("/table")
	{
		Table.POST("/create", authMiddleWare.VerificarTokenMiddleWare, isAdminMiddleWare.VerificarIsAdminMiddleWare, tableWeb.CreateTableWeb)
		Table.GET("/list", tableWeb.GetAllTableWeb)
		Table.GET("/getNumber", tableWeb.GetTablePerNumberWeb)
		Table.POST("/empty", tableWeb.GetTablesEmptyDataWeb)
	}
	Reserve := r.Group("/reserve")
	{
		Reserve.GET("/list", authMiddleWare.VerificarTokenMiddleWare, reserveWeb.GetAllReserveWeb)
		Reserve.POST("/create", authMiddleWare.VerificarTokenMiddleWare, reserveWeb.CreateReserveWeb)
	}
	User := r.Group("/user")
	{
		User.POST("/register", userWeb.RegisterUserWeb)
		User.POST("/login", userWeb.LoginUserWeb)
	}
	r.Run(":3030")
}
