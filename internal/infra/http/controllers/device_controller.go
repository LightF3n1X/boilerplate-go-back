package controllers

import (
	"errors"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/upper/db/v4"

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

		room, err := c.roomService.FindById(*device.RoomId)
		if err != nil && !errors.Is(err, db.ErrNoMoreRows) {
			log.Printf("DeviceController.Save(c.roomService.Find): %s", err)
			InternalServerError(w, err)
			return
		}

		if room.Id == 0 {
			err = errors.New("room not found")
			BadRequest(w, err)
			return
		}

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

func (c DeviceController) FindById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		device, err := requests.Bind(r, requests.UpdateDeviceRequest{}, domain.Device{})
		if err != nil {
			log.Printf("DeviceController.FindById(requests.Bind): %s", err)
			BadRequest(w, errors.New("invalid request body"))
			return
		}

		oldDevice := r.Context().Value(DeviceKey).(domain.Device)
		device, err = c.deviceService.Update(oldDevice, device)
		if err != nil {
			log.Printf("DeviceController.FindById(c.deviceService.Update): %s", err)
			InternalServerError(w, err)
			return
		}

		var deviceDto resources.DeviceDto
		deviceDto = deviceDto.DomainToDto(device)
		Success(w, deviceDto)
	}
}

func (c DeviceController) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		device, err := requests.Bind(r, requests.UpdateDeviceRequest{}, domain.Device{})
		if err != nil {
			log.Printf("DeviceController.Update(requests.Bind): %s", err)
			BadRequest(w, errors.New("invalid request body"))
			return
		}

		oldDevice := r.Context().Value(DeviceKey).(domain.Device)
		device, err = c.deviceService.Update(oldDevice, device)
		if err != nil {
			log.Printf("DeviceController.Update(c.deviceService.Update): %s", err)
			InternalServerError(w, err)
			return
		}

		var deviceDto resources.DeviceDto
		deviceDto = deviceDto.DomainToDto(device)
		Success(w, deviceDto)
	}
}

func (c DeviceController) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		device := r.Context().Value(DeviceKey).(domain.Device)

		err := c.deviceService.Delete(device.Id)
		if err != nil {
			log.Printf("DeviceController.Delete(c.deviceService.Delete): %s", err)
			InternalServerError(w, err)
			return
		}
		noContent(w)
	}
}
