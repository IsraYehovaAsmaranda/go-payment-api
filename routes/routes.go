package routes

import (
	"net/http"

	"github.com/IsraYehovaAsmaranda/go-payment-api/handlers"
	"github.com/go-chi/chi"
)

func AuthRoutes() http.Handler {
	r := chi.NewRouter()

	r.Post("/register", handlers.RegisterHandler)
	r.Post("/login", handlers.LoginHandler)

	return r
}

func RegisterRoutes() http.Handler {
	r := chi.NewRouter()

	r.Mount("/auth", AuthRoutes())

	return r
}
