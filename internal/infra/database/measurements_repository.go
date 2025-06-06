package database

import (
	"time"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/upper/db/v4"
)

const MeasurementsTableName = "measurements"

type measurement struct {
	Id          uint64     `db:"id,omitempty"`
	DeviceId    uint64     `db:"device_id"`
	RoomId      *uint64    `db:"room_id"`
	Value       uint64     `db:"value"`
	CreatedDate time.Time  `db:"created_date,omitempty"`
	UpdatedDate time.Time  `db:"updated_date,omitempty"`
	DeletedDate *time.Time `db:"deleted_date,omitempty"`
}

type MeasurementsRepository interface {
}

type measurementsRepository struct {
	coll db.Collection
	sess db.Session
}

func NewMeasurementsRepository(sess db.Session) MeasurementsRepository {
	return measurementsRepository{
		coll: sess.Collection(MeasurementsTableName),
		sess: sess,
	}
}

func (r measurementsRepository) mapDomainToModel(m domain.Measurement) measurement {
	return measurement{
		Id:          m.Id,
		DeviceId:    m.DeviceId,
		RoomId:      m.RoomId,
		Value:       m.Value,
		CreatedDate: m.CreatedDate,
		UpdatedDate: m.UpdatedDate,
		DeletedDate: m.DeletedDate,
	}
}

func (r measurementsRepository) mapModelToDomain(m measurement) domain.Measurement {
	return domain.Measurement{
		Id:          m.Id,
		DeviceId:    m.DeviceId,
		RoomId:      m.RoomId,
		Value:       m.Value,
		CreatedDate: m.CreatedDate,
		UpdatedDate: m.UpdatedDate,
		DeletedDate: m.DeletedDate,
	}
}
