package controllers

import (
  "net/http"

  "secsys/libs"
	"log"
)

// GetCurrentUser GET /user
func GetCurrentUser(w http.ResponseWriter, r *http.Request) {
  userID := r.Context().Value(libs.ContextKey("userid")).(string)
  log.Println("userID: ", userID)
}

// UpdateCurrentUser PATCH /user
func UpdateCurrentUser(w http.ResponseWriter, r *http.Request) {

}

// DeleteCurrentUser DELETE /user
func DeleteCurrentUser(w http.ResponseWriter, r *http.Request) {

}

// GetUser GET /users/:userID
func GetUser(w http.ResponseWriter, r *http.Request) {

}