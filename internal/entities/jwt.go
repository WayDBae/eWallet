package entities

import "github.com/golang-jwt/jwt/v5"

type CustomClaims struct {
	Name string `json:"name,omitempty"`
	jwt.RegisteredClaims // содержит exp, iat и т.д.
}
