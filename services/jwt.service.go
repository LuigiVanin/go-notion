package services

import (
	"fmt"
	"main/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AuthPayload struct {
	UserId uint  `json:"userId"`
	Time   int64 `json:"time"`
	jwt.RegisteredClaims
}

func CreateJwtToken(payload AuthPayload) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, AuthPayload{
		UserId: payload.UserId,
		Time:   payload.Time,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Unix(payload.Time+3600, 0)),
			IssuedAt:  jwt.NewNumericDate(time.Unix(payload.Time, 0)),
		},
	})
	tokenString, err := token.SignedString([]byte(config.GlobalEnv.JwtKey))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ParseJwtToken(token string) (*AuthPayload, error) {

	parsedToken, err := jwt.ParseWithClaims(token, &AuthPayload{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.GlobalEnv.JwtKey), nil
	})

	if err != nil {
		return nil, err
	}

	// Check if the token is valid and has valid claims
	if claims, ok := parsedToken.Claims.(*AuthPayload); ok && parsedToken.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("Token is not valid")
}
