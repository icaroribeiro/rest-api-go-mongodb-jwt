package controllers

import (
  "fmt"
  "net/http"
  "encoding/json"
  "github.com/gorilla/mux"
  "../models"
  "../utils"
  "../database"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
  var products []models.Product
  products, err := database.GetProducts();

  if err != nil {
    code := http.StatusInternalServerError
    message := err.Error()
    payload := map[string]string{"message": message}
    utils.RespondWithJson(w, code, payload)
    utils.LogError(r, code, message)
    return
  }
  
  if products == nil {
    code := http.StatusNotFound
    message := "no product is registered"
    payload := map[string]string{"message": message}
    utils.RespondWithJson(w, code, payload)
    utils.LogError(r, code, message)
    return
  }
  
  code := http.StatusOK
  payload := products
  utils.RespondWithJson(w, code, payload)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
  var product models.Product
  err := json.NewDecoder(r.Body).Decode(&product)

  if err != nil {
    code := http.StatusBadRequest
    message := "invalid request payload"
    payload := map[string]string{"message": message}
    utils.RespondWithJson(w, code, payload)
    utils.LogError(r, code, message)
    return
  }

  if product.Name == "" {
    code := http.StatusBadRequest
    message := "product name is required in the payload"
    payload := map[string]string{"message": message}
    utils.RespondWithJson(w, code, payload)
    utils.LogError(r, code, message)
    return
  }

   err = database.CreateProduct(product);

  if err != nil {
    code := http.StatusInternalServerError
    message := err.Error()
    payload := map[string]string{"message": message}
    utils.RespondWithJson(w, code, payload)
    utils.LogError(r, code, message)
    return
  }

  code := http.StatusCreated
  message := "product created successfully"
  payload := map[string]string{"message": message}
  utils.RespondWithJson(w, code, payload)
}

func GetProductById(w http.ResponseWriter, r *http.Request) {
  var product models.Product
  params := mux.Vars(r)
  product, err := database.GetProductById(params["id"]);

  if err != nil {
    code := http.StatusInternalServerError
    message := err.Error()
    payload := map[string]string{"message": message}
    utils.RespondWithJson(w, code, payload)
    utils.LogError(r, code, message)
    return
  }

  if product.Id == "" {
    code := http.StatusNotFound
    message := fmt.Sprintf("product id %s was not found", params["id"])
    payload := map[string]string{"message": message}
    utils.RespondWithJson(w, code, payload)
    utils.LogError(r, code, message)
    return
  }

  code := http.StatusOK
  payload := product
  utils.RespondWithJson(w, code, payload)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
  var product models.Product
  params := mux.Vars(r)
  product, err := database.GetProductById(params["id"]);

  if err != nil {
    code := http.StatusInternalServerError
    message := err.Error()
    payload := map[string]string{"message": message}
    utils.RespondWithJson(w, code, payload)
    utils.LogError(r, code, message)
    return
  }

  if product.Id == "" {
    code := http.StatusNotFound
    message := fmt.Sprintf("product id %s was not found", params["id"])
    payload := map[string]string{"message": message}
    utils.RespondWithJson(w, code, payload)
    utils.LogError(r, code, message)
    return
  }

  err = json.NewDecoder(r.Body).Decode(&product)

  if err != nil {
    code := http.StatusBadRequest
    message := "invalid request payload"
    payload := map[string]string{"message": message}
    utils.RespondWithJson(w, code, payload)
    utils.LogError(r, code, message)
    return
  }

   err = database.UpdateProduct(product)

  if err != nil {
    code := http.StatusInternalServerError
    message := err.Error()
    payload := map[string]string{"message": message}
    utils.RespondWithJson(w, code, payload)
    utils.LogError(r, code, message)
    return
  }

  code := http.StatusOK
  message := "product updated successfully"
  payload := map[string]string{"message": message}
  utils.RespondWithJson(w, code, payload)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
  var product models.Product
  params := mux.Vars(r)
  product, err := database.GetProductById(params["id"]);

  if err != nil {
    code := http.StatusInternalServerError
    message := err.Error()
    payload := map[string]string{"message": message}
    utils.RespondWithJson(w, code, payload)
    utils.LogError(r, code, message)
    return
  }

  if product.Id == "" {
    code := http.StatusNotFound
    message := fmt.Sprintf("product id %s was not found", params["id"])
    payload := map[string]string{"message": message}
    utils.RespondWithJson(w, code, payload)
    utils.LogError(r, code, message)
    return
  }

  err = database.DeleteProduct(product)

  if err != nil {
    code := http.StatusInternalServerError
    message := err.Error()
    payload := map[string]string{"message": message}
    utils.RespondWithJson(w, code, payload)
    utils.LogError(r, code, message)
    return
  }

  code := http.StatusOK
  message := "product deleted successfully"
  payload := map[string]string{"message": message}
  utils.RespondWithJson(w, code, payload)
}
