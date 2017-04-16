package config

import (
  "log"
  "flag"
)

var (
  Port string
  JWTSecret string
  ServerURL string
  VersionURL string
  BasicURL string
  DbURI string
)

// InitConfig init the config parameters.
func InitConfig() {
  flag.StringVar(&Port, "port", "3000", "port service listen at")
  flag.StringVar(&JWTSecret, "jwt_secret", "secsys_secret", "secret of jwt hmac algorithm")
  flag.StringVar(&ServerURL, "server_url", "https://api.secsys.net", "server address")
  flag.StringVar(&VersionURL, "version_url", "/v1", "url of api version")
  flag.StringVar(&DbURI, "db_uri", "postgres://bzjgtwxllnppiq:31a68d8748d99ded0a1d42f95a6c28eef56f0cf5fd892250f1905f99bde7ce95@ec2-54-235-90-107.compute-1.amazonaws.com:5432/d5imq2f3o1cs7a", "uri of database")
  flag.Parse()
  BasicURL = ServerURL + VersionURL
  log.Println("Config loaded")
}