package service

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
)

type JWTService struct {
	SecretKey string
}

func NewJWTService() *JWTService {
	return &JWTService{SecretKey: "cfcf0b5ff953c9a8801c1014de6aee167b2ebf6cc915f6e1a7a8ee5b1accb223f6b64350e73691817a94be36a5fcacd177b49ffc8b0ba4766a741302212766ce139bb8833f726a89c65c8bfb3f59f056a2a45aa3ce5b4fe8703e9e816b8d2f521df28f2b837c01cb8d7d11cf8d71e3acc528ea9facb14bc3b3707209387fd6fe1e564322b0e9985e2e8d846d26e9d6a49b81f61973251e9e9fa6b3b758b4055900049666de6f494e84e102ead2fd9ec9a0667b004de58d4fb6d871448def174a0255de91c729dbf49b1fa9856268f32224f9680674834f17b8ad09c5400159ce0234a1ea54189c33b24751fa05d64d43b28b9b462724b045d609ba8ad07fad9f"}
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
