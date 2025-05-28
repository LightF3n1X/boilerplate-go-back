package requests

import (
	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
)

type DeviceRequest struct {
	SerialNumber     string                `json:"serial_number" validate:"required"`
	Characteristics  *string               `json:"characteristics"`
	Category         domain.DeviceCategory `json:"category" validate:"required"`
	Units            *string               `json:"units"`
	PowerConsumption *int                  `json:"power_consumption"`
}

type UpdateDeviceRequest struct {
	SerialNumber     *string                `json:"serial_number"`
	Characteristics  *string                `json:"characteristics"`
	Category         *domain.DeviceCategory `json:"category"`
	Units            *string                `json:"units"`
	PowerConsumption *int                   `json:"power_consumption"`
}

func (r DeviceRequest) ToDomainModel() (interface{}, error) {
	return domain.Device{
		SerialNumber:     r.SerialNumber,
		Characteristics:  r.Characteristics,
		Category:         r.Category,
		Units:            r.Units,
		PowerConsumption: r.PowerConsumption,
	}, nil
}

func (r UpdateDeviceRequest) ToDomainModel() (interface{}, error) {
	var (
		serialNumber     string
		characteristics  *string
		category         *domain.DeviceCategory
		units            *string
		powerConsumption *int
	)
	if r.SerialNumber != nil {
		serialNumber = *r.SerialNumber
	}
	if r.Characteristics != nil {
		characteristics = r.Characteristics
	}
	if r.Category != nil {
		category = r.Category
	}
	if r.Units != nil {
		units = r.Units
	}
	if r.PowerConsumption != nil {
		powerConsumption = r.PowerConsumption
	}

	var categoryValue domain.DeviceCategory
	if category != nil {
		categoryValue = *category
	}
	return domain.Device{
		SerialNumber:     serialNumber,
		Characteristics:  characteristics,
		Category:         categoryValue,
		Units:            units,
		PowerConsumption: powerConsumption,
	}, nil
}
