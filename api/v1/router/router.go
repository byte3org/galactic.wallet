package router

import "github.com/go-chi/chi"

func SetupRoutes() *chi.Mux {
	r := chi.NewRouter()
	r.Mount("/wallet", Wallet{}.Routes())
	return r
}
