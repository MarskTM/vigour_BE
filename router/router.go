package router

import (
	"net/http"

	"vigour/controller"
	_ "vigour/docs"
	"vigour/infrastructure"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	_ "github.com/go-chi/jwtauth"
	"github.com/go-chi/render"
	httpSwagger "github.com/swaggo/http-swagger"
)

func Router() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.URLFormat)
	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	Cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // Use this to allow specific origin hosts
		// AllowedOrigins: []string{"*"},
		// AllowOriginFunc:  checkOrigin,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	r.Use(Cors.Handler)

	// Swagger route
	r.Get("/api/v1/swagger/*", httpSwagger.Handler(
		httpSwagger.URL(infrastructure.GetHTTPSwagger()),
	))

	// Declear Controller
	userController := controller.NewUserController()
	profileController := controller.NewProfileController()

	r.Route("/api/v1", func(router chi.Router) {
		// Public routes
		// Protected routes
		router.Post("/user/create", userController.CreateUser)
		router.Post("/user/login", userController.Login)
		router.Post("/user/login/jwt", userController.LoginWithToken)
		router.Put("/profile/upsert", profileController.Upsert)

		router.Group(func(protectedRoute chi.Router) {
			// Setup middleware
			// protectedRoute.Use(jwtauth.Verifier(infrastructure.GetEncodeAuth()))
			// protectedRoute.Use(jwtauth.Authenticator)

			protectedRoute.Route("/user", func(subRoute chi.Router) {
				subRoute.Get("/all", userController.GetAll)
				subRoute.Get("/{uid}", userController.GetById)
				subRoute.Put("/update", userController.UpdateUser)
				subRoute.Delete("/delete/{uid}", userController.DeleteUser)
				subRoute.Get("/wname", userController.GetByUsername)
			})

			protectedRoute.Route("/profile", func(subRoute chi.Router) {
				subRoute.Get("/all", profileController.GetAll)
				subRoute.Get("/{id}", profileController.GetById)
				subRoute.Get("/user/{user_id}", profileController.GetByUserId)
				subRoute.Post("/create", profileController.Create)
				subRoute.Put("/update", profileController.Update)
				subRoute.Delete("/delete/{uid}", profileController.Delete)
			})
		})
	})
	return r
}
