package routes

import (
  "net/http"
  mux "github.com/gorilla/mux"
  "../models"
)

type Routes []models.Route

func CreateRouter() *mux.Router {
  router := mux.NewRouter().StrictSlash(true)

  routesList := []Routes {
    RootRoute,
    UserRoutes,
    ProductRoutes,
  }

  for _, routes := range routesList {
    for _, route := range routes {
      var handler http.Handler

      handler = route.HandlerFunc

      if route.MiddlewareFunc != nil {
        router.
          Methods(route.Method).
          Path(route.Pattern).
          Name(route.Name).
          Handler(route.MiddlewareFunc(handler))
      } else {
        router.
          Methods(route.Method).
          Path(route.Pattern).
          Name(route.Name).
          Handler(handler)
      }
    }
  }

  return router
}
