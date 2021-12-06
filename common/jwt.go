package common

import (
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/pbatista27/portafolio-serve/models"
)

type Token struct{}

func NewToken() *Token {
	return &Token{}
}

func (t Token) Generar(u models.User) (string, error) {

	semilla := []byte(os.Getenv("SECRECT_JWT_SEED"))

	payload := jwt.MapClaims{
		"_id":  u.Id.Hex(),
		"name": u.Name,
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	tokenstr, err := token.SignedString(semilla)

	if err != nil {
		return tokenstr, err
	}

	return tokenstr, nil

}
