package models

import (
	"time"
	"secsys/db"
)

// Log type represents log
type Log struct {
  ID string `json:"id" db:"id"`
  UserID string `json:"userID" db:"user_id"`
  Created time.Time `json:"created" db:"created"`
  ComName string `json:"comName" db:"com_name"`
  IP string `json:"ip" db:"ip"`
  Address string `json:"address" db:"address"`
  Status string `json:"status" db:"status"`
}

// CreateLog create log and return id
func CreateLog(log Log) (string, error) {
  var id string
  createLogSQL := `INSERT INTO logs (user_id, com_name, ip, address, status)
    VALUES($1, $2, $3, $4, $5)
    RETURNING id;`
  err := db.Pool.Get(&id, createLogSQL, log.UserID, log.ComName, log.IP, log.Address, log.Status)
  return id, err
}

// GetLogs return Log list
func GetLogs() ([]Log, error) {
  logs := []Log{}
  getLogsSQL := `SELECT * FROM logs ORDER BY created DESC`
  err := db.Pool.Select(&logs, getLogsSQL)
  return logs, err
}
