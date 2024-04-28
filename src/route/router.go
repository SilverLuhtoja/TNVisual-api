package route

import (
	"net/http"

	"github.com/SilverLuhtoja/TNVisual/src/api/middleware"
	"github.com/SilverLuhtoja/TNVisual/src/common"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

func NewRouter(controllers *Controllers) *chi.Mux {
	main_router := newApiCorsRouter()
	apiRouter := chi.NewRouter()

	apiRouter.Get("/health", ReadinessHandler)
	apiRouter.Get("/errorCheck", ErrHandler)

	apiRouter.Post("/users", controllers.UserController.Create)

	apiRouter.Post("/login", controllers.AuthController.Login)
	// apiRouter.Post("/auth", controllers.AuthController.Authenticate)

	apiRouter.Get("/projects", controllers.ProjectController.GetProjects)
	apiRouter.Post("/projects", middleware.Authenticate(controllers.ProjectController.CreateProjects))

	apiRouter.Get("/verify", controllers.AuthController.VerifyKey)

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

func ReadinessHandler(w http.ResponseWriter, r *http.Request) {
	res := map[string]string{"status": "ok"}
	common.RespondWithJSON(w, 200, res)
}

func ErrHandler(w http.ResponseWriter, r *http.Request) {
	common.RespondWithError(w, 500, "Internal Server Error")
}
