package database

import (
  mgo "gopkg.in/mgo.v2"
  "../utils"
)

var mongodb *mgo.Database

const (
  path = "./database/cfg.json"
)

type Configuration struct {
  Ip   string
  Database string
}

func EstablishConnection() (error) {
  cfg := Configuration{}

  err := utils.ReadJsonObjects(path, &cfg)

  if err != nil {
    return err
  }

  var session *mgo.Session
  session, err = mgo.Dial(cfg.Ip)
  
  if err != nil {
    return err
  }

	session.SetMode(mgo.Monotonic, true)
  mongodb = session.DB(cfg.Database)

  return nil
}
