package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)


const secretKey = "secret"

func GenerateJwt(Issuer string)(string , error){
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer: Issuer,
	
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
	})

  return  claims.SignedString([]byte(secretKey))
}


func Parsejwt(cookie string)(string , error){

	token , err := jwt.ParseWithClaims(cookie , &jwt.RegisteredClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(secretKey) , nil
	} )

	if err != nil || !token.Valid {
		return  "", err
	}

	claims := token.Claims.(*jwt.RegisteredClaims)

	return  claims.Issuer ,nil
}