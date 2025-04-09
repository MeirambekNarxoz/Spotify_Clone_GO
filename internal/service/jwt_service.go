package service

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
)

type JWTService struct {
	SecretKey string
}

func NewJWTService() *JWTService {
	return &JWTService{SecretKey: "70eb82a1680c9254c9fc747223c661b7ae700db76b5a91cb964530de6853c65ec0641e5e86b31df3b5ec89882c6dee3e5a4368ad78b7558bc6d0be0e103f0c6a0c4685f828e8ade65367f8d73700821af784018dc199dbff49c47645f301f0df5003c532ed1d2adc0ce3b0a36b1e037c090d5c1296fc1c97834fb3f8908318229bb4328fe2e666f160a6481da603899ea9e9ceaf27c6f5ee6faa35cb0c0535bf489f57ff23235a3ae8a7e857bc9f2601636a75447c4db9ac0dc2caa342e1534318ae9132a9e9ab36a7a654c7d3876cdd84646c79f4d86e311a7637987262a61ecd044df6bbf7a66c58031c5a102057877b6702d28652b376a35e0e3264320f50"}
}

func (s *JWTService) ValidateToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token signing method")
		}
		return []byte(s.SecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
