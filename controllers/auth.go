package controllers

import (
	"net/http"
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	
	"secsys/models"
	"secsys/libs"
)

type tokenResponse struct {
	Token string `json:"token"`
}

// SignUp POST /auth/signup
func SignUp(w http.ResponseWriter, r *http.Request) {
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
	user.ID, err = models.CreateUser(user.Phone, string(hashedPassword))
	if err != nil {
		libs.ResponseError(w, r, "Error on create user: " + err.Error(), http.StatusInternalServerError)
		return
	}
	token, err := libs.GenerateJWT(user.ID)
	if err != nil {
		libs.ResponseError(w, r, "Error on sign token: " + err.Error(), http.StatusInternalServerError)
		return
	}
	
	libs.ResponseJSON(w, r, tokenResponse{token})
}

// LogIn POST /auth/login
func LogIn(w http.ResponseWriter, r *http.Request) {
	//userid := r.Context().Value(libs.ContextKey("userid"))
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
	//get user from db
	dbUser, err := models.FindUserByPhone(user.Phone)
	if err != nil {
		libs.ResponseError(w, r, "User not exist: " + err.Error(), http.StatusUnauthorized)
		return
	}
	// check password
	err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password))
	if err != nil {
		libs.ResponseError(w, r, "Password not correct: " + err.Error(), http.StatusUnauthorized)
		return		
	}
	// generate JWT
	token, err := libs.GenerateJWT(dbUser.ID)
	if err != nil {
		libs.ResponseError(w, r, "Error on sign token: " + err.Error(), http.StatusInternalServerError)
		return
	}
	
	libs.ResponseJSON(w, r, tokenResponse{token})
}