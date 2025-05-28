package resources

import (
	"time"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
)

type RoomDto struct {
	Id          uint64     `json:"id"`
	Name        string     `json:"name"`
	HouseId     uint64     `json:"houseId"`
	Description *string    `json:"description,omitempty"`
	CreatedDate time.Time  `json:"createdDate"`
	UpdatedDate time.Time  `json:"updatedDate"`
	DeleteDate  *time.Time `json:"deletedDate,omitempty"`
}

func (d RoomDto) DomainToDtoCollection(rooms []domain.Room) []RoomDto {
	rs := make([]RoomDto, len(rooms))
	for i, r := range rooms {
		rs[i] = d.DomainToDto(r)
	}
	return rs
}

func (d RoomDto) DomainToDto(r domain.Room) RoomDto {
	return RoomDto{
		Id:          r.Id,
		Name:        r.Name,
		HouseId:     r.HouseId,
		Description: r.Description,
		CreatedDate: r.CreatedDate,
		UpdatedDate: r.UpdatedDate,
		DeleteDate:  r.DeletedDate,
	}
}
