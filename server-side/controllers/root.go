package controllers

import (
  "net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
  http.ServeFile(w, r, "..\\client-side\\index.html")
}
