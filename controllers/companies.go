package controllers

import (
  "net/http"
)

// GetCompanies GET /companies
func GetCompanies(w http.ResponseWriter, r *http.Request) {

}

// CreateCompany POST /companies
func CreateCompany(w http.ResponseWriter, r *http.Request) {

}

// GetCompany GET /companies/:companyID
func GetCompany(w http.ResponseWriter, r *http.Request) {
	//companyID := chi.URLParam(r, "companyID")
}

// UpdateCompany PUT /companies/:companyID
func UpdateCompany(w http.ResponseWriter, r *http.Request) {

}

// DeleteCompany DELETE /companies/:companyID
func DeleteCompany(w http.ResponseWriter, r *http.Request) {

}

// GetCompanyContracts GET /companies/:companyID/contracts
func GetCompanyContracts(w http.ResponseWriter, r *http.Request) {

}