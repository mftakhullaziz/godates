package jsonwebtoken

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var jwtSecret = []byte("app-godating-dealls")

type JWTTokenClaims struct {
	UserId    int64  `json:"user_id"`
	AccountId int64  `json:"account_id"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	jwt.RegisteredClaims
}

func GenerateJWTToken(userId int64, accountId int64, email string, claims JWTTokenClaims) (string, error) {
	expireAt := time.Now().Add(24 * time.Hour)
	claims = JWTTokenClaims{
		UserId:    userId,
		AccountId: accountId,
		Email:     email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireAt),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func VerifyJWTToken(accessToken string) (*JWTTokenClaims, error) {
	token, err := jwt.ParseWithClaims(accessToken, &JWTTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*JWTTokenClaims); ok && token.Valid {
		// Check if the token is expired
		if claims.ExpiresAt.Time.Before(time.Now()) {
			return nil, fmt.Errorf("token has expired")
		}
		return claims, nil
	}
	return nil, fmt.Errorf("invalid token")
}
