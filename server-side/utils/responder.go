package utils

import (
  "net/http"
  "encoding/json"
)

func RespondWithJson(w http.ResponseWriter, code int, payload interface{}) {
  response, err := json.Marshal(payload)

  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  w.Header().Set("Content-Type", "application/json; charset=UTF-8")
  w.WriteHeader(code)
  w.Write(response)
}
