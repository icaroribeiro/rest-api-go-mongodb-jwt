package security

import (
  "log"
  "io/ioutil"
  "crypto/rsa"
  jwt "github.com/dgrijalva/jwt-go"
)

const (
  // Create a RSA public key.
  // $ openssl rsa -in privatekey.rsa -pubout > publickey.rsa.pub
  publicKeyPath = "./security/publickey.rsa.pub"

  // Create a RSA private key.
  // $ openssl genrsa -out privatekey.rsa 1024
  privateKeyPath = "./security/privatekey.rsa"
)

var (
  rsaPublicKey *rsa.PublicKey 
  rsaPrivateKey *rsa.PrivateKey 
)

func InitializeKeys() (error) {
  publicKey, err := ioutil.ReadFile(publicKeyPath)

  if err != nil {
    log.Fatal("Error while reading the RSA public key file")
    return err
  }

  // Convert bytes into the RSA public key.
  rsaPublicKey, err = jwt.ParseRSAPublicKeyFromPEM(publicKey)

  if err != nil {
    log.Fatal("Error while parsing the RSA public key file")
    return err
  }

  privateKey, err := ioutil.ReadFile(privateKeyPath)

  if err != nil {
    log.Fatal("Error while reading the RSA private key file")
    return err
  }

  // Convert bytes into the RSA private key.
  rsaPrivateKey, err = jwt.ParseRSAPrivateKeyFromPEM(privateKey)

  if err != nil {
    log.Fatal("Error while parsing the RSA private key file")
    return err
  }

  return nil
}
