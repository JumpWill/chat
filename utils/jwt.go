package utils

import (
	global "chat/global"
	schema "chat/schema"

	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func Secret(secret []byte) jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		return secret, nil // 这是我的secret
	}
}

func JwtEncoder(payload schema.Payload) (string, error) {
	claim := schema.Jwt{
		Payload: payload,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(global.Settings.Jwt.Duration) * time.Hour)), // 过期时间3小时
			IssuedAt:  jwt.NewNumericDate(time.Now()),                                                              // 签发时间
			NotBefore: jwt.NewNumericDate(time.Now()),                                                              // 生效时间
		}}
	cli := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	token, err := cli.SignedString([]byte(global.Settings.Jwt.Secret))
	Error(err, "Sign JWT token error!")
	return token, nil

}

func JwtDecoder(token string) (*schema.Jwt, error) {

	data, err := jwt.ParseWithClaims(token, &schema.Jwt{}, Secret([]byte(global.Settings.Jwt.Secret)))

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				err = errors.New("Token非法认证")
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				err = errors.New("Token已经过期")
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				err = errors.New("Token未被激活")
			} else {
				err = errors.New("Token无法解析")
			}
		}
		return nil, err
	}
	claims, _ := data.Claims.(*schema.Jwt)
	return claims, nil
}
