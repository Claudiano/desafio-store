package repositories

import (
	"desafio-store/models"
	"fmt"
)

type AccountRepository struct {
}

func (repository AccountRepository) FindAllAccounts() string {
	return "retorna todas as contas"
}
func (repository AccountRepository) FindAccountById(id uint64) string {
	return fmt.Sprintf("Conta com id: %v", id)
}
func (repository AccountRepository) Save(account models.Account) string {

	var result string
	if account.GetId() == 0 {
		result = repository.create(account)
	} else {
		result = repository.update(account)
	}

	return result
}

func (AccountRepository) update(account models.Account) string {
	return "usuario atualizado"
}

func (AccountRepository) create(account models.Account) string {
	return "usuario criado"
}
