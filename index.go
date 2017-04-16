package main

import (
	"net/http"
	"log"
	"os"
	"github.com/pressly/chi"
  "github.com/pressly/chi/middleware"

	"secsys/config"
	"secsys/libs"
	"secsys/controllers"
)

func main() {
	config.InitConfig()
	libs.InitDb()

  r := chi.NewRouter()
  // A good base middleware stack
  r.Use(middleware.RequestID)
  r.Use(middleware.RealIP)
  r.Use(middleware.Logger)
  r.Use(middleware.Recoverer)
	r.Use(libs.CORSMiddleware)
	// root router for test
  r.Get("/", func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("secsys!"))
  })
	// add version prefix path
  r.Mount(config.VersionURL, router()) 
	// try to get heroku port first
	port := os.Getenv("PORT")
	if port == "" {
		port = config.Port
	}	
	log.Println("Service started at port "+port)
	http.ListenAndServe(":"+port, r)
}

func router() http.Handler {
  r := chi.NewRouter()
  r.Route("/auth", func(r chi.Router) {
    r.Post("/signup", controllers.SignUp)       // POST /auth/signup
    r.Post("/login", controllers.LogIn)         // POST /auth/login
  })
	// need jwt token for the api below
	r.Group(func(r chi.Router) {
		r.Use(libs.ValidateJWTMiddleware)
		r.Get("/user", controllers.GetUser)         // GET /user
	})
	return r
}