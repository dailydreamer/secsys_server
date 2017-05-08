package controllers

import (
  "encoding/json"
  "net/http"
  "github.com/pressly/chi"
  "secsys/models"
  "secsys/libs"
)

// GetScores GET /scores
func GetScores(w http.ResponseWriter, r *http.Request) {
  scoreList, err := models.GetScores()
  if err != nil {
    libs.ResponseError(w, r, "Error on get scores: " + err.Error(), http.StatusInternalServerError)
  }
  libs.ResponseJSON(w, r, scoreList)
}

// CreateScore POST /scores
func CreateScore(w http.ResponseWriter, r *http.Request) {
	var score models.Score
	err := json.NewDecoder(r.Body).Decode(&score)
	if err != nil {
		libs.ResponseError(w, r, "Error on parse json: " + err.Error(), http.StatusBadRequest)
		return
	}
	// check required field
	if score.ComName == "" {
		libs.ResponseError(w, r, "Field comName required", http.StatusUnprocessableEntity)
		return
	}
  dbUser, err := models.GetUserByComName(score.ComName)
	if err != nil {
		libs.ResponseError(w, r, "Error on get user by comName: " + err.Error(), http.StatusInternalServerError)
		return
	}
  score.UserID = dbUser.ID
  score.ID, err = models.CreateScore(score)
  if err != nil {
		libs.ResponseError(w, r, "Error on create score: " + err.Error(), http.StatusInternalServerError)
		return
	}
  libs.ResponseSuccess(w, r)
}

// GetScore GET /scores/:scoreID
func GetScore(w http.ResponseWriter, r *http.Request) {
  scoreID := chi.URLParam(r, "scoreID")
  score, err := models.GetScoreByID(scoreID)
  if err != nil {
    libs.ResponseError(w, r, "Error on get score: " + err.Error(), http.StatusInternalServerError)
  }
  libs.ResponseJSON(w, r, score)
}


// UpdateScore PUT /scores/:scoreID
func UpdateScore(w http.ResponseWriter, r *http.Request) {
  scoreID := chi.URLParam(r, "scoreID")
	var score models.Score
	err := json.NewDecoder(r.Body).Decode(&score)
	if err != nil {
		libs.ResponseError(w, r, "Error on parse json: " + err.Error(), http.StatusBadRequest)
		return
	}
  if scoreID != score.ID {
    libs.ResponseError(w, r, "scoreID not match", http.StatusUnauthorized)
		return
  }
  err = models.UpdateScore(score)
  if err != nil {
		libs.ResponseError(w, r, "Error on update score: " + err.Error(), http.StatusInternalServerError)
		return  
  }
  libs.ResponseSuccess(w, r)
}

// DeleteScore DELETE /scores/:scoreID
func DeleteScore(w http.ResponseWriter, r *http.Request) {
  scoreID := chi.URLParam(r, "scoreID")
  err := models.DeleteScoreByID(scoreID)
  if err != nil {
		libs.ResponseError(w, r, "Error on delete score: " + err.Error(), http.StatusInternalServerError)
		return  
  }
  libs.ResponseSuccess(w, r)
}