package utils

import (
	"chat/global"

	"chat/schema"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func Secret(secret []byte) jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		return secret, nil // 这是我的secret
	}
}

func JwtEncoder(user_id string, role string) (string, error) {
	secret := global.Settings.Jwt.Secret
	claim := schema.Jwt{
		// TODO Payload
		UserId: user_id,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(global.Settings.Jwt.Duration) * time.Hour)), // 过期时间3小时
			IssuedAt:  jwt.NewNumericDate(time.Now()),                                                              // 签发时间
			NotBefore: jwt.NewNumericDate(time.Now()),                                                              // 生效时间
		}}
	cli := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	token, err := cli.SignedString([]byte(secret))
	Error(err, "Sign JWT token error!")
	return token, nil

}

func JwtDecoder(token string) (*schema.Jwt, error) {

	data, err := jwt.ParseWithClaims(token, &schema.Jwt{}, Secret([]byte("string")))
	info := ""
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				err = errors.New("invalid token")
				info = "Invalid token"
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				err = errors.New("token has been expired")
				info = "Token has been expired"
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				err = errors.New("token has not")
				info = "Token not active yet"
			} else {
				err = errors.New("couldn't parse token")
				info = "Couldn't parse token"
			}
		}
	}

	Error(err, info)
	claims, _ := data.Claims.(*schema.Jwt)
	return claims, nil
}
