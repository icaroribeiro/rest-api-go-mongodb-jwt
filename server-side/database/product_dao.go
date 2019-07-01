package database

import (
  mgo "gopkg.in/mgo.v2"
  bson "gopkg.in/mgo.v2/bson"
  "../models"
)

func GetProducts() ([]models.Product, error) {
  var products []models.Product
  err := mongodb.C("products").Find(nil).All(&products)

  return products, err
}

func CreateProduct(product models.Product) (error) {
  product.Id = bson.NewObjectId()
  err := mongodb.C("products").Insert(&product)

  return err
}

func GetProductById(id string) (models.Product, error) {
  var product models.Product

  if bson.IsObjectIdHex(id) == true {
    err := mongodb.C("products").FindId(bson.ObjectIdHex(id)).One(&product)

    if err != nil && err != mgo.ErrNotFound {
      return product, err
    }
  }

  return product, nil
}

func UpdateProduct(product models.Product) (error) {
  err := mongodb.C("products").UpdateId(product.Id, &product)

  return err
}

func DeleteProduct(product models.Product) (error) {
  err := mongodb.C("products").Remove(&product)

  return err
}
