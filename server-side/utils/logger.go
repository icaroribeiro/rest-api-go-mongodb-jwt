package utils

import (
  "fmt"
  "log"
  "net/http"
)

func LogError(r *http.Request, code int, message string) {
  log.Printf(
    "%s %s %s %s",
    r.Method,
    r.RequestURI,
    fmt.Sprintf("-> code: %d", code),
    fmt.Sprintf("message: %s", message),
  )
}
