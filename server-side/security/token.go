package security

import (
  "fmt"
  "time"
  "net/http"
  "strings"
  jwt "github.com/dgrijalva/jwt-go"
  "../models"
)

const (
  // The time in minutes the token will remain valid
  tokenExpTime = 5
)

func GetToken(r *http.Request) (string, error) {
  authHeader := r.Header.Get("authorization")

  if authHeader != "" {
    bearerToken := strings.Split(authHeader, " ")

    if len(bearerToken) == 2 {
      return bearerToken[1], nil
    } else {
      return "", fmt.Errorf("token was not informed")
    }
  } else {
    return "", fmt.Errorf("authorization header was not informed")
  }
}

func CreateAndSignToken(email string) (string, error) {
  claims := models.Claims{
    Email: email,
    StandardClaims: jwt.StandardClaims{
      ExpiresAt: time.Now().Add(time.Minute * tokenExpTime).Unix(),
    },
  }

  // Create a signed token based on the RSA 256 signing method.
  token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
  signedToken, err := token.SignedString(rsaPrivateKey)

  if err != nil {
    return "", err
  }

  return signedToken, nil
}

func RefreshToken(bearerToken string) (string, error) {
  claims := models.Claims{}

  _, err := jwt.ParseWithClaims(bearerToken, &claims, func (token *jwt.Token) (interface{}, error) {
    if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
      return nil, fmt.Errorf("token signature failed")
    }
    return rsaPublicKey, nil
  })

  if err != nil {
    v, _ := err.(*jwt.ValidationError)

    if v.Errors != jwt.ValidationErrorExpired {
      return "", err
    }
  }

  // We ensure that a new token is not issued until enough time has elapsed
  // In this case, a new token will only be issued if the old token is within
  // 30 seconds of expiry. Otherwise, return a bad request status
  if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30 * time.Second {
    return "", fmt.Errorf("token is not within 30 seconds of expiry")
  }

  signedToken, err := CreateAndSignToken(claims.Email)

  if err != nil {
    return "", err
  }

  return signedToken, nil
}
