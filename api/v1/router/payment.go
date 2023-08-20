package router

import (
	"github.com/byte3/galactic.wallet/api/v1/handlers"
	"github.com/go-chi/chi"
)

type Payment struct{}

func (p Payment) Routes() chi.Router {
	r := chi.NewRouter()

	r.Group(func(r chi.Router) {
		r.Get("/history", handlers.PaymentHandler{}.GetPaymentHistory)
		r.Post("/pay", handlers.PaymentHandler{}.CreatePayment)
	})

	return r
}
