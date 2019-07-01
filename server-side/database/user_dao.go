package database

import (
  mgo "gopkg.in/mgo.v2"
  bson "gopkg.in/mgo.v2/bson"
  "../models"
)

func CreateUser(user models.User) (error) {
  user.Id = bson.NewObjectId()
  err := mongodb.C("users").Insert(&user)

  return err
}

func GetUserByEmail(email string) (models.User, error) {
  var user models.User
  err := mongodb.C("users").Find(bson.M{"email": email}).One(&user)

  if err != nil && err != mgo.ErrNotFound {
    return user, err
  }

  return user, nil
}
