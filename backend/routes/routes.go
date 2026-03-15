package routes

import (
	"net/http"
	"rapid-rise-backend/handlers"
	"rapid-rise-backend/middlewares"

	"github.com/go-chi/chi/v5"
)

func RegisterRoutes() http.Handler {
	r := chi.NewRouter()

	// Global middlewares
	r.Use(middlewares.Logger)
	r.Use(middlewares.Recoverer)
	r.Use(middlewares.Cors)

	// Public routes
	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("server working"))
	})
	r.Get("/employees", handlers.GetEmployeesHandler)
	r.Post("/login", handlers.LoginHandler)

	// Protected routes – apply AuthMiddleware only to these
	r.With(middlewares.AuthMiddleware).Post("/scan", handlers.ScanHandler)
	// (Add other protected routes here)

	return r
}
