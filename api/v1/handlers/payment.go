package handlers

import (
	"io/ioutil"
	"net/http"

	"github.com/byte3/galactic.payment/internal/database"
)

type PaymentHandler struct{}

func (ph PaymentHandler) GetPaymentHistory(w http.ResponseWriter, r *http.Request) {
	// pass the token to auth
	// get the user id
	// query the db for user id
	// return payments
}

func (ph PaymentHandler) CreatePayment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Header.Get("Content-Type") != "application/json" {
		msg := "Content-Type Header is not in the required format (json)"
		http.Error(w, msg, http.StatusUnsupportedMediaType)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// make request to the auth server for payment
	// if confirmed update the db
	paymentCollection := database.GetPaymentCollection()
}
