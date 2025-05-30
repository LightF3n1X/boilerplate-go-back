package requests

import (
	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
)

type DeviceRequest struct {
	RoomID           uint64                `json:"room_id" validate:"required"`
	SerialNumber     string                `json:"serial_number" validate:"required"`
	Characteristics  *string               `json:"characteristics"`
	Category         domain.DeviceCategory `json:"category" validate:"oneof=ACTUATOR SENSOR"`
	Units            *string               `json:"units"`
	PowerConsumption *int                  `json:"power_consumption"`
}

/*type UpdateDeviceRequest struct {
	RoomID           *uint64                `json:"room_id"`
	SerialNumber     *string                `json:"serial_number"`
	Characteristics  *string                `json:"characteristics"`
	Category         *domain.DeviceCategory `json:"category"`
	Units            *string                `json:"units"`
	PowerConsumption *int                   `json:"power_consumption"`
}*/

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
