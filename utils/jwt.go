package utils

import (
	"chat/schema"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

func Secret(secret []byte) jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		return secret, nil // 这是我的secret
	}
}

func JwtEncoder(user_id string, role string) (string, error) {
	// TODO 整理
	secret := "string"
	claim := schema.Jwt{
		UserId: user_id,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(3 * time.Hour * time.Duration(1))), // 过期时间3小时
			IssuedAt:  jwt.NewNumericDate(time.Now()),                                       // 签发时间
			NotBefore: jwt.NewNumericDate(time.Now()),                                       // 生效时间
		}}
	cli := jwt.NewWithClaims(jwt.SigningMethodHS256, claim) // 使用HS256算法
	token, err := cli.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return token, nil

}

func JwtDecoder(token string) (*schema.Jwt, error) {

	token2, err := jwt.ParseWithClaims(token, &schema.Jwt{}, Secret([]byte("string")))
	// TODO 优化
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errors.New("invalid token")
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, errors.New("token has been expired")
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, errors.New("token has not")
			} else {
				return nil, errors.New("couldn't parse token")
			}
		}
	}
	claims, _ := token2.Claims.(*schema.Jwt)
	return claims, nil
}
