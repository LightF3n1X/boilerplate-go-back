package database

import (
	"time"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/upper/db/v4"
)

const RoomsTableName = "rooms"

type room struct {
	Id          uint64     `db:"id,omitempty"`
	Name        string     `db:"name"`
	HouseId     uint64     `db:"house_id"`
	Description *string    `db:"description"`
	CreatedDate time.Time  `db:"created_date"`
	UpdatedDate time.Time  `db:"updated_date"`
	DeletedDate *time.Time `db:"deleted_date"`
}

type RoomRepository interface {
	Save(r domain.Room) (domain.Room, error)
	Find(id uint64) (domain.Room, error)
	FindByHouseId(hid uint64) ([]domain.Room, error)
	Update(r domain.Room) (domain.Room, error)
	Delete(id uint64) error
}

type roomRepository struct {
	coll db.Collection
	sess db.Session
}

func NewRoomRepository(sess db.Session) RoomRepository {
	return roomRepository{
		coll: sess.Collection(RoomsTableName),
		sess: sess,
	}
}

func (r roomRepository) Save(rm domain.Room) (domain.Room, error) {
	rs := r.mapDomainToModel(rm)
	rs.CreatedDate = time.Now()
	rs.UpdatedDate = time.Now()

	err := r.coll.InsertReturning(&rs)
	if err != nil {
		return domain.Room{}, err
	}

	rm = r.mapModelToDomain(rs)
	return rm, nil
}

func (r roomRepository) Find(id uint64) (domain.Room, error) {
	var rs room
	err := r.coll.Find(db.Cond{"id": id, "deleted_date": nil}).One(&rs)
	if err != nil {
		return domain.Room{}, err
	}

	rss := r.mapModelToDomain(rs)
	return rss, nil
}
func (r roomRepository) FindByHouseId(hid uint64) ([]domain.Room, error) {
	rs := []room{}
	err := r.coll.Find(db.Cond{"house_id": hid, "deleted_date": nil}).All(&rs)
	if err != nil {
		return nil, err
	}

	rooms := make([]domain.Room, len(rs))
	for i, room := range rs {
		rooms[i] = r.mapModelToDomain(room)
	}
	return rooms, nil
}

func (r roomRepository) Update(rm domain.Room) (domain.Room, error) {
	rs := r.mapDomainToModel(rm)
	rs.UpdatedDate = time.Now()
	err := r.coll.
		Find(db.Cond{
			"id":           rm.Id,
			"deleted_date": nil}).Update(&rs)
	if err != nil {
		return domain.Room{}, err
	}
	return r.mapModelToDomain(rs), nil
}

func (r roomRepository) Delete(id uint64) error {
	return r.coll.Find(db.Cond{"id": id, "deleted_date": nil}).Update(map[string]interface{}{"deleted_date": time.Now()})
}

func (r roomRepository) mapDomainToModel(rm domain.Room) room {
	return room{
		Id:          rm.Id,
		Name:        rm.Name,
		HouseId:     rm.HouseId,
		Description: rm.Description,
		CreatedDate: rm.CreatedDate,
		UpdatedDate: rm.UpdatedDate,
		DeletedDate: rm.DeletedDate,
	}
}
func (r roomRepository) mapModelToDomain(rs room) domain.Room {
	return domain.Room{
		Id:          rs.Id,
		Name:        rs.Name,
		HouseId:     rs.HouseId,
		Description: rs.Description,
		CreatedDate: rs.CreatedDate,
		UpdatedDate: rs.UpdatedDate,
		DeletedDate: rs.DeletedDate,
	}
}

func (r roomRepository) mapModelToDomainCollection(rooms []room) []domain.Room {
	hs := make([]domain.Room, len(rooms))
	for i, h := range rooms {
		hs[i] = r.mapModelToDomain(h)
	}
	return hs
}
