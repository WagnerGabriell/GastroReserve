package web

import (
	dto "GastroReserve/internal/DTO"
	"GastroReserve/internal/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ReserveWeb struct {
	CreateReserveUseCase *usecases.CreateReserveUseCase
	GetAllReserveUseCase *usecases.GetAllReserveUseCase
}

func NewReserveWeb(createReserveUseCase *usecases.CreateReserveUseCase, getAllReserveUseCase *usecases.GetAllReserveUseCase) *ReserveWeb {
	return &ReserveWeb{
		CreateReserveUseCase: createReserveUseCase,
		GetAllReserveUseCase: getAllReserveUseCase,
	}
}

func (w *ReserveWeb) CreateReserveWeb(c *gin.Context) {
	var reserveInputDTO dto.ReserveInputDTO
	err := c.ShouldBindJSON(&reserveInputDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	id, exists := c.Get("Id")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	reserveOutputDTO, err := w.CreateReserveUseCase.Execute(&reserveInputDTO, id.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"ReserveOutput": reserveOutputDTO})
}

func (w *ReserveWeb) GetAllReserveWeb(c *gin.Context) {
	reserveOutput, err := w.GetAllReserveUseCase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Reserve": reserveOutput})
}
