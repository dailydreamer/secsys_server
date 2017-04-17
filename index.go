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
    r.Post("/signup", controllers.SignUp)       								// POST /auth/signup
    r.Post("/login", controllers.LogIn)         								// POST /auth/login
  })			
	// need jwt token for the api below			
	r.Group(func(r chi.Router) {			
		r.Use(libs.ValidateJWTMiddleware)

		r.Route("/user", func(r chi.Router) {
			r.Get("/", controllers.GetCurrentUser)												// GET /user
			r.Patch("/", controllers.UpdateCurrentUser)										// PATCH /user
			r.Delete("/", controllers.DeleteCurrentUser)									// DELETE /user
		})

		r.Group(func(r chi.Router) {
			r.Use(libs.ValidateAdminOrCurrentCompanyMiddleware)
			r.Route("/companies/:companyID", func(r chi.Router) {
				r.Get("/", controllers.GetCompany)												// GET /companies/:companyID
				r.Patch("/", controllers.UpdateCompany)										// PATCH /companies/:companyID
				r.Delete("/", controllers.DeleteCompany)									// DELETE /companies/:companyID
			})
		})

		r.Group(func(r chi.Router) {
			r.Use(libs.ValidateCurrentCompanyMiddleware)
			r.Route("/companies/:companyID", func(r chi.Router) {
				r.Get("/contracts", controllers.GetCompanyContracts)				// GET /companies/:companyID/contracts
				r.Post("/contracts", controllers.CreateContract)						// POST /companies/:companyID/contracts
				r.Route("/contracts/:contractID", func(r chi.Router) {
					r.Get("/", controllers.GetContract)							// GET /companies/:companyID/contracts/:contractID
					r.Patch("/", controllers.UpdateContract)				// PATCH /companies/:companyID/contracts/:contractID
					r.Delete("/", controllers.DeleteContract)				// DELETE /companies/:companyID/contracts/:contractID
				})
			})
		})

		r.Group(func(r chi.Router) {
			r.Use(libs.ValidateAdminMiddleware)
			r.Get("/companies", controllers.GetCompanies)								// GET /companies
			r.Get("/contracts", controllers.GetContracts)								// GET /contracts
			r.Post("/contracts", controllers.CreateContract)						// POST /contracts
			r.Route("/contracts/:contractID", func(r chi.Router) {
				r.Get("/", controllers.GetContract)												// GET /contracts/:contractID
				r.Patch("/", controllers.UpdateContract)									// PATCH /contracts/:contractID
				r.Delete("/", controllers.DeleteContract)									// DELETE /contracts/:contractID
			})
		})

	})
	return r
}