package models

import (
	"time"
	"secsys/db"
)

// User type represents the registered user.
type User struct {
  ID string `json:"id"`
  IsAdmin bool `json:"isAdmin"`
  Phone string `json:"phone"`
  Nickname *string `json:"nickname"`    //make it pointer so that it can be null
  Avator *string `json:"avator"`
  Created time.Time `json:"created"`
  Modified time.Time `json:"modified"`
  Password string `json:"password,omitempty"`
}

// CreateUser create user and return id
func CreateUser(phone string, hashedPassword string, isAdmin bool) (string, error) {
  var id string
  createUserSQL := `INSERT INTO users (phone, password, is_admin) 
    VALUES($1, $2, $3) 
    RETURNING id;`
  err := db.Pool.Get(&id, createUserSQL, phone, hashedPassword, isAdmin)
  return id, err
}

// FindUserByPhone find user by phone
func FindUserByPhone(phone string) (User, error) {
  var user User
  findUserSQL := `SELECT * FROM users WHERE phone=$1`
  err := db.Pool.Get(&user, findUserSQL, phone)
  return user, err
}

