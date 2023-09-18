package services

import (
	"time"

	"github.com/NithishNithi/GoTask/constants"
	"github.com/dgrijalva/jwt-go"
)

func CreateToken(customerid,email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,"customerid":customerid,
		"exp": time.Now().Add(time.Hour * 1).Unix(),
	})

	tokenString, err := token.SignedString([]byte(constants.SecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
