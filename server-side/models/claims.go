package models

import (
  jwt "github.com/dgrijalva/jwt-go"
)

type Claims struct {
  Email string `json:"email"`
  jwt.StandardClaims
}
