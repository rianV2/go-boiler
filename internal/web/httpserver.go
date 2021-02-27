package web

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/remnv/go-boiler/internal/config"
	"github.com/remnv/go-boiler/internal/repository/gormrepo"
	"github.com/remnv/go-boiler/internal/usecase"
	"github.com/remnv/go-boiler/internal/web/controller"
	"github.com/remnv/go-boiler/internal/web/middleware"
	"github.com/sirupsen/logrus"
)

type HttpServer interface {
	Start() error
	GetHandler() (http.Handler, error)
}

type httpServer struct {
	db          *gorm.DB
	config      config.Config
	engine      *gin.Engine
	controllers controllers
}

type controllers struct {
	player controller.Player
}

func NewHttpServer(db *gorm.DB, cfg config.Config) HttpServer {
	controllers := controllers{
		*controller.NewPlayer(cfg, usecase.NewPlayer(cfg, gormrepo.NewPlayer(db))),
	}

	gin.SetMode(gin.ReleaseMode)
	if strings.ToLower(cfg.LogLevel) == gin.DebugMode {
		gin.SetMode(gin.DebugMode)
	}

	engine := newGinEngine()

	requestHandler := &httpServer{db, cfg, engine, controllers}
	requestHandler.setupRouting()

	return requestHandler
}

func (h *httpServer) Start() error {
	return h.engine.Run(fmt.Sprintf("%s:%s", h.config.Service.Host, h.config.Service.Port))
}

func (h *httpServer) GetHandler() (http.Handler, error) {
	return h.engine, nil
}

func newGinEngine() *gin.Engine {
	r := gin.New()

	// Add global middlewares
	r.Use(middleware.RequestId())
	r.Use(middleware.LogrusLogger(logrus.StandardLogger()), gin.Recovery())

	return r
}

//-- Builder
type HttpServerBuilder struct {
	db          *gorm.DB
	config      config.Config
	controllers controllers
}

func NewHttpServerBuilder(db *gorm.DB, cfg config.Config) *HttpServerBuilder {
	controllers := controllers{
		*controller.NewPlayer(cfg, usecase.NewPlayer(cfg, gormrepo.NewPlayer(db))),
	}

	return &HttpServerBuilder{db, cfg, controllers}
}

func (h *HttpServerBuilder) SetPlayerUsecase(u usecase.Player) {
	h.controllers.player.SetUseCase(u)
}

func (h *HttpServerBuilder) Build() HttpServer {

	gin.SetMode(gin.ReleaseMode)
	if strings.ToLower(h.config.LogLevel) == gin.DebugMode {
		gin.SetMode(gin.DebugMode)
	}

	engine := newGinEngine()

	requestHandler := &httpServer{h.db, h.config, engine, h.controllers}
	requestHandler.setupRouting()

	return requestHandler
}
