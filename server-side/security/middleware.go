package security

import (
  "fmt"
  "net/http"
  jwt "github.com/dgrijalva/jwt-go"
  "../models"
  "../utils"
)

func ValidateMiddleware(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {    
    bearerToken, err := GetToken(r)

    if err != nil {
      code := http.StatusBadRequest
      message := err.Error()
      payload := map[string]string{"message": message}
      utils.RespondWithJson(w, code, payload)
      utils.LogError(r, code, message)
      return
    }

    _, err = jwt.ParseWithClaims(bearerToken, &models.Claims{}, func (token *jwt.Token) (interface{}, error) {
      if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
        return nil, fmt.Errorf("token signature failed")
      }

      return rsaPublicKey, nil
    })

    if err != nil {
      code := http.StatusInternalServerError
      message := err.Error()
      payload := map[string]string{"message": message}
      utils.RespondWithJson(w, code, payload)
      utils.LogError(r, code, message)
      return
    }

    // JsonWebToken is valid.
    next.ServeHTTP(w, r)
  })
}
