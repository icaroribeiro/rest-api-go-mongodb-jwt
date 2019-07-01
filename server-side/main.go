package main

import (
  "log"
  "time"
  "net/http"
  cors "github.com/rs/cors"
  "./database"
  "./routes"
  "./security"
)

func init() {
  err := security.InitializeKeys()

  if err != nil {
    log.Fatal("Error while creating authenticaton keys:", err)
  }

  err = database.EstablishConnection()

  if err != nil {
    log.Fatal("Error while establishing database connection:", err)
  }
}

func main() {
  router := routes.CreateRouter()

  // It enables CORS (Cross Origin Resource Sharing) so that
  // an API can be accessible by JavaScript in-browser client-side code.
  corsOpts := cors.New(cors.Options{
    AllowedHeaders: []string{"*"},
    AllowedMethods: []string{
      http.MethodGet,
      http.MethodPost,
      http.MethodPut,
      http.MethodPatch,
      http.MethodDelete,
      http.MethodOptions,
      http.MethodHead,
    },
  })

  server := &http.Server{
    Addr: "127.0.0.1:8080",
    Handler: corsOpts.Handler(router),
    ReadTimeout: 30 * time.Second,
    WriteTimeout: 30 * time.Second,
  }

  log.Println("Starting server...")

  go func() {
    for {
      time.Sleep(time.Second)

      log.Println("Checking if server started...")
      
      r, err := http.Get("http://127.0.0.1:8080")
      
      if err != nil {
        log.Println("Error while starting the server:", err)
        continue
      }

      r.Body.Close()

      if r.StatusCode != http.StatusOK {
        log.Println("Error while starting the server:", r.StatusCode)
        continue
      }

      break
    }
    log.Println("Server up and running successfully!")
  }()

  err := server.ListenAndServe()

  if err != nil {
    log.Fatal("Error while starting the server:", err)
  }
}
