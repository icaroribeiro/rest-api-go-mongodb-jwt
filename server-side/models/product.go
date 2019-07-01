package models

import (
  bson "gopkg.in/mgo.v2/bson"
)

type Product struct {
  Id  bson.ObjectId `bson:"_id" json:"id"`
  Name string `json:"name"`
}
