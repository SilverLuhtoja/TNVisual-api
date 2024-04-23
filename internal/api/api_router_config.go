package api

import (
	"github.com/SilverLuhtoja/TNVisual/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

type ApiConfig struct {
	DB *database.Queries
}

func NewRouter(cfg *ApiConfig) *chi.Mux {
	main_router := newApiCorsRouter()

	apiRouter := chi.NewRouter()

	apiRouter.Get("/health", ReadinessHandler)
	apiRouter.Get("/errorCheck", ErrHandler)
	apiRouter.Post("/login", cfg.LoginHandler)
	apiRouter.Post("/auth", cfg.AuthenticateKeyHandler)

	apiRouter.Post("/users", cfg.CreateUserHandler)

	apiRouter.Get("/projects", cfg.GetAllProjects)
	apiRouter.Post("/projects", cfg.middlewareAuth(cfg.CreateProjectHandler))

	apiRouter.Get("/verify", GetCookieHandler)

	main_router.Mount("/", apiRouter)
	return main_router
}

func newApiCorsRouter() *chi.Mux {
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		MaxAge:           300,
	}))
	return router
}
