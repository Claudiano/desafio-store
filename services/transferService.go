package services

import (
	"desafio-store/dtos"
	"desafio-store/models"
	"desafio-store/repositories"
	"errors"
)

var transferRepository = repositories.TranferRepository{}

type TransferService struct{}

func (TransferService) FindAllTransfer() ([]models.Transfer, error) {

	return transferRepository.FindAllTransfers()
}

func (TransferService) CreateTransfer(transferDto dtos.TransferDto) error {
	var transfer models.Transfer

	if transferDto.Amount <= 0 {
		return errors.New("Valor não permitido para ser tranferido.")
	}

	accountOrigin, err := accountRepository.FindAccountById(transferDto.Account_origin_id)
	if err != nil {
		return err
	}

	accountDestination, err := accountRepository.FindAccountById(transferDto.Account_destination_id)
	if err != nil {
		return err
	}

	if accountOrigin.Ballance < transferDto.Amount {
		return errors.New("A conta origem não possui saldo suficiente para tranferencia.")
	}

	accountOrigin.Ballance -= transferDto.Amount
	accountDestination.Ballance += transferDto.Amount

	transfer.Account_origin_id = transferDto.Account_origin_id
	transfer.Account_destination_id = transferDto.Account_destination_id
	transfer.Amount = transferDto.Amount

	err = transferRepository.Save(transfer, accountOrigin, accountDestination)
	return err

}
