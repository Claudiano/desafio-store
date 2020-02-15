package controllers

import (
	"desafio-store/services"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

var serviceAccount = services.ServiceAccount{}

type AccountController struct{}

func (AccountController) FindAllAccount(w http.ResponseWriter, r *http.Request) {

	res := serviceAccount.FindAllAccount()
	fmt.Fprintf(w, "%v", res)
}

func (AccountController) FindAccountById(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseUint(chi.URLParam(r, "account_id"), 10, 32)
	res := serviceAccount.FindAccountById(id)
	fmt.Fprintf(w, "Hello World! %s", res)
}

func (AccountController) CreateAccount(w http.ResponseWriter, r *http.Request) {
	res := serviceAccount.CreateAccount()
	fmt.Fprintf(w, "%v", res)
}
