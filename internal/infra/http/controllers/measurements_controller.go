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

type MeasurementsController struct {
	measurementsService app.MeasurementsService
	roomService         app.RoomService
	deviceService       app.DeviceService
}

func NewMeasurementsController(ms app.MeasurementsService, rs app.RoomService, ds app.DeviceService) MeasurementsController {
	return MeasurementsController{
		measurementsService: ms,
		roomService:         rs,
		deviceService:       ds,
	}
}

func (c MeasurementsController) Save() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		measurement, err := requests.Bind(r, requests.MeasurementRequest{}, domain.Measurement{})
		if err != nil {
			log.Printf("MeasurementsController.Save(requests.Bind): %s", err)
			BadRequest(w, errors.New("invalid request body"))
			return
		}
		room, err := c.roomService.FindById(*measurement.RoomId)
		if err != nil {
			log.Printf("MeasurementsController.Save(c.roomService.FindById): %s", err)
			InternalServerError(w, err)
			return
		}
		if room.Id == 0 {
			err = errors.New("room not found")
			BadRequest(w, err)
			return
		}
		device, err := c.deviceService.FindById(measurement.DeviceId)
		if err != nil {
			log.Printf("MeasurementsController.Save(c.deviceService.Find): %s", err)
			InternalServerError(w, err)
			return
		}
		if device.Id == 0 {
			err = errors.New("device not found")
			BadRequest(w, err)
			return
		}

		measurement.RoomId = &room.Id
		measurement.DeviceId = &device.Id

		measurement, err = c.measurementsService.Save(measurement)
		if err != nil {
			log.Printf("MeasurementsController.Save(c.measurementsService.Save): %s", err)
			InternalServerError(w, err)
			return
		}

		var measurementDto resources.MeasurementDto
		measurementDto = measurementDto.DomainToDto(measurement)
		Success(w, measurementDto)
	}
}
