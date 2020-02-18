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
