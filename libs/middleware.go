package libs

import (
  "net/http"
	"github.com/pressly/chi"
	"log"
)

// CORSMiddleware write some headers to handle CORS 
func CORSMiddleware(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    origin := r.Header.Get("Origin")
    //TODO verify origin whitelist here
    w.Header().Set("Access-Control-Allow-Origin", origin)
    w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PATCH, DELETE")
    w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
    w.Header().Set("Access-Control-Allow-Credentials", "true")
    next.ServeHTTP(w, r)
  })
}

// ValidateAdminOrCurrentUserMiddleware validate is admin or is resource belongs to current user
func ValidateAdminOrCurrentUserMiddleware(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    userClaims := r.Context().Value(ContextKey("userClaims")).(*UserClaims)
    if userClaims.IsAdmin || (userClaims.ID == chi.URLParam(r, "userID")) {
      next.ServeHTTP(w, r)
    } else {
      ResponseError(w, r, "Neither admin nor current company", http.StatusUnauthorized)
    }
  })
}

// ValidateAdminMiddleware validate is admin
func ValidateAdminMiddleware(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    userClaims := r.Context().Value(ContextKey("userClaims")).(*UserClaims)
    if userClaims.IsAdmin {
      next.ServeHTTP(w, r)
    } else {
      ResponseError(w, r, "Not admin", http.StatusUnauthorized)
    }
    next.ServeHTTP(w, r)
  })
}
