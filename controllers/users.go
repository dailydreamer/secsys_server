package controllers

import (
  "net/http"

  "secsys/libs"
	"log"
  "secsys/models"
	"github.com/pressly/chi"
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
		libs.ResponseError(w, r, "Error on update user: " + err.Error(), http.StatusInternalServerError)
		return  
  }
  libs.ResponseSuccess(w, r)
}

// DeleteUser DELETE /users/:userID
func DeleteUser(w http.ResponseWriter, r *http.Request) {
  userID := chi.URLParam(r, "userID")
  err := models.DeleteUser(userID)
  if err != nil {
		libs.ResponseError(w, r, "Error on delete user: " + err.Error(), http.StatusInternalServerError)
		return  
  }
  libs.ResponseSuccess(w, r)
}

// GetUserContracts GET /user/:userID/contracts
func GetUserContracts(w http.ResponseWriter, r *http.Request) {
  userID := chi.URLParam(r, "userID")
  contracts, err := models.GetUserContracts(userID)
  if err != nil {
		libs.ResponseError(w, r, "Error on get user contracts: " + err.Error(), http.StatusInternalServerError)
		return  
  }
  libs.ResponseJSON(w, r, contracts)
}

// CreateUserContract POST /user/:userID/contracts
func CreateUserContract(w http.ResponseWriter, r *http.Request) {
  userID := chi.URLParam(r, "userID")
	var contract models.Contract
	err := json.NewDecoder(r.Body).Decode(&contract)
	if err != nil {
		libs.ResponseError(w, r, "Error on parse json: " + err.Error(), http.StatusBadRequest)
		return
	}
  dbUser, err := models.GetUserByComName(contract.ComName)
	if err != nil {
		libs.ResponseError(w, r, "Error on get user by comName: " + err.Error(), http.StatusInternalServerError)
		return
	}
  if dbUser.ID != userID {
		libs.ResponseError(w, r, "Wrong company name: " + err.Error(), http.StatusUnprocessableEntity)
		return    
  }
  contract.UserID = dbUser.ID
  contract.ID, err := models.CreateContract(contract)
  if err != nil {
		libs.ResponseError(w, r, "Error on create contract: " + err.Error(), http.StatusInternalServerError)
		return
	}
  libs.ResponseSuccess(w, r)
}

// GetUserContract GET /user/:userID/contracts/:contractID
func GetUserContract(w http.ResponseWriter, r *http.Request) {
  userID := chi.URLParam(r, "userID")
  contractID := chi.URLParam(r, "contractID")
  dbContract, err := models.GetContractByID(contractID)
	if err != nil {
		libs.ResponseError(w, r, "Error on get contract by id: " + err.Error(), http.StatusInternalServerError)
		return
	}
  if userID != dbContract.UserID {
    libs.ResponseError(w, r, "Contract not belong to the company: " + err.Error(), http.StatusUnauthorized)
		return
  }
  libs.ResponseJSON(w, r, dbContract)
}


// UpdateUserContract PUT /user/:userID/contracts/:contractID
func UpdateUserContract(w http.ResponseWriter, r *http.Request) {
  userID := chi.URLParam(r, "userID")
  contractID := chi.URLParam(r, "contractID")
  dbContract, err := models.GetContractByID(contractID)
	if err != nil {
		libs.ResponseError(w, r, "Error on get contract by id: " + err.Error(), http.StatusInternalServerError)
		return
	}
  if userID != dbContract.UserID {
    libs.ResponseError(w, r, "Contract not belong to the company: " + err.Error(), http.StatusUnauthorized)
		return
  }
	var contract models.Contract
	err := json.NewDecoder(r.Body).Decode(&contract)
	if err != nil {
		libs.ResponseError(w, r, "Error on parse json: " + err.Error(), http.StatusBadRequest)
		return
	}
  err := models.UpdateContract(contract)
	if err != nil {
		libs.ResponseError(w, r, "Error on update contract: " + err.Error(), http.StatusInternalServerError)
		return
	}
  libs.ResponseSuccess(w, r)
}

// DeleteUserContract DELETE /user/:userID/contracts/:contractID
func DeleteUserContract(w http.ResponseWriter, r *http.Request) {
  userID := chi.URLParam(r, "userID")
  contractID := chi.URLParam(r, "contractID")
  dbContract, err := models.GetContractByID(contractID)
	if err != nil {
		libs.ResponseError(w, r, "Error on get contract by id: " + err.Error(), http.StatusInternalServerError)
		return
	}
  if userID != dbContract.UserID {
    libs.ResponseError(w, r, "Contract not belong to the company: " + err.Error(), http.StatusUnauthorized)
		return
  }
  err := models.DeleteContractByID(contractID)
	if err != nil {
		libs.ResponseError(w, r, "Error on delete contract: " + err.Error(), http.StatusInternalServerError)
		return
	}
  libs.ResponseSuccess(w, r)
}

// GetUserScores GET /users/:userID/scores
func GetUserScores(w http.ResponseWriter, r *http.Request) {
  userID := chi.URLParam(r, "userID")
  scores, err := models.GetUserScores(userID)
  if err != nil {
		libs.ResponseError(w, r, "Error on get user scores: " + err.Error(), http.StatusInternalServerError)
		return  
  }
  libs.ResponseJSON(w, r, scores)
}