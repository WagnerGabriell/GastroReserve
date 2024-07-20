package web

import (
	dto "GastroReserve/internal/DTO"
	"GastroReserve/internal/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TableWeb struct {
	CreateTable *usecases.CreateTableUseCase
}

func NewTableWeb(createTable *usecases.CreateTableUseCase) *TableWeb {
	return &TableWeb{
		CreateTable: createTable,
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
	c.JSON(http.StatusInternalServerError, gin.H{"Table": tableOut})
}
