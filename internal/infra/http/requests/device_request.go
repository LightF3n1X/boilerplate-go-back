package requests

import (
	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
)

type DeviceRequest struct {
	RoomID           uint64                `json:"room_id" validate:"required"`
	SerialNumber     string                `json:"serial_numbers" validate:"required"`
	Characteristics  *string               `json:"characteristics"`
	Category         domain.DeviceCategory `json:"category" validate:"oneof=ACTUATOR SENSOR"`
	Units            *string               `json:"units"`
	PowerConsumption *int                  `json:"power_consumption"`
}

type UpdateDeviceRequest struct {
	RoomID           *uint64                `json:"room_id"`
	SerialNumber     *string                `json:"serial_numbers"`
	Characteristics  *string                `json:"characteristics"`
	Category         *domain.DeviceCategory `json:"category" validate:"oneof=ACTUATOR SENSOR"`
	Units            *string                `json:"units"`
	PowerConsumption *int                   `json:"power_consumption"`
}

func (r DeviceRequest) ToDomainModel() (interface{}, error) {
	return domain.Device{
		RoomId:           &r.RoomID,
		SerialNumber:     r.SerialNumber,
		Characteristics:  r.Characteristics,
		Category:         domain.DeviceCategory(r.Category),
		Units:            r.Units,
		PowerConsumption: r.PowerConsumption,
	}, nil
}

func (r UpdateDeviceRequest) ToDomainModel() (interface{}, error) {
	var (
		roomId           *uint64
		serialNumber     string
		characteristics  *string
		category         domain.DeviceCategory
		units            *string
		powerConsumption *int
	)

	if r.RoomID != nil {
		roomId = r.RoomID
	}
	if r.SerialNumber != nil {
		serialNumber = *r.SerialNumber
	}
	if r.Characteristics != nil {
		characteristics = r.Characteristics
	}
	if r.Category != nil {
		category = *r.Category
	}
	if r.Units != nil {
		units = r.Units
	}
	if r.PowerConsumption != nil {
		powerConsumption = r.PowerConsumption
	}

	return domain.Device{
		RoomId:           roomId,
		SerialNumber:     serialNumber,
		Characteristics:  characteristics,
		Category:         category,
		Units:            units,
		PowerConsumption: powerConsumption,
	}, nil
}
