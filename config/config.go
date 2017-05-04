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
  flag.StringVar(&DbURI, "db_uri", "postgres://zsgogdpabujfvb:b8ff515b7eb9becefb8455d1f6c890e0cc66246487f1b2ae35535fb4f96f27d2@ec2-54-235-119-27.compute-1.amazonaws.com:5432/d9ijitr6tas3b", "uri of database")
  flag.Parse()
  BasicURL = ServerURL + VersionURL
  log.Println("Config loaded")
}