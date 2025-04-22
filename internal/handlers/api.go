package handlers

import (
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
	"github.com/ryamgoh/go-api/internal/middleware"
)

func Handler(r *chi.Mux) {
	// Global Middleware
	// This will strip away endpoints with trailing backslash which could cause 404 error
	r.Use(chimiddle.StripSlashes)

	r.Route("/account", func(router chi.Router) {
		// Middleware for /account route
		router.Use(middleware.Authorization)

		router.Get("/credits", GetCreditBalance)
	})
}
