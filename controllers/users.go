package controllers

import (
  "net/http"

  "secsys/libs"
	"log"
)

// GetUsers GET /users
func GetUsers(w http.ResponseWriter, r *http.Request) {
  userID := r.Context().Value(libs.ContextKey("userid")).(string)
  log.Println("userID: ", userID)
}

// CreateUser POST /users
func CreateUser(w http.ResponseWriter, r *http.Request) {

}

// GetUser GET /users/:userID
func GetUser(w http.ResponseWriter, r *http.Request) {

}

// UpdateUser PATCH /users/:userID
func UpdateUser(w http.ResponseWriter, r *http.Request) {

}

// DeleteUser DELETE /users/:userID
func DeleteUser(w http.ResponseWriter, r *http.Request) {

}