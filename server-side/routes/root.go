package routes

import (
  "strings"
  "../controllers"
  "../models"
)

var RootRoute = Routes{
  models.Route{
    "Index",
    strings.ToUpper("Get"),
    "/",
    controllers.Index,
    nil,
  },
}
