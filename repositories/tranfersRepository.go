package repositories

import "desafio-store/models"

type TranferRepository struct{}

func (TranferRepository) FindAllTransfers() ([]models.Transfer, error) {
	var transfers []models.Transfer
	db := conn.Pgdb

	err := db.Find(&transfers).Error

	return transfers, err
}

func (TranferRepository) Save(transfer models.Transfer) error {
	db := conn.Pgdb

	err := db.Create(&transfer).Error
	return err
}
