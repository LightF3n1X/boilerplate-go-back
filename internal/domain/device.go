package domain

import "time"

type Device struct {
	Id               uint64
	HouseId          *uint64
	RoomId           *uint64
	UUID             string
	SerialNumber     string
	Characteristics  *string
	Category         DeviceCategory
	Units            *string
	PowerConsumption *int
	CreatedDate      time.Time
	UpdatedDate      time.Time
	DeletedDate      *time.Time
}

type DeviceCategory string

const (
	SensorType   DeviceCategory = "SENSOR"
	ActuatorType DeviceCategory = "ACTUATOR"
)
