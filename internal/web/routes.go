package web

import (
	"github.com/gin-gonic/gin"
	"github.com/remnv/go-boiler/internal/web/middleware"
)

func (h *httpServer) setupRouting() {

	router := h.engine

	router.GET("/ping", func(context *gin.Context) {
		context.String(200, "Ok")
	})

	router.Use(middleware.TokenMiddleware())
	{
		router.POST("/players", h.controllers.player.Create)
	}
}
