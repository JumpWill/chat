package schema

import (
	"github.com/golang-jwt/jwt/v4"
)

type Jwt struct {
	UserId string
	Role   string
	jwt.RegisteredClaims
}
