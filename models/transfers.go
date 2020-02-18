package models

import "time"

type Transfer struct {
	Id                     uint
	Account_origin_id      uint64
	Account_destination_id uint64
	Amount                 float32
	Created_at             time.Time
}
