package routes

import (
  "strings"
  "../controllers"
  "../models"
  "../security"
)

var ProductRoutes = Routes{
  models.Route{
    "GetProducts",
    strings.ToUpper("Get"),
    "/products",
    controllers.GetProducts,
    security.ValidateMiddleware,
  },

  models.Route{
    "CreateProduct",
    strings.ToUpper("Post"),
    "/products",
    controllers.CreateProduct,
    security.ValidateMiddleware,
  },

  models.Route{
    "GetProductById",
    strings.ToUpper("Get"),
    "/products/{id}",
    controllers.GetProductById,
    security.ValidateMiddleware,
  },

  models.Route{
    "UpdateProduct",
    strings.ToUpper("Put"),
    "/products/{id}",
    controllers.UpdateProduct,
    security.ValidateMiddleware,
  },

  models.Route{
    "DeleteProduct",
    strings.ToUpper("Delete"),
    "/products/{id}",
    controllers.DeleteProduct,
    security.ValidateMiddleware,
  },
}
