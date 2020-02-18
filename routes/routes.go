package routes

import (
	"desafio-store/controllers"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	_ "github.com/joho/godotenv/autoload"
)

func StartServer() {
	var accountController = controllers.AccountController{}
	var transferController = controllers.TransferController{}

	port := os.Getenv("SERVER_PORT")

	routes := chi.NewRouter()

	routes.Route("/accounts", func(r chi.Router) {
		r.Get("/", accountController.FindAllAccount)
		r.Get("/{account_id}/ballance", accountController.FindAccountById)
		r.Post("/", accountController.CreateAccount)
	})

	routes.Route("/transfers", func(r chi.Router) {
		r.Get("/", transferController.FindAllTransfer)
		r.Post("/", transferController.CreateTransfer)
	})

	http.ListenAndServe(port, routes)
}
