package helper

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"time"

	"github.com/Haroxa/Integrated_documentation/common"
)

type JwtClaim struct {
	UserId int
	jwt.RegisteredClaims
}

// this key is the most dangerous!!!! MUST BE DIFFICULT TO GUESS
var myKey = []byte("fahkdslfhakldsjfklasdk321084710jfd")

func CreatToken(UserId int) (string, error) {
	claim := JwtClaim{
		UserId: UserId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().In(common.ChinaTime).Add(168 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now().In(common.ChinaTime)),
		}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, err := token.SignedString(myKey)
	if err != nil {
		return "", nil
	}
	return tokenString, nil
}

func VerifyToken(token string) (int, error) {
	tempToken, err := jwt.ParseWithClaims(token, &JwtClaim{}, func(token *jwt.Token) (interface{}, error) {
		return myKey, nil
	})
	if err != nil {
		return -1, err
	}
	claims, ok := tempToken.Claims.(*JwtClaim)
	if !ok {
		return -1, errors.New("claims error")
	}
	if err := tempToken.Claims.Valid(); err != nil {
		return -1, err
	}
	return claims.UserId, nil
}
