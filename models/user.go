package models

import (
	"time"
	"secsys/libs"
)

// User type represents the registered user.
type User struct {
  ID string `json:"id"`
  Phone string `json:"phone"`
  Nickname *string `json:"nickname"`    //make it pointer so that it can be null
  Avator *string `json:"avator"`
  Created time.Time `json:"created"`
  Modified time.Time `json:"modified"`
  Password string `json:"password,omitempty"`
}

// CreateUser create user and return id
func CreateUser(phone string, hashedPassword string) (string, error) {
  var id string
  createUserSQL := `INSERT INTO users (phone, password) 
    VALUES($1, $2) 
    RETURNING id;`
  err := libs.Db.Get(&id, createUserSQL, phone, hashedPassword)
  return id, err
}

// FindUserByPhone find user by phone
func FindUserByPhone(phone string) (User, error) {
  var user User
  findUserSQL := `SELECT * FROM users WHERE phone=$1`
  err := libs.Db.Get(&user, findUserSQL, phone)
  return user, err
}

