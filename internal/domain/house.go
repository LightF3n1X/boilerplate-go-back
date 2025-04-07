package domain

import "time"

type House struct {
	Id          uint64
	UserId      uint64
	Name        string
	Description *string
	City        string
	Adress      string
	Lat         float64
	Lon         float64
	CreatedDate time.Time
	UpdatedDate time.Time
	DeleteDate  *time.Time
}
