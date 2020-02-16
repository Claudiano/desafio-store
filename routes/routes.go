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

	port := os.Getenv("SERVER_PORT")

	routes := chi.NewRouter()

	routes.Route("/accounts", func(r chi.Router) {
		r.Get("/", accountController.FindAllAccount)
		r.Get("/{account_id}/balance", accountController.FindAccountById)
		r.Post("/", accountController.CreateAccount)
	})

	routes.Route("/transfers", func(r chi.Router) {
		r.Get("/", accountController.FindAllAccount)
		r.Post("/", accountController.CreateAccount)
	})

	http.ListenAndServe(port, routes)
}
