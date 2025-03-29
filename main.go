package main

import (
	"net/http"
	"os"

	"github.com/IsraYehovaAsmaranda/go-payment-api/routes"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	godotenv.Load()

	portString := os.Getenv("PORT")
	if portString == "" {
		portString = "8080"
	}

	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	router.Mount("/v1", routes.RegisterRoutes())

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	logrus.Printf("Server starting on port %v", portString)

	if err := srv.ListenAndServe(); err != nil {
		logrus.Fatal(err)
	}
}
