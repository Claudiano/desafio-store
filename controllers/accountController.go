package controllers

import (
	"desafio-store/dtos"
	"desafio-store/services"
	"desafio-store/utils"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

var serviceAccount = services.ServiceAccount{}

type AccountController struct{}

func (AccountController) FindAllAccount(w http.ResponseWriter, r *http.Request) {

	accounts, err := serviceAccount.FindAllAccount()
	if err != nil {
		utils.RespondWithErrorJSON(w, http.StatusUpgradeRequired, err)
	} else {
		utils.RespondwithJSON(w, http.StatusAccepted, accounts)
	}
}

func (AccountController) FindAccountById(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseUint(chi.URLParam(r, "account_id"), 10, 32)
	ballance, err := serviceAccount.FindBallanceById(id)
	if err != nil {
		log.Println("error: ", err)
		utils.RespondWithErrorJSON(w, http.StatusBadRequest, err)
	} else {
		utils.RespondwithJSON(w, http.StatusAccepted, ballance)
	}
}

func (AccountController) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var accountDto = dtos.AccountDto{}

	err := json.NewDecoder(r.Body).Decode(&accountDto)
	if err != nil {
		log.Println("Error: %v", err)
		utils.RespondWithErrorJSON(w, http.StatusBadRequest, err)
	}

	err = serviceAccount.CreateAccount(accountDto)
	if err != nil {
		log.Println("Error: %v", err)
		utils.RespondWithErrorJSON(w, http.StatusBadGateway, err)
	} else {
		utils.RespondwithJSON(w, http.StatusCreated, "usuario criado")
	}

}
