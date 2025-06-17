package requests

import "github.com/BohdanBoriak/boilerplate-go-back/internal/domain"

type MeasurementRequest struct {
	RoomId *uint64 `json:"room_id,omitempty"`
	Value  uint64  `json:"value" validate:"required"`
}

func (r MeasurementRequest) ToDomainModel() (interface{}, error) {
	return domain.Measurement{
		RoomId: r.RoomId,
		Value:  r.Value,
	}, nil
}
