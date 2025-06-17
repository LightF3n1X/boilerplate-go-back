package database

import (
	"time"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/upper/db/v4"
)

const DevicesTableName = "devices"

type device struct {
	Id               uint64                `db:"id,omitempty"`
	HouseId          *uint64               `db:"house_id"`
	RoomId           *uint64               `db:"room_id"`
	UUID             string                `db:"uuid"`
	SerialNumber     string                `db:"serial_numbers"`
	Characteristics  *string               `db:"characteristics"`
	Category         domain.DeviceCategory `db:"category"`
	Units            *string               `db:"units"`
	PowerConsumption *int                  `db:"power_consumption"`
	CreatedDate      time.Time             `db:"created_date"`
	UpdatedDate      time.Time             `db:"updated_date"`
	DeletedDate      *time.Time            `db:"deleted_date"`
}

type DeviceRepository interface {
	Save(d domain.Device) (domain.Device, error)
	Find(id uint64) (domain.Device, error)
	FindById(id uint64) (domain.Device, error)
	Update(d domain.Device) (domain.Device, error)
	Delete(id uint64) error
}

type deviceRepository struct {
	coll db.Collection
	sess db.Session
}

func NewDeviceRepository(sess db.Session) DeviceRepository {
	return deviceRepository{
		coll: sess.Collection(DevicesTableName),
		sess: sess,
	}
}

func (r deviceRepository) Save(d domain.Device) (domain.Device, error) {
	ds := r.mapDomainToModel(d)
	ds.CreatedDate = time.Now()
	ds.UpdatedDate = time.Now()

	err := r.coll.InsertReturning(&ds)
	if err != nil {
		return domain.Device{}, err
	}
	d = r.mapModelToDomain(ds)
	return d, nil
}

func (r deviceRepository) Find(id uint64) (domain.Device, error) {
	var ds device
	err := r.coll.Find(db.Cond{"id": id, "deleted_date": nil}).One(&ds)
	if err != nil {
		return domain.Device{}, err
	}
	d := r.mapModelToDomain(ds)
	return d, nil
}

func (r deviceRepository) FindById(id uint64) (domain.Device, error) {
	var ds device
	err := r.coll.Find(db.Cond{"id": id, "deleted_date": nil}).One(&ds)
	if err != nil {
		return domain.Device{}, err
	}
	d := r.mapModelToDomain(ds)
	return d, nil
}

func (r deviceRepository) Update(d domain.Device) (domain.Device, error) {
	ds := r.mapDomainToModel(d)
	ds.UpdatedDate = time.Now()

	err := r.coll.
		Find(db.Cond{
			"id":           ds.Id,
			"deleted_date": nil}).Update(&ds)
	if err != nil {
		return domain.Device{}, err
	}
	d = r.mapModelToDomain(ds)
	return d, nil
}

func (r deviceRepository) Delete(id uint64) error {
	return r.coll.Find(db.Cond{"id": id, "deleted_date": nil}).Update(map[string]interface{}{"deleted_date": time.Now()})
}

func (r deviceRepository) mapDomainToModel(d domain.Device) device {
	return device{
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
		DeletedDate:      d.DeletedDate,
	}
}

func (r deviceRepository) mapModelToDomain(d device) domain.Device {
	return domain.Device{
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
		DeletedDate:      d.DeletedDate,
	}
}
