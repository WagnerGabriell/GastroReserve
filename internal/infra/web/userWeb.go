package web

import (
	dto "GastroReserve/internal/DTO"
	"GastroReserve/internal/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserWeb struct {
	RegistroUser      *usecases.RegisterUserUseCase
	LoginUser         *usecases.LoginUseCase
	VerificationToken *usecases.VerifyTokenUseCase
}

func NewUserWeb(registroUser *usecases.RegisterUserUseCase, loginUser *usecases.LoginUseCase, verificationToken *usecases.VerifyTokenUseCase) *UserWeb {
	return &UserWeb{
		RegistroUser:      registroUser,
		LoginUser:         loginUser,
		VerificationToken: verificationToken,
	}
}
func (w *UserWeb) RegisterUserWeb(c *gin.Context) {
	var userInput dto.UserInputDTO
	err := c.ShouldBindJSON(&userInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	err = w.RegistroUser.Execute(userInput)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	token, err := w.LoginUser.Execute(dto.UserLoginInputDTO{Email: userInput.Email, Password: userInput.Password})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Token": token})
}

func (w *UserWeb) LoginUserWeb(c *gin.Context) {
	var loginInputDTO dto.UserLoginInputDTO
	err := c.ShouldBindJSON(&loginInputDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	token, err := w.LoginUser.Execute(loginInputDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Token": token})
}

func (u *UserWeb) VerificarTokenUseCase(c *gin.Context) {
	jwtToken := c.GetHeader("Token")
	if jwtToken == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Token invalid"})
		c.Abort()
	}
	ClaimsOutputDTO, err := u.VerificationToken.Execute(jwtToken)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err})
		c.Abort()
	}
	c.Set("Id", ClaimsOutputDTO.Id)
	c.Set("IsAdmin", ClaimsOutputDTO.IsAdmin)
	c.Next()
}
