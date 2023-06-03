package schema

import (
	"github.com/golang-jwt/jwt/v4"
)

type Token struct {
	AccessToken  string
	RefreshToken string
}

type Jwt struct {
	UserId string
	Role   string
	jwt.RegisteredClaims
}
