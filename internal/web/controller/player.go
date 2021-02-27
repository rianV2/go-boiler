package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/remnv/go-boiler/internal/config"
	"github.com/remnv/go-boiler/internal/usecase"
	"github.com/remnv/go-boiler/internal/web/middleware"
	"github.com/remnv/go-boiler/internal/web/request"
	"github.com/remnv/go-boiler/internal/web/response"
)

type Player struct {
	config  config.Config
	useCase usecase.Player
}

func NewPlayer(cfg config.Config, useCase usecase.Player) *Player {
	return &Player{cfg, useCase}
}

func (p *Player) Create(c *gin.Context) {
	// auth
	user, err := middleware.GetJWTData(c)
	if err != nil {
		response.Error(c, response.ErrUnauthorized, err.Error())
		return
	}

	// validation
	var req request.PlayerCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, response.ErrBadRequest, err.Error())
		return
	}
	err = validOrFail(c, req)
	if err != nil {
		return
	}

	// action
	player, err := req.ToModel()
	player.UserId = &user.ID

	addPlayer, err := p.useCase.Create(*player)
	if err != nil {
		sendErrorResponse(c, err, nil)
		return
	}

	c.JSON(http.StatusOK, addPlayer)
}

func (p *Player) SetUseCase(useCase usecase.Player) {
	p.useCase = useCase
}
