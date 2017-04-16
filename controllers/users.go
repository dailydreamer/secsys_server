package controllers

import (
  "net/http"

  "secsys/libs"
	"log"
)

// GetUser GET /user
func GetUser(w http.ResponseWriter, r *http.Request) {
  userID := r.Context().Value(libs.ContextKey("userid")).(string)
  log.Println("userID: ", userID)
}

// UpdateUser PATCH /user
func UpdateUser() {

}

// DeleteUser DELETE /user
func DeleteUser() {

}

// GetUserByIDController GET /users/:id
func GetUserByIDController(w http.ResponseWriter, r *http.Request) {

}