package resources

import (
	"time"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
)

type MeasurementDto struct {
	Id          uint64     `json:"id"`
	DeviceId    uint64     `json:"device_id"`
	RoomId      *uint64    `json:"room_id,omitempty"`
	Value       uint64     `json:"value"`
	CreatedDate time.Time  `json:"createdDate"`
	UpdatedDate time.Time  `json:"updatedDate"`
	DeletedDate *time.Time `json:"deletedDate,omitempty"`
}

func (m MeasurementDto) DomainToDtoCollection(measurements []domain.Measurement) []MeasurementDto {
	ms := make([]MeasurementDto, len(measurements))
	for i, measurement := range measurements {
		ms[i] = m.DomainToDto(measurement)
	}
	return ms
}

func (m MeasurementDto) DomainToDto(measurement domain.Measurement) MeasurementDto {
	return MeasurementDto{
		Id:          m.Id,
		DeviceId:    m.DeviceId,
		RoomId:      m.RoomId,
		Value:       m.Value,
		CreatedDate: m.CreatedDate,
		UpdatedDate: m.UpdatedDate,
		DeletedDate: m.DeletedDate,
	}
}
