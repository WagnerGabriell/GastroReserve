package web

import (
	dto "GastroReserve/internal/DTO"
	"GastroReserve/internal/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ReserveWeb struct {
	CreateReserveUseCase *usecases.CreateReserveUseCase
}

func NewReserveWeb(createReserveUseCase *usecases.CreateReserveUseCase) *ReserveWeb {
	return &ReserveWeb{
		CreateReserveUseCase: createReserveUseCase,
	}
}

func (w *ReserveWeb) CreateReserveWeb(c *gin.Context) {
	var reserveInputDTO dto.ReserveInputDTO
	err := c.ShouldBindJSON(&reserveInputDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	reserveOutputDTO, err := w.CreateReserveUseCase.Execute(&reserveInputDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"ReserveOutput": reserveOutputDTO})
}
