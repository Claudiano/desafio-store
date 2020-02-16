package services

import (
	"desafio-store/dtos"
	"desafio-store/models"
	"desafio-store/repositories"
)

var repositoryAccount = repositories.AccountRepository{}

type ServiceAccount struct{}

func (ServiceAccount) FindAllAccount() string {

	return repositoryAccount.FindAllAccounts()
}

func (ServiceAccount) FindAccountById(id uint64) string {

	return repositoryAccount.FindAccountById(id)
}

func (ServiceAccount) CreateAccount(accountDto dtos.AccountDto) string {

	var account models.Account

	return repositoryAccount.Save()
}
