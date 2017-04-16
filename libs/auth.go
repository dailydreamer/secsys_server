package libs

import (
  "time"
  "net/http"
  "strings"
  "fmt"
  jwt "github.com/dgrijalva/jwt-go"
	"log"
	"context"
  "secsys/config"
)

type UserClaims struct {
  ID string `json:"id"`
  jwt.StandardClaims
}

type ContextKey string

// GenerateJWT sign jwt token
func GenerateJWT(id string) (string, error) {
  claims := UserClaims{
    id,
    jwt.StandardClaims{
      Issuer: "pootoo",
      ExpiresAt: time.Now().Add(time.Minute * 60 * 24 * 30).Unix(),
    },
  }  
	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
  return rawToken.SignedString([]byte(config.JWTSecret))
}

// ValidateJWT middleware to validate jwt token and set context
func ValidateJWT(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
  rawToken := TokenFromAuthHeader(r)
  if rawToken == "" {
    // no auth token
    next(w, r)
    return
  }
  token, err := jwt.ParseWithClaims(rawToken, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
    if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
        return "", fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
    }
    return []byte(config.JWTSecret), nil
  })
  if err != nil {
		switch err.(type) {
		case *jwt.ValidationError: // JWT validation error
			vErr := err.(*jwt.ValidationError)
			switch vErr.Errors {
			case jwt.ValidationErrorExpired: //JWT expired
        ResponseError(w, r, "Access Token is expired, get a new Token", http.StatusUnauthorized)
				return
			default:
        log.Println(err)
			}
		default:
      log.Println(err)
		}
  }
  if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
    ctx := context.WithValue(r.Context(), ContextKey("userid"), claims.ID)
    r = r.WithContext(ctx)
  }
  next(w, r)
}

// TokenFromAuthHeader is a "TokenExtractor" that takes a given request and extracts
// the JWT token from the Authorization header.
func TokenFromAuthHeader(r *http.Request) string {
	// Look for an Authorization header
	if ah := r.Header.Get("Authorization"); ah != "" {
		// Should be a bearer token
		if len(ah) > 6 && strings.ToUpper(ah[0:6]) == "BEARER" {
			return ah[7:]
		}
	}
  return ""
}