package routes

import (
	"net/http"

	"github.com/IsraYehovaAsmaranda/go-payment-api/handlers"
	"github.com/IsraYehovaAsmaranda/go-payment-api/middleware"
	"github.com/go-chi/chi"
)

func AuthRoutes() http.Handler {
	r := chi.NewRouter()

	r.Post("/register", handlers.RegisterHandler)
	r.Post("/login", handlers.LoginHandler)

	r.Group(func(protected chi.Router) {
		protected.Use(middleware.AuthMiddleware)
		protected.Post("/logout", handlers.LogoutHandler)
	})

	return r
}

func PaymentRoutes() http.Handler {
	r := chi.NewRouter()

	r.Group(func(protected chi.Router) {
		protected.Use(middleware.AuthMiddleware)
		protected.Post("/", handlers.PaymentHandler)
	})

	return r
}

func RegisterRoutes() http.Handler {
	r := chi.NewRouter()

	r.Mount("/auth", AuthRoutes())
	r.Mount("/payment", PaymentRoutes())

	return r
}
