package models

import (
	"time"
)

type Account struct {
	id         uint64
	name       string
	cpf        string
	ballance   float32
	created_at time.Time
}

func (account Account) GetId() uint64 {
	return account.id
}

func (account *Account) SetId(id uint64) {
	account.id = id
}

func (account Account) GetName() string {
	return account.name
}
func (account *Account) SetName(name string) {
	account.name = name
}

func (account Account) GetCpf() string {
	return account.cpf
}

func (account *Account) SetCpf(cpf string) {
	account.cpf = cpf
}

func (account Account) GetBallance() float32 {
	return account.ballance
}

func (account *Account) SetBallance(ballance float32) {
	account.ballance = ballance
}

func (account Account) GetCreatedAt() time.Time {
	return account.created_at
}
