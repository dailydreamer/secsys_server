package controllers

import (
  "net/http"
  "secsys/models"
  "secsys/libs"
)

// GetMessages GET /messages
func GetMessages(w http.ResponseWriter, r *http.Request) {
  messageList, err := models.GetMessages()
  if err != nil {
    libs.ResponseError(w, r, "Error on get messages: " + err.Error(), http.StatusInternalServerError)
  }
  libs.ResponseJSON(w, r, messageList)
}