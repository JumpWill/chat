package schema

import (
	"github.com/golang-jwt/jwt/v4"
)

type Payload struct {
	Name   string // 名称
	UserID string // uuid
	Role   string //

}

type Token struct {
	AccessToken  string
	RefreshToken string
}

type Jwt struct {
	Payload
	jwt.RegisteredClaims
}
