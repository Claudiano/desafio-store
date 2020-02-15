package services

import "fmt"

type ServiceAccount struct{}

func (ServiceAccount) FindAllAccount() string {

	return "todas as contas"
}

func (ServiceAccount) FindAccountById(id uint64) string {

	return fmt.Sprintf("Conta com id: %v", id)
}

func (ServiceAccount) CreateAccount() string {

	return "Criar conta"
}
