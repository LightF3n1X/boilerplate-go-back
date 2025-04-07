package resources

import (
	"time"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
)

type HouseDto struct {
	Id          uint64     `json:"id"`
	UserId      uint64     `json:"userId"`
	Name        string     `json:"name"`
	Description *string    `json:"description,omitempty"`
	City        string     `json:"city"`
	Adress      string     `json:"address"`
	Lat         float64    `json:"lat"`
	Lon         float64    `json:"lon"`
	CreatedDate time.Time  `json:"createdDate"`
	UpdatedDate time.Time  `json:"updatedDate"`
	DeleteDate  *time.Time `json:"deletedDate,omitempty"`
}

func (d HouseDto) DomainToDto(h domain.House) HouseDto {
	return HouseDto{
		Id:          h.Id,
		UserId:      h.UserId,
		Name:        h.Name,
		Description: h.Description,
		City:        h.City,
		Adress:      h.Adress,
		Lat:         h.Lat,
		Lon:         h.Lon,
		CreatedDate: h.CreatedDate,
		UpdatedDate: h.UpdatedDate,
		DeleteDate:  h.DeleteDate,
	}
}
