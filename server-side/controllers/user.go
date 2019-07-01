package controllers

import (
  "net/http"
  "encoding/json"
  bcrypt "golang.org/x/crypto/bcrypt"
  "../database"
  "../models"
  "../security"
  "../utils"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
  var user models.User
  err := json.NewDecoder(r.Body).Decode(&user)

  if err != nil {
    code := http.StatusBadRequest
    message := "invalid request payload"
    payload := map[string]string{"message": message}
    utils.RespondWithJson(w, code, payload)
    utils.LogError(r, code, message)
    return
  }

  if user.Email == "" {
    code := http.StatusBadRequest
    message := "user email is required in the payload"
    payload := map[string]string{"message": message}
    utils.RespondWithJson(w, code, payload)
    utils.LogError(r, code, message)
    return
  }

  if user.Password == "" {
    code := http.StatusBadRequest
    message := "user password is required in the payload"
    payload := map[string]string{"message": message}
    utils.RespondWithJson(w, code, payload)
    utils.LogError(r, code, message)
    return
  }

  var userAux models.User
  userAux, err = database.GetUserByEmail(user.Email)

  if err != nil {
    code := http.StatusInternalServerError
    message := err.Error()
    payload := map[string]string{"message": message}
    utils.RespondWithJson(w, code, payload)
    utils.LogError(r, code, message)
    return
  }
  
  if user.Email == userAux.Email {
    code := http.StatusBadRequest
    message := "user email already exists, then create a new one"
    payload := map[string]string{"message": message}
    utils.RespondWithJson(w, code, payload)
    utils.LogError(r, code, message)
    return
  }

  // Hash the password using the bcrypt algorithm and a default cost of hashing.
  hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

  if err != nil {
    code := http.StatusInternalServerError
    message := err.Error()
    payload := map[string]string{"message": message}
    utils.RespondWithJson(w, code, payload)
    utils.LogError(r, code, message)
    return
  }

  user.Password = string(hashedPassword)
  err = database.CreateUser(user)

  if err != nil {
    code := http.StatusInternalServerError
    message := err.Error()
    payload := map[string]string{"message": message}
    utils.RespondWithJson(w, code, payload)
    utils.LogError(r, code, message)
    return
  }

  code := http.StatusCreated
  message := "user created successfully"
  payload := map[string]string{"message": message}
  utils.RespondWithJson(w, code, payload)
}

func SignIn(w http.ResponseWriter, r *http.Request) {
  var user models.User
  err := json.NewDecoder(r.Body).Decode(&user)

  if err != nil {
    code := http.StatusBadRequest
    message := "invalid request payload"
    payload := map[string]string{"message": message}
    utils.RespondWithJson(w, code, payload)
    utils.LogError(r, code, message)
    return
  }

  if user.Email == "" {
    code := http.StatusBadRequest
    message := "user email is required in the payload"
    payload := map[string]string{"message": message}
    utils.RespondWithJson(w, code, payload)
    utils.LogError(r, code, message)
    return
  }

  if user.Password == "" {
    code := http.StatusBadRequest
    message := "user password is required in the payload"
    payload := map[string]string{"message": message}
    utils.RespondWithJson(w, code, payload)
    utils.LogError(r, code, message)
    return
  }

  var userAux models.User
  userAux, err = database.GetUserByEmail(user.Email)

  if err != nil {
    code := http.StatusInternalServerError
    message := err.Error()
    payload := map[string]string{"message": message}
    utils.RespondWithJson(w, code, payload)
    utils.LogError(r, code, message)
    return
  }
  
  if userAux.Email == "" {
    code := http.StatusBadRequest
    message := "user credentials are not registered"
    payload := map[string]string{"message": message}
    utils.RespondWithJson(w, code, payload)
    utils.LogError(r, code, message)
    return
  }

  // Compare the stored hashed password from database
  // with the password that was received in the payload
  if err = bcrypt.CompareHashAndPassword([]byte(userAux.Password), []byte(user.Password)); err != nil {
    code := http.StatusUnauthorized
    message := "user password is invalid"
    payload := map[string]string{"message": message}
    utils.RespondWithJson(w, code, payload)
    utils.LogError(r, code, message)
    return
  }

  token, err := security.CreateAndSignToken(user.Email)

  if err != nil {
    code := http.StatusInternalServerError
    message := err.Error()
    payload := map[string]string{"message": message}
    utils.RespondWithJson(w, code, payload)
    utils.LogError(r, code, message)
    return
  }

  code := http.StatusOK
  message := token
  payload := map[string]string{"token": message}
  utils.RespondWithJson(w, code, payload)
}

func RefreshToken(w http.ResponseWriter, r *http.Request) {
  bearerToken, err := security.GetToken(r)

  if err != nil {
    code := http.StatusBadRequest
    message := err.Error()
    payload := map[string]string{"message": message}
    utils.RespondWithJson(w, code, payload)
    utils.LogError(r, code, message)
    return
  }

  token, err := security.RefreshToken(bearerToken)

  if err != nil {
    code := http.StatusInternalServerError
    message := err.Error()
    payload := map[string]string{"message": message}
    utils.RespondWithJson(w, code, payload)
    utils.LogError(r, code, message)
    return
  }

  code := http.StatusOK
  message := token
  payload := map[string]string{"token": message}
  utils.RespondWithJson(w, code, payload)
}
