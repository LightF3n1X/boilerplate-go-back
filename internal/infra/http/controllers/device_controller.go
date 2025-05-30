package controllers

import (
	"errors"
	"log"
	"net/http"

	"github.com/google/uuid"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/app"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/http/requests"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/http/resources"
)

type DeviceController struct {
	deviceService app.DeviceService
	roomService   app.RoomService
}

func NewDeviceController(ds app.DeviceService, rs app.RoomService) DeviceController {
	return DeviceController{
		deviceService: ds,
		roomService:   rs,
	}
}

func (c DeviceController) Save() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		device, err := requests.Bind(r, requests.DeviceRequest{}, domain.Device{})
		if err != nil {
			log.Printf("DeviceController.Save(requests.Bind): %s", err)
			BadRequest(w, errors.New("invalid request body"))
			return
		}
		device.UUID = uuid.New().String()

		room, err := c.roomService.Find(*device.RoomId)
		if err != nil {
			log.Printf("DeviceController.Save(c.roomService.Find): %s", err)
			InternalServerError(w, err)
			return
		}

		room := r.Context().Value(RoomKey).(domain.Room)
		device.RoomId = &room.Id

		device, err = c.deviceService.Save(device)
		if err != nil {
			log.Printf("DeviceController.Save(c.deviceService.Save): %s", err)
			InternalServerError(w, err)
			return
		}

		var deviceDto resources.DeviceDto
		deviceDto = deviceDto.DomainToDto(device)
		Success(w, deviceDto)
	}
}

func (c DeviceController) Find() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		device := r.Context().Value(DeviceKey).(domain.Device)

		var deviceDto resources.DeviceDto
		deviceDto = deviceDto.DomainToDto(device)
		Success(w, deviceDto)
	}
}
