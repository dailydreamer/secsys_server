package controllers

import (
  "net/http"

  "secsys/libs"
	"log"
)

// GetUsers GET /users
func GetUsers(w http.ResponseWriter, r *http.Request) {

}

// CreateUser POST /users
func CreateUser(w http.ResponseWriter, r *http.Request) {

}

// GetUser GET /users/:userID
func GetUser(w http.ResponseWriter, r *http.Request) {
  userClaims := r.Context().Value(libs.ContextKey("userClaims")).(*libs.UserClaims)
  log.Println("userID: ", userClaims.ID)
  w.Write([]byte(userClaims.ID))
}

// UpdateUser PATCH /users/:userID
func UpdateUser(w http.ResponseWriter, r *http.Request) {

}

// DeleteUser DELETE /users/:userID
func DeleteUser(w http.ResponseWriter, r *http.Request) {

}

// GetUserContracts GET /user/:userID/contracts
func GetUserContracts(w http.ResponseWriter, r *http.Request) {

}

// CreateUserContract POST /user/:userID/contracts
func CreateUserContract(w http.ResponseWriter, r *http.Request) {

}

// GetUserContract GET /user/:userID/contracts/:contractID
func GetUserContract(w http.ResponseWriter, r *http.Request) {

}


// UpdateUserContract PATCH /user/:userID/contracts/:contractID
func UpdateUserContract(w http.ResponseWriter, r *http.Request) {

}

// DeleteUserContract DELETE /user/:userID/contracts/:contractID
func DeleteUserContract(w http.ResponseWriter, r *http.Request) {

}

// GetUserScores GET /users/:userID/scores
func GetUserScores(w http.ResponseWriter, r *http.Request) {

}