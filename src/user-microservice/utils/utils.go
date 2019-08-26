package utils

import (
	"time"

	jwt_lib "github.com/dgrijalva/jwt-go"
	"github.com/gillepool/MovieBackend/src/movie-microservice/utils"
)

//SdtClaim struct
type SdtClaim struct {
	Name string `json:"name"`
	Role string `json:"role"`
	jwt_lib.StandardClaims
}

//Jwt struct
type Utils struct {
}

//GenerateJWT generates token
func (u *Utils) GenerateJWT(name string, role string) (string, error) {
	claims := SdtClaim{
		name,
		role,
		jwt_lib.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
			Issuer:    utils.Config.Issuer,
		},
	}

	token := jwt_lib.NewWithClaims(jwt_lib.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(utils.Config.JwtSecretPassword))

	return tokenString, err
}
