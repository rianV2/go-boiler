package gormrepo

import (
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/remnv/go-boiler/internal/helpers"
)

type GenericStorager interface {
	Create(data interface{}) error
	Update(data interface{}) error
	Delete(filter interface{}) error
	Fetch(where interface{}, out interface{}, limit int, offset int) error
	FetchWithAssoc(where interface{}, out interface{}, limit int, offset int, assocs ...string) error
	FetchLike(field string, query string, out interface{}, limit string, offset string) error
	Get(outFilter interface{}) error
	GetWithAssociations(outFilter interface{}, associations ...string) error
}

type generic struct {
	db *gorm.DB
}

func NewGeneric(db *gorm.DB) GenericStorager {
	return &generic{db}
}

func (s *generic) Create(data interface{}) error {
	if err := s.db.Create(data).Error; err != nil {
		return err
	}
	return nil
}
func (s *generic) Delete(data interface{}) error {
	if !helpers.IsZero(data) {
		r := s.db.Where(data).Delete(data)
		if r.Error != nil {
			return r.Error
		}
		if r.RowsAffected < 1 {
			return errors.New(ENoRowsAffected)
		}
	}
	return nil
}
func (s *generic) Fetch(where interface{}, out interface{}, limit int, offset int) error {
	r := s.db.Where(where).Limit(limit).Offset(offset).Find(out)
	if r.Error != nil {
		return r.Error
	}
	return nil
}

func (s *generic) FetchWithAssoc(where interface{}, out interface{}, limit int, offset int, associations ...string) error {
	q := s.db.Where(where).Limit(limit).Offset(offset)
	for _, assoc := range associations {
		q = q.Preload(assoc)
	}
	r := q.Find(out)
	if r.Error != nil {
		return r.Error
	}
	return nil
}

func (s *generic) Get(data interface{}) error {
	if !helpers.IsZero(data) {
		r := s.db.Where(data).First(data)
		if r.Error != nil {
			return r.Error
		}
		return nil
	}
	return nil
}

func (s *generic) GetWithAssociations(data interface{}, associations ...string) error {
	if helpers.IsZero(data) {
		return nil
	}

	q := s.db.Model(data)
	for _, assoc := range associations {
		q = q.Preload(assoc)
	}
	r := q.First(data)
	if r.Error != nil {
		return r.Error
	}

	return nil
}

func (s *generic) Update(data interface{}) error {
	r := s.db.Model(data).Updates(data)
	if r.Error != nil {
		return r.Error
	}
	if r.RowsAffected < 1 {
		return errors.New(ENoRowsAffected)
	}
	return nil
}

func (s *generic) FetchLike(field string, query string, out interface{}, limit string, offset string) error {
	field = fmt.Sprintf("%s LIKE ?", field)
	r := s.db.Where(field, query).Limit(limit).Offset(offset).Find(out)
	if r.Error != nil {
		return r.Error
	}
	if r.RowsAffected < 1 {
		return errors.New(ENoRowsAffected)
	}
	return nil
}
