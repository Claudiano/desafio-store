package services

import (
	"desafio-store/dtos"
	"desafio-store/models"
	"desafio-store/repositories"
)

var accountRepository = repositories.AccountRepository{}

type ServiceAccount struct{}

func (ServiceAccount) FindAllAccount() ([]models.Account, error) {

	return accountRepository.FindAllAccounts()
}

func (ServiceAccount) FindBallanceById(id uint64) (float32, error) {

	account, err := accountRepository.FindAccountById(id)

	return account.Ballance, err
}

func (ServiceAccount) CreateAccount(accountDto dtos.AccountDto) error {

	var account = models.Account{}

	account.Name = accountDto.Name
	account.Cpf = accountDto.Cpf

	return accountRepository.Save(account)
}
