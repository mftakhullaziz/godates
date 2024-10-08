package jsonwebtoken

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"godating-dealls/internal/common"
	"time"
)

var jwtSecret = []byte(common.StringEncoder("key-app-godating-dealls"))

type JWTTokenClaims struct {
	UserId    int64  `json:"user_id"`
	AccountId int64  `json:"account_id"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	jwt.RegisteredClaims
}

func GenerateJWTToken(userId int64, accountId int64, email string) (string, error) {
	expireAt := time.Now().Add(24 * time.Hour)
	claims := JWTTokenClaims{
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

//// ExtractToken to extract and decode the token
//func ExtractToken(tokenString string) (*JWTTokenClaims, error) {
//	token, err := jwt.ParseWithClaims(tokenString, &JWTTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
//		// Typically, the key used to sign the token should be passed here.
//		// For simplicity, we're using a hardcoded secret. In a real-world scenario, ensure to use a proper method to handle the secret.
//		return []byte("your-256-bit-secret"), nil
//	})
//
//	if claims, ok := token.Claims.(*JWTTokenClaims); ok && token.Valid {
//		return claims, nil
//	} else {
//		return nil, err
//	}
//}
