package gormrepo

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/remnv/go-boiler/internal/model"
	"github.com/segmentio/ksuid"
)

type Player struct {
	ID        *string
	UserId    *string
	Name      *string
	Level     *string
	Job       *string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

func (p *Player) BeforeCreate(scope *gorm.Scope) error {
	if p.ID != nil {
		return nil
	}
	return scope.SetColumn("ID", ksuid.New().String())
}

func (p Player) FromModel(player model.Player) (*Player, error) {
	return &Player{
		ID:        player.ID,
		UserId:    player.UserId,
		Name:      player.Name,
		Level:     player.Level,
		Job:       player.Job,
		CreatedAt: player.CreatedAt,
		UpdatedAt: player.UpdatedAt,
	}, nil
}

func (p Player) ToModel() (*model.Player, error) {
	return &model.Player{
		ID:        p.ID,
		UserId:    p.UserId,
		Name:      p.Name,
		Level:     p.Level,
		Job:       p.Job,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}, nil
}

func (p Player) ToModels(players []Player) (ret []model.Player, err error) {
	for _, v := range players {
		m, err := v.ToModel()
		if err != nil {
			return nil, err
		}
		ret = append(ret, *m)
	}
	return
}
