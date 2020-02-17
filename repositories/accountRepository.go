package repositories

import (
	"desafio-store/models"
	"errors"
	"fmt"
	"log"
)

type AccountRepository struct {
}

func (repository AccountRepository) FindAllAccounts() string {

	return "retorna todas as contas"
}
func (repository AccountRepository) FindAccountById(id uint64) string {
	return fmt.Sprintf("Conta com id: %v", id)
}
func (repository AccountRepository) Save(account models.Account) error {
	var err error
	if account.GetId() == 0 {
		err = repository.createAccount(account)
	} else {
		err = repository.update(account)
	}

	return err
}

func (AccountRepository) update(account models.Account) error {
	db := conn.PgConnect()
	log.Println(db.dbName)

	return errors.New("usuario não atualizado")
}

func (AccountRepository) createAccount(account models.Account) error {
	db := conn.Pgdb

	err := db.Create(&account).Error

	if err != nil {
		log.Println("erro")
	}

	return err
}
