package main

import (
	"log"
	"net/http"
	"time"

	"github.com/byte3/galactic.wallet/api/v1/router"
	"github.com/byte3/galactic.wallet/config"
	"github.com/byte3/galactic.wallet/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func intitializePaymentRoutes(r *chi.Mux) {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("galactic payment gateway"))
	})
	r.Route("/v1", func(r chi.Router) {
		r.Mount("/", router.SetupRoutes())
	})
}

func initializePaymentServer(config *config.Config) *chi.Mux {
	r := chi.NewRouter()

	r.Use(
		middleware.RedirectSlashes,
		middleware.Recoverer,
		middleware.Heartbeat("/health"),
		middleware.Logger,
		middleware.AllowContentType("application/json"),
	)

	r.Use(middleware.Timeout(20 * time.Second))

	intitializePaymentRoutes(r)

	return r
}

func main() {
	log.Println("[!] Starting galactic.payment service...")

	// load config file
	config := config.GetConfig()

	// initialize database
	database.Initialize(config)

	r := initializePaymentServer(config)
}
