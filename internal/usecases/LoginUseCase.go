package usecases

import (
	dto "GastroReserve/internal/DTO"
	"GastroReserve/internal/infra/repositories"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type LoginUseCase struct {
	Repository repositories.IUserRepositoryMySql
}

func NewLoginUseCase(repository repositories.IUserRepositoryMySql) *LoginUseCase {
	return &LoginUseCase{Repository: repository}
}
func (u *LoginUseCase) Execute(UserLoginInputDTO dto.UserLoginInputDTO) (string, error) {
	user, err := u.Repository.GetUserPerEmail(UserLoginInputDTO.Email)
	if err != nil {
		return "", err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(UserLoginInputDTO.Password))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"Id":      user.Id,
		"IsAdmin": user.IsAdmin,
		"exp":     time.Now().Add(time.Hour * 1).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRETKEY")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
