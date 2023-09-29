package application

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/mmvergara/golang-gala/backend/handlers"
	"github.com/mmvergara/golang-gala/backend/repository/post"
)

func (app *App) loadRoutes() {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("Hello World"))
	})

	router.Route("/post", app.loadPostRoutes)

	app.router = router
}

func (app *App) loadPostRoutes(router chi.Router) {
	postHandler := &handlers.Post{
		Repo: &post.RedisRepo{
			Client: app.redisDb,
		},
	}

	router.Post("/", postHandler.Create)
	router.Get("/", postHandler.List)
	router.Get("/{id}", postHandler.GetByID)
	router.Delete("/{id}", postHandler.DeleteByID)
}
