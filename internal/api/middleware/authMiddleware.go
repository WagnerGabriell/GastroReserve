package middleware

import (
	"GastroReserve/internal/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthMiddleWare struct {
	VerificarTokenUseCase *usecases.VerificaTokenUseCase
}

func NewAuthMiddleWare(verificarTokenUseCase *usecases.VerificaTokenUseCase) *AuthMiddleWare {
	return &AuthMiddleWare{VerificarTokenUseCase: verificarTokenUseCase}
}

func (m *AuthMiddleWare) VerificarTokenMiddleWare(c *gin.Context) {
	jwtToken := c.GetHeader("Token")
	if jwtToken == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Token invalid"})
		c.Abort()
	}
	claimsOutputDTO, err := m.VerificarTokenUseCase.Execute(jwtToken)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err})
		c.Abort()
	}
	c.Set("Id", claimsOutputDTO.Id)
	c.Set("IsAdmin", claimsOutputDTO.IsAdmin)
	c.Next()
}
