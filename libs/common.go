package libs

import (
  "net/http"
  "encoding/json"
)

// JSONError is error json respones struct
type JSONError struct {
  Error string `json:"error"`
}

// CORSMiddleware write some headers to handle CORS 
func CORSMiddleware(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    origin := r.Header.Get("Origin")
    //TODO verify origin whitelist here
    w.Header().Set("Access-Control-Allow-Origin", origin)
    w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
    w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
    w.Header().Set("Access-Control-Allow-Credentials", "true")
    next.ServeHTTP(w, r)
  })
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