package usecase

import (
	"github.com/remnv/go-boiler/internal/config"
	"github.com/remnv/go-boiler/internal/model"
	"github.com/remnv/go-boiler/internal/repository"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
)

type Player struct {
	config     config.Config
	playerRepo repository.Player
}

func NewPlayer(config config.Config, playerRepo repository.Player) Player {
	return Player{
		config:     config,
		playerRepo: playerRepo,
	}
}

// Create player if does not exist
func (p Player) Create(player model.Player) (*model.Player, error) {
	logger := logrus.WithField("_sid", uuid.NewV4().String()).WithField("_method", "Create")

	newPl, err := p.playerRepo.Add(player)
	if err != nil {
		logger.WithError(err).Warning("Failed add player")
		return nil, err
	}

	return newPl, err
}
