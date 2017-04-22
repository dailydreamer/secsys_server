package libs

import (
  "net/http"
  "encoding/json"
)

// JSONError is error json respones struct
type JSONError struct {
  Error string `json:"error"`
}

// JSONSuccess is success json respones struct
type JSONSuccess struct {
  Success bool `json:"success"`
}

// ResponseJSON return json http respones
func ResponseJSON(w http.ResponseWriter, r *http.Request, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(v)
}

// ResponseError return json http error respones
func ResponseError(w http.ResponseWriter, r *http.Request, error string, code int) {
	w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(code)
	json.NewEncoder(w).Encode(JSONError{error})
}

// ResponseSuccess return json success respones
func ResponseSuccess(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(JSONSuccess{true})
}