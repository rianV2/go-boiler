package gormrepo

import (
	"github.com/jinzhu/gorm"
	"github.com/remnv/go-boiler/internal/model"
)

type PlayerGorm struct {
	GenericStorager
	db *gorm.DB
}

func NewPlayer(db *gorm.DB) *PlayerGorm {
	newGeneric := NewGeneric(db)
	return &PlayerGorm{newGeneric, db}
}

func (p *PlayerGorm) Add(player model.Player) (*model.Player, error) {
	pl, err := Player{}.FromModel(player)
	if err != nil {
		return nil, err
	}
	err = p.Create(&pl)
	if err != nil {
		// if strings.Contains(strings.ToLower(err.Error()), "duplicate") {
		// 	return nil, usecase.NewError(err.Error(), usecase.ErrorDuplicate)
		// }
		return nil, err
	}
	return pl.ToModel()
}
