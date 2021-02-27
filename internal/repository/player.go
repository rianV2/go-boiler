package repository

import (
	"github.com/remnv/go-boiler/internal/model"
)

type Player interface {
	Add(player model.Player) (*model.Player, error)
}
