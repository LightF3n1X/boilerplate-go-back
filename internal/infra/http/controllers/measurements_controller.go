package controllers

import "github.com/BohdanBoriak/boilerplate-go-back/internal/app"

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
