package controllers

import (
  "net/http"

  "secsys/libs"
	"log"
  "secsys/models"
)

// GetUsers GET /users
func GetUsers(w http.ResponseWriter, r *http.Request) {
  companyBasicList, err := models.GetUsers()
  if err != nil {
    libs.ResponseError(w, r, "Error on get users: " + err.Error(), http.StatusInternalServerError)
  }
  libs.ResponseJSON(w, r, companyBasicList)
}

// CreateUser POST /users
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		libs.ResponseError(w, r, "Error on parse json: " + err.Error(), http.StatusBadRequest)
		return
	}
	// check required field
	if user.Phone == "" {
		libs.ResponseError(w, r, "Field phone required", http.StatusUnprocessableEntity)
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		libs.ResponseError(w, r, "Error on encrypt password: " + err.Error(), http.StatusInternalServerError)
		return
	}
	user.ID, err = models.CreateUser(user, string(hashedPassword), false)
	if err != nil {
		libs.ResponseError(w, r, "Error on create user: " + err.Error(), http.StatusInternalServerError)
		return
	}
	
  libs.ResponseSuccess(w, r)
}

// GetUser GET /users/:userID
func GetUser(w http.ResponseWriter, r *http.Request) {
  userID := chi.URLParam(r, "userID")
  user, err := models.GetUserByID(userID)
  if err != nil {
    libs.ResponseError(w, r, "Error on get user: " + err.Error(), http.StatusInternalServerError)
  }
  libs.ResponseJSON(w, r, user)
}

// UpdateUser PUT /users/:userID
func UpdateUser(w http.ResponseWriter, r *http.Request) {
  userID := chi.URLParam(r, "userID")
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		libs.ResponseError(w, r, "Error on parse json: " + err.Error(), http.StatusBadRequest)
		return
	}
  err := models.UpdateUser(user)
  if err != nil {
		libs.ResponseError(w, r, "Error on update user: " + err.Error(), http.StatusBadRequest)
		return  
  }
  libs.ResponseSuccess(w, r)
}

// DeleteUser DELETE /users/:userID
func DeleteUser(w http.ResponseWriter, r *http.Request) {
  userID := chi.URLParam(r, "userID")
  err := models.DeleteUser(userID)
  if err != nil {
		libs.ResponseError(w, r, "Error on delete user: " + err.Error(), http.StatusBadRequest)
		return  
  }
  libs.ResponseSuccess(w, r)
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


// UpdateUserContract PUT /user/:userID/contracts/:contractID
func UpdateUserContract(w http.ResponseWriter, r *http.Request) {

}

// DeleteUserContract DELETE /user/:userID/contracts/:contractID
func DeleteUserContract(w http.ResponseWriter, r *http.Request) {

}

// GetUserScores GET /users/:userID/scores
func GetUserScores(w http.ResponseWriter, r *http.Request) {

}