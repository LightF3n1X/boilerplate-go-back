package controllers

import (
	"errors"
	"log"
	"net/http"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/app"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/http/requests"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/http/resources"
)

type RoomController struct {
	roomService app.RoomService
}

func NewRoomController(rs app.RoomService) RoomController {
	return RoomController{
		roomService: rs,
	}
}

func (c RoomController) Save() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		room, err := requests.Bind(r, requests.RoomRequest{}, domain.Room{})
		if err != nil {
			log.Printf("RoomController.Save(requests.Bind): %s", err)
			BadRequest(w, errors.New("invalid request body"))
			return
		}

		house := r.Context().Value(HouseKey).(domain.House)
		room.HouseId = house.Id

		room, err = c.roomService.Save(room)
		if err != nil {
			log.Printf("RoomController.Save(c.roomService.Save): %s", err)
			InternalServerError(w, err)
			return
		}

		var roomDto resources.RoomDto
		roomDto = roomDto.DomainToDto(room)
		Success(w, roomDto)
	}
}

func (c RoomController) Find() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		room := r.Context().Value(RoomKey).(domain.Room)

		var roomDto resources.RoomDto
		roomDto = roomDto.DomainToDto(room)
		Success(w, roomDto)
	}
}

func (c RoomController) FindByHouseId() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		house := r.Context().Value(HouseKey).(domain.House)

		rooms, err := c.roomService.FindByHouseId(house.Id)
		if err != nil {
			log.Printf("RoomController.FindByHouseId(c.roomService.FindByHouseId): %s", err)
			InternalServerError(w, err)
			return
		}

		var roomsDto []resources.RoomDto
		for _, room := range rooms {
			var roomDto resources.RoomDto
			roomDto = roomDto.DomainToDto(room)
			roomsDto = append(roomsDto, roomDto)
		}
		Success(w, roomsDto)
	}
}

func (c RoomController) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		room, err := requests.Bind(r, requests.UpdateRoomRequest{}, domain.Room{})
		if err != nil {
			log.Printf("RoomController.Update(requests.Bind): %s", err)
			BadRequest(w, errors.New("invalid request body"))
			return
		}

		oldRoom := r.Context().Value(RoomKey).(domain.Room)
		room, err = c.roomService.Update(oldRoom, room)
		if err != nil {
			log.Printf("RoomController.Update(c.roomService.Update): %s", err)
			InternalServerError(w, err)
			return
		}

		var roomDto resources.RoomDto
		roomDto = roomDto.DomainToDto(room)
		Success(w, roomDto)
	}
}

func (c RoomController) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		room := r.Context().Value(RoomKey).(domain.Room)

		err := c.roomService.Delete(room.Id)
		if err != nil {
			log.Printf("RoomController.Delete(c.roomService.Delete): %s", err)
			InternalServerError(w, err)
			return
		}
		noContent(w)
	}
}
