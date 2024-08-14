package usecases

import (
	dto "GastroReserve/internal/DTO"
	"fmt"
	"os"

	"github.com/golang-jwt/jwt"
)

type VerificaTokenUseCase struct{}

func NewVerifyTokenUseCase() *VerificaTokenUseCase {
	return &VerificaTokenUseCase{}
}

func (u *VerificaTokenUseCase) Execute(tokenString string) (*dto.ClaimsOutputDTO, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("token inválido")
		}
		return []byte(os.Getenv("SECRETKEY")), nil
	})

	if err != nil {
		return &dto.ClaimsOutputDTO{}, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return &dto.ClaimsOutputDTO{}, fmt.Errorf("invalid token or claims")
	}
	id, ok := claims["Id"].(string)
	if !ok {
		return &dto.ClaimsOutputDTO{}, fmt.Errorf("token inválido")
	}
	isAdmin, ok := claims["IsAdmin"].(bool)
	if !ok {
		return &dto.ClaimsOutputDTO{}, fmt.Errorf("token inválido")
	}
	return &dto.ClaimsOutputDTO{
		Id:      id,
		IsAdmin: isAdmin,
	}, nil
}
