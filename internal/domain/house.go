package domain

import "time"

type House struct {
	Id          uint64
	UserId      uint64
	Name        string
	Description *string
	City        string
	Address     string
	Lat         float64
	Lon         float64
	Rooms       []Room
	CreatedDate time.Time
	UpdatedDate time.Time
	DeletedDate *time.Time
}
