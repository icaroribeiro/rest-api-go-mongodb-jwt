package models

import (
  "net/http"
  mux "github.com/gorilla/mux"
)

type Route struct {
  Name string
  Method string
  Pattern string
  HandlerFunc http.HandlerFunc
  MiddlewareFunc mux.MiddlewareFunc
}
