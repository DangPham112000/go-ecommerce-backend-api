package auth

import (
	"time"

	"github.com/DangPham112000/go-ecommerce-backend-api/global"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type PayloadClaims struct {
	jwt.StandardClaims
}

func GenTokenJWT(payload jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return token.SignedString([]byte(global.Config.JWT.API_SECRET))
}

func CreateToken(uuidToken string) (string, error) {
	// Set time expiration
	timeExp := global.Config.JWT.JWT_EXPIRATION
	if timeExp == "" {
		timeExp = "1h"
	}
	expiration, err := time.ParseDuration(timeExp)
	if err != nil {
		return "", err
	}

	now := time.Now()
	expiresAt := now.Add(expiration)
	return GenTokenJWT(&PayloadClaims{
		StandardClaims: jwt.StandardClaims{
			Id:        uuid.New().String(),
			ExpiresAt: expiresAt.Unix(),
			IssuedAt:  now.Unix(),
			Issuer:    "shopgolang",
			Subject:   uuidToken,
		},
	})
}
