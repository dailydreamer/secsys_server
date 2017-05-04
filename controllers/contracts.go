package controllers

import (
  "encoding/json"
  "net/http"
  "github.com/pressly/chi"
  "secsys/models"
  "secsys/libs"
)

// GetContracts GET /contracts
func GetContracts(w http.ResponseWriter, r *http.Request) {
  contractList, err := models.GetContracts()
  if err != nil {
    libs.ResponseError(w, r, "Error on get contracts: " + err.Error(), http.StatusInternalServerError)
  }
  libs.ResponseJSON(w, r, contractList)
}

// CreateContract POST /contracts
func CreateContract(w http.ResponseWriter, r *http.Request) {
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
  contract.UserID = dbUser.ID
  contract.ID, err = models.CreateContract(contract)
  if err != nil {
		libs.ResponseError(w, r, "Error on create contract: " + err.Error(), http.StatusInternalServerError)
		return
	}
  libs.ResponseSuccess(w, r)
}

// GetContract GET /contracts/:contractID
func GetContract(w http.ResponseWriter, r *http.Request) {
  contractID := chi.URLParam(r, "contractID")
  contract, err := models.GetContractByID(contractID)
  if err != nil {
    libs.ResponseError(w, r, "Error on get contract: " + err.Error(), http.StatusInternalServerError)
  }
  libs.ResponseJSON(w, r, contract)
}


// UpdateContract PUT /contracts/:contractID
func UpdateContract(w http.ResponseWriter, r *http.Request) {
  contractID := chi.URLParam(r, "contractID")
	var contract models.Contract
	err := json.NewDecoder(r.Body).Decode(&contract)
	if err != nil {
		libs.ResponseError(w, r, "Error on parse json: " + err.Error(), http.StatusBadRequest)
		return
	}
  if contractID != contract.ID {
    libs.ResponseError(w, r, "contractID not match", http.StatusUnauthorized)
		return
  }
  err = models.UpdateContract(contract)
  if err != nil {
		libs.ResponseError(w, r, "Error on update contract: " + err.Error(), http.StatusInternalServerError)
		return  
  }
  libs.ResponseSuccess(w, r)
}

// DeleteContract DELETE /contracts/:contractID
func DeleteContract(w http.ResponseWriter, r *http.Request) {
  contractID := chi.URLParam(r, "contractID")
  err := models.DeleteContractByID(contractID)
  if err != nil {
		libs.ResponseError(w, r, "Error on delete contract: " + err.Error(), http.StatusInternalServerError)
		return  
  }
  libs.ResponseSuccess(w, r)
}