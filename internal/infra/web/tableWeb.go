package web

import (
	dto "GastroReserve/internal/DTO"
	"GastroReserve/internal/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TableWeb struct {
	CreateTable       *usecases.CreateTableUseCase
	GetAllTable       *usecases.GetAllTableUseCase
	GetTablePerNumber *usecases.GetTablePerNumberUseCase
}

func NewTableWeb(createTable *usecases.CreateTableUseCase, getAllTable *usecases.GetAllTableUseCase, getTablePerNumber *usecases.GetTablePerNumberUseCase) *TableWeb {
	return &TableWeb{
		CreateTable:       createTable,
		GetAllTable:       getAllTable,
		GetTablePerNumber: getTablePerNumber,
	}
}

func (w *TableWeb) CreateTableWeb(c *gin.Context) {
	var tableInput dto.TableInputDTO
	err := c.ShouldBindJSON(&tableInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err})
		return
	}
	tableOut, err := w.CreateTable.Execute(tableInput)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"Table": tableOut})
}

func (w *TableWeb) GetAllTableWeb(c *gin.Context) {
	tableOutput, err := w.GetAllTable.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Table": tableOutput})
}
func (w *TableWeb) GetTablePerNumberWeb(c *gin.Context) {
	var number dto.TableNumberInputDTO
	err := c.ShouldBindJSON(&number)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err})
		return
	}
	tableOutput, err := w.GetTablePerNumber.Execute(number.Number)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Table": tableOutput})
}
