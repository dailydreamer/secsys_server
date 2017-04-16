package libs

import (
  "log"
  "github.com/jmoiron/sqlx"
  _ "github.com/lib/pq"
  "secsys/config"
)

// Db database connect pool struct
var Db *sqlx.DB

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
  Db = db
  log.Println("Database connected")
}