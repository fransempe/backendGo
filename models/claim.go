package models

import (
	jwt "github.com/dgrijalva/jwt-go"
)

//Claim is a struct that process jwt
type Claim struct {
	Username string `json:"username"`

	jwt.StandardClaims
}
