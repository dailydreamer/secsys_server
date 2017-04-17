package libs

import (
  "net/http"
)

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

// ValidateAdminMiddleware validate is admin
func ValidateAdminMiddleware(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    //TODO
  })
}

// ValidateCurrentCompanyMiddleware validate is resource belongs to current company
func ValidateCurrentCompanyMiddleware(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    //TODO
  })
}

// ValidateAdminOrCurrentCompanyMiddleware validate is admin or is resource belongs to current company
func ValidateAdminOrCurrentCompanyMiddleware(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    //TODO
  })
}