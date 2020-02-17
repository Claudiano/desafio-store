package controllers

import (
	"desafio-store/dtos"
	"desafio-store/services"
	"encoding/json"
	"fmt"
	"log"
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
	var accountDto = dtos.AccountDto{}

	err := json.NewDecoder(r.Body).Decode(&accountDto)
	checkError(err)

	res := serviceAccount.CreateAccount(accountDto)
	if res != nil {
		log.Println("Error: %v", err)
	}

	fmt.Fprintf(w, "Hello")
}

func checkError(err error) {
	if err != nil {
		log.Panic(err)
	}
}
