package controllers

import (
  "net/http"
  "secsys/models"
  "secsys/libs"
)

// GetLogs GET /logs
func GetLogs(w http.ResponseWriter, r *http.Request) {
  logList, err := models.GetLogs()
  if err != nil {
    libs.ResponseError(w, r, "Error on get logs: " + err.Error(), http.StatusInternalServerError)
  }
  libs.ResponseJSON(w, r, logList)
}