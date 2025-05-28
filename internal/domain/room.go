package domain

import "time"

type Room struct {
	Id          uint64
	Name        string
	HouseId     uint64
	Description *string
	CreatedDate time.Time
	UpdatedDate time.Time
	DeletedDate *time.Time
}
