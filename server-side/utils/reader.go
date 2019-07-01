package utils

import (
  "os"
  "encoding/json"
)

func ReadJsonObjects(path string, data interface{}) (error) {
  file, err := os.Open(path)

  if err != nil {
    return err
  }

  defer file.Close()
  decoder := json.NewDecoder(file)
  err = decoder.Decode(&data)

  if err != nil {
    return err
  }

  return nil
}
