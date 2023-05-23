package models

import (
	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	name string `json:"name"`
	jwt.StandardClaims
}
