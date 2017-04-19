package db

import (
  "log"
  "github.com/jmoiron/sqlx"
  _ "github.com/lib/pq"   // init postgresql driver
  "secsys/config"
)

// Pool database connect pool struct
var Pool *sqlx.DB

// InitDb build connect pool to database
func InitDb() {
  db, err := sqlx.Open("postgres", config.DbURI)
  if err != nil {
    log.Fatalln("Database source URI error: " + err.Error())
  }
  err = db.Ping()
  if err != nil {
    log.Fatalln("Database connect error: " + err.Error())
  }
  Pool = db
  log.Println("Database connected")
}