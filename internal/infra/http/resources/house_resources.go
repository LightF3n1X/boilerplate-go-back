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
	Rooms       []RoomDto  `json:"rooms,omitempty"`
	CreatedDate time.Time  `json:"createdDate"`
	UpdatedDate time.Time  `json:"updatedDate"`
	DeleteDate  *time.Time `json:"deletedDate,omitempty"`
}

func (d HouseDto) DomainToDtoCollection(houses []domain.House) []HouseDto {
	hs := make([]HouseDto, len(houses))
	for i, h := range houses {
		hs[i] = d.DomainToDto(h)
	}
	return hs
}

func (d HouseDto) DomainToDto(h domain.House) HouseDto {
	rooms := RoomDto{}.DomainToDtoCollection(h.Rooms)
	return HouseDto{
		Id:          h.Id,
		UserId:      h.UserId,
		Name:        h.Name,
		Description: h.Description,
		City:        h.City,
		Adress:      h.Address,
		Lat:         h.Lat,
		Lon:         h.Lon,
		Rooms:       rooms,
		CreatedDate: h.CreatedDate,
		UpdatedDate: h.UpdatedDate,
		DeleteDate:  h.DeletedDate,
	}
}
