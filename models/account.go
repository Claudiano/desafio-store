package models

import (
	"time"
)

type Account struct {
	Id         uint64  `gorm:PRIMARY_KEY;AUTO_INCREMENT"`
	Name       string  `gorm:"type:varchar(200)"`
	Cpf        string  `gorm:"type:varchar(11);unique;not null"`
	Ballance   float32 `gorm:"DEFAULT:0.0"`
	Created_at time.Time
}

func (account Account) GetId() uint64 {
	return account.Id
}

func (account *Account) SetId(Id uint64) {
	account.Id = Id
}

func (account Account) GetName() string {
	return account.Name
}
func (account *Account) SetName(Name string) {
	account.Name = Name
}

func (account Account) GetCpf() string {
	return account.Cpf
}

func (account *Account) SetCpf(Cpf string) {
	account.Cpf = Cpf
}

func (account Account) GetBallance() float32 {
	return account.Ballance
}

func (account *Account) SetBallance(Ballance float32) {
	account.Ballance = Ballance
}

func (account Account) GetCreatedAt() time.Time {
	return account.Created_at
}
