package models

import (
  bson "gopkg.in/mgo.v2/bson"
)

type User struct {
  Id  bson.ObjectId `bson:"_id" json:"id"`
  Email string `json:"email"`
  Password string `json:"password"`
}
