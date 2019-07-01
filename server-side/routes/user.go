package routes

import (
  "strings"
  "../controllers"
  "../models"
)

var UserRoutes = Routes{
  models.Route{
    "SignUp",
    strings.ToUpper("Post"),
    "/user/signup",
    controllers.SignUp,
    nil,
  },

  models.Route{
    "SignIn",
    strings.ToUpper("Post"),
    "/user/signin",
    controllers.SignIn,
    nil,
  },

  models.Route{
    "RefreshToken",
    strings.ToUpper("Get"),
    "/user/refresh_token",
    controllers.RefreshToken,
    nil,
  },
}
