package router

import (
	"github.com/byte3/galactic.wallet/api/v1/handlers"
	"github.com/byte3/galactic.wallet/api/v1/middlewares"
	"github.com/go-chi/chi"
)

type Wallet struct{}

func (w Wallet) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", handlers.WalletHandler{}.GetAllWallets)
	r.Post("/", handlers.WalletHandler{}.CreateWallet)
	r.Group(func(r chi.Router) {
		r.Use(middlewares.ExtractUserId)
		r.Get("/user/balance", handlers.WalletHandler{}.GetWalletBalance)
		r.Get("/user/transactions", handlers.WalletHandler{}.GetWalletTransactions)
	})
	r.Group(func(r chi.Router) {
		r.Use(middlewares.ExtractUserId)
		r.Use(middlewares.SignatureVerify)
		r.Post("/user/deposit", handlers.WalletHandler{}.DepositToWallet)
		r.Post("/user/withdraw", handlers.WalletHandler{}.WithdrawFromWallet)
	})
	return r
}
