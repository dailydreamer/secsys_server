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
    r.Post("/signup", controllers.SignUp)       							// POST /auth/signup
    r.Post("/login", controllers.LogIn)         							// POST /auth/login
  })			
	// need to be admin or current user for the api below			
	r.Group(func(r chi.Router) {			
		r.Use(libs.ValidateJWTMiddleware)
		r.Use(libs.ValidateAdminOrCurrentUserMiddleware)

		r.Route("/users/:userID", func(r chi.Router) {
			r.Get("/", controllers.GetUser)												// GET /users/:userID
			r.Patch("/", controllers.UpdateUser)									// PATCH /users/:userID
			r.Delete("/", controllers.DeleteUser)									// DELETE /users/:userID

			r.Route("/contracts", func(r chi.Router) {
				r.Get("/", controllers.GetContracts)								// GET /users/:userID/contracts
				r.Post("/", controllers.CreateContract)							// POST /users/:userID/contracts
				r.Route("/:contractID", func(r chi.Router) {
					r.Get("/", controllers.GetContract)								// GET /users/:userID/contracts/:contractID
					r.Patch("/", controllers.UpdateContract)					// PATCH /users/:userID/contracts/:contractID
					r.Delete("/", controllers.DeleteContract)					// DELETE /users/:userID/contracts/:contractID
				})
			})

			r.Get("/scores", controllers.GetScores)								// GET /users/:userID/scores
		})
		// need to be admin for the api below	
		r.Group(func(r chi.Router) {
			r.Use(libs.ValidateAdminMiddleware)

			r.Get("/users", controllers.GetUsers)									// GET /users
			r.Post("/users", controllers.CreateUser)							// POST /users

			r.Route("/contracts", func(r chi.Router) {
				r.Get("/", controllers.GetContracts)								// GET /contracts
				r.Post("/", controllers.CreateContract)							// POST /contracts
				r.Route("/:contractID", func(r chi.Router) {
					r.Get("/", controllers.GetContract)								// GET /contracts/:contractID
					r.Patch("/", controllers.UpdateContract)					// PATCH /contracts/:contractID
					r.Delete("/", controllers.DeleteContract)					// DELETE /contracts/:contractID
				})
			})

			r.Route("/scores", func(r chi.Router) {
				r.Get("/", controllers.GetScores)										// GET /scores
				r.Post("/", controllers.CreateScore)								// POST /scores
				r.Route("/:scoreID", func(r chi.Router) {
					r.Get("/", controllers.GetScore)									// GET /scores/:scoreID
					r.Patch("/", controllers.UpdateScore)							// PATCH /scores/:scoreID
					r.Delete("/", controllers.DeleteScore)						// DELETE /scores/:scoreID
				})
			})

			r.Get("/logs", controllers.GetLogs)										// GET /logs
			r.Get("/messages", controllers.GetMessages)						// GET /messages
		})
	})
	return r
}