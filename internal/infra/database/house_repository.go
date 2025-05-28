package database

import (
	"time"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/upper/db/v4"
)

const HousesTableName = "houses"

type house struct {
	Id          uint64     `db:"id,omitempty"`
	UserId      uint64     `db:"user_id"`
	Name        string     `db:"name"`
	Description *string    `db:"description"`
	City        string     `db:"city"`
	Address     string     `db:"address"`
	Lat         float64    `db:"lat"`
	Lon         float64    `db:"lon"`
	CreatedDate time.Time  `db:"created_date"`
	UpdatedDate time.Time  `db:"updated_date"`
	DeletedDate *time.Time `db:"deleted_date"`
}

type HouseRepository interface {
	Save(h domain.House) (domain.House, error)
	FindList(uId uint64) ([]domain.House, error)
	Find(id uint64) (domain.House, error)
	Update(h domain.House) (domain.House, error)
	Delete(id uint64) error
}

type houseRepository struct {
	coll db.Collection
	sess db.Session
}

func NewHouseRepository(sess db.Session) HouseRepository {
	return houseRepository{
		coll: sess.Collection(HousesTableName),
		sess: sess,
	}
}

func (r houseRepository) Save(h domain.House) (domain.House, error) {
	hs := r.mapDomainToModel(h)
	hs.CreatedDate = time.Now()
	hs.UpdatedDate = time.Now()

	err := r.coll.InsertReturning(&hs)
	if err != nil {
		return domain.House{}, err
	}

	h = r.mapModelToDomain(hs)
	return h, nil
}

func (r houseRepository) FindList(uId uint64) ([]domain.House, error) {
	var houses []house
	err := r.coll.
		Find(db.Cond{
			"user_id":      uId,
			"deleted_date": nil}).All(&houses)
	if err != nil {
		return nil, err
	}

	hs := r.mapModelToDomainCollection(houses)
	return hs, nil
}

func (r houseRepository) Find(id uint64) (domain.House, error) {
	var h house
	err := r.coll.
		Find(db.Cond{
			"id":           id,
			"deleted_date": nil}).One(&h)
	if err != nil {
		return domain.House{}, err
	}

	hs := r.mapModelToDomain(h)
	return hs, nil
}

func (r houseRepository) Update(h domain.House) (domain.House, error) {
	hs := r.mapDomainToModel(h)
	hs.UpdatedDate = time.Now()
	err := r.coll.
		Find(db.Cond{
			"id":           hs.Id,
			"deleted_date": nil}).Update(&hs)
	if err != nil {
		return domain.House{}, err
	}
	return r.mapModelToDomain(hs), nil
}

func (r houseRepository) Delete(id uint64) error {
	return r.coll.Find(db.Cond{"id": id, "deleted_date": nil}).Update(map[string]interface{}{"deleted_date": time.Now()})
}

func (r houseRepository) mapDomainToModel(d domain.House) house {
	return house{
		Id:          d.Id,
		UserId:      d.UserId,
		Name:        d.Name,
		Description: d.Description,
		City:        d.City,
		Address:     d.Address,
		Lat:         d.Lat,
		Lon:         d.Lon,
		CreatedDate: d.CreatedDate,
		UpdatedDate: d.UpdatedDate,
		DeletedDate: d.DeletedDate,
	}
}

func (r houseRepository) mapModelToDomain(d house) domain.House {
	return domain.House{
		Id:          d.Id,
		UserId:      d.UserId,
		Name:        d.Name,
		Description: d.Description,
		City:        d.City,
		Address:     d.Address,
		Lat:         d.Lat,
		Lon:         d.Lon,
		CreatedDate: d.CreatedDate,
		UpdatedDate: d.UpdatedDate,
		DeletedDate: d.DeletedDate,
	}
}

func (r houseRepository) mapModelToDomainCollection(houses []house) []domain.House {
	hs := make([]domain.House, len(houses))
	for i, h := range houses {
		hs[i] = r.mapModelToDomain(h)
	}

	return hs
}
