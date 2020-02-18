package controllers

import (
	"desafio-store/dtos"
	"desafio-store/services"
	"desafio-store/utils"
	"encoding/json"
	"log"
	"net/http"
)

var transferService = services.TransferService{}

type TransferController struct{}

func (TransferController) FindAllTransfer(w http.ResponseWriter, r *http.Request) {

	transfers, err := transferService.FindAllTransfer()
	if err != nil {
		utils.RespondWithErrorJSON(w, http.StatusUpgradeRequired, err)
	} else {
		utils.RespondwithJSON(w, http.StatusAccepted, transfers)
	}
}

func (TransferController) CreateTransfer(w http.ResponseWriter, r *http.Request) {
	var transferDto = dtos.TransferDto{}

	err := json.NewDecoder(r.Body).Decode(&transferDto)
	if err != nil {
		log.Println("Error: %v", err)
		utils.RespondWithErrorJSON(w, http.StatusBadRequest, err)
	}

	err = transferService.CreateTransfer(transferDto)
	if err != nil {
		log.Println("Error: %v", err)
		utils.RespondWithErrorJSON(w, http.StatusBadGateway, err)
	} else {
		utils.RespondwithJSON(w, http.StatusCreated, "usuario criado")
	}

}
