package repositories

import (
	"desafio-store/models"
)

type AccountRepository struct {
}

func (repository AccountRepository) FindAllAccounts() ([]models.Account, error) {

	var accounts []models.Account

	db := conn.Pgdb

	err := db.Find(&accounts).Error

	return accounts, err
}

func (repository AccountRepository) FindAccountById(id uint64) (models.Account, error) {
	var account models.Account
	db := conn.Pgdb

	err := db.Where("id = ?", id).First(&account).Error

	return account, err
}

func (repository AccountRepository) Save(account models.Account) error {
	db := conn.Pgdb

	err := db.Create(&account).Error

	return err
}

func (AccountRepository) UpdateBallance(account models.Account) error {
	db := conn.Pgdb
	err := db.Model(&account).Where("id = ?", account.Id).Update("ballance", account.Ballance).Error
	return err
}
