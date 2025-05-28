package resources

import (
	"time"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
)

type DeviceDto struct {
	Id               uint64                `json:"id"`
	HouseId          uint64                `json:"houseId"`
	RoomId           uint64                `json:"roomId"`
	UUID             string                `json:"uuid"`
	SerialNumber     string                `json:"serialNumber"`
	Characteristics  *string               `json:"characteristics,omitempty"`
	Category         domain.DeviceCategory `json:"category"`
	Units            *string               `json:"units,omitempty"`
	PowerConsumption *int                  `json:"powerConsumption,omitempty"`
	CreatedDate      time.Time             `json:"createdDate"`
	UpdatedDate      time.Time             `json:"updatedDate"`
	DeleteDate       *time.Time            `json:"deletedDate,omitempty"`
}

func (d DeviceDto) DomainToDtoCollection(devices []domain.Device) []DeviceDto {
	ds := make([]DeviceDto, len(devices))
	for i, device := range devices {
		ds[i] = d.DomainToDto(device)
	}
	return ds
}

func (d DeviceDto) DomainToDto(device domain.Device) DeviceDto {
	return DeviceDto{
		Id:               d.Id,
		HouseId:          d.HouseId,
		RoomId:           d.RoomId,
		UUID:             d.UUID,
		SerialNumber:     d.SerialNumber,
		Characteristics:  d.Characteristics,
		Category:         d.Category,
		Units:            d.Units,
		PowerConsumption: d.PowerConsumption,
		CreatedDate:      d.CreatedDate,
		UpdatedDate:      d.UpdatedDate,
		DeleteDate:       d.DeleteDate,
	}
}
