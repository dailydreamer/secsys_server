package models

import (
	"time"
	"secsys/db"
)

// Message type represents message
type Message struct {
  ID string `json:"id" db:"id"`
  UserID string `json:"userID" db:"user_id"`
  Created time.Time `json:"created" db:"created"`
  ComName string `json:"comName" db:"com_name"`
  Message string `json:"message" db:"message"`
}

// CreateMessage create message and return id
func CreateMessage(message Message) (string, error) {
  var id string
  createMessageSQL := `INSERT INTO messages (user_id, com_name, message)
    VALUES($1, $2, $3)
    RETURNING id;`
  err := db.Pool.Get(&id, createMessageSQL, message.UserID, message.ComName, message.Message)
  return id, err
}

// GetMessages return Message list
func GetMessages() ([]Message, error) {
  messages := []Message{}
  getMessagesSQL := `SELECT * FROM messages ORDER BY created DESC`
  err := db.Pool.Select(&messages, getMessagesSQL)
  return messages, err
}
