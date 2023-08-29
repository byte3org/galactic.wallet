package router

import (
	"github.com/byte3/galactic.wallet/api/v1/handlers"
	"github.com/byte3/galactic.wallet/api/v1/middlewares"
	"github.com/go-chi/chi"
)

type Wallet struct{}

func (w Wallet) Routes() chi.Router {
	r := chi.NewRouter()

    r.Route("/", func(publicRoutes chi.Router) {
        publicRoutes.Get("/", handlers.WalletHandler{}.GetAllWallets)
        publicRoutes.Post("/", handlers.WalletHandler{}.CreateWallet)
    })
	r.Route("user", func(userRoutes chi.Router) {
		userRoutes.Use(middlewares.ExtractUserId)
		userRoutes.Get("/user/balance", handlers.WalletHandler{}.GetWalletBalance)
		userRoutes.Get("/user/transactions", handlers.WalletHandler{}.GetWalletTransactions)
	})
	r.Group(func(protectedRoutes chi.Router) {
		protectedRoutes.Use(middlewares.ExtractUserId)
		protectedRoutes.Use(middlewares.SignatureVerify)
		protectedRoutes.Post("/user/deposit", handlers.WalletHandler{}.DepositToWallet)
		protectedRoutes.Post("/user/withdraw", handlers.WalletHandler{}.WithdrawFromWallet)
	})
	return r
}
