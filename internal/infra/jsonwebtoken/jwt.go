package jsonwebtoken

type Jwt interface {
	GenerateToken(userID string) (string, error)
	ValidateToken(tokenString string) (string, error)
}
