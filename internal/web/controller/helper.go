package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/remnv/go-boiler/internal/usecase"
	"github.com/remnv/go-boiler/internal/web/response"
	"github.com/sirupsen/logrus"
)

func sendErrorResponse(c *gin.Context, err error, payload interface{}) {
	logger := getLogger(c)
	logger.WithError(err).Warning("Error when running use case")

	t, isValidType := err.(usecase.Error)
	if !isValidType {
		response.Error(c, response.ErrServerError, err.Error())
		return
	}

	if t.Code == usecase.ErrorDuplicate {
		c.JSON(http.StatusConflict, payload)
		return
	}

	response.Error(c, response.ErrorResponse{
		Code:     strconv.Itoa(t.Code),
		Message:  t.Error(),
		Error:    "",
		HttpCode: t.Code,
	}, t.Error())
}

func getLogger(c *gin.Context) *logrus.Entry {
	return logrus.WithField("requestId", c.GetString("RequestId"))
}
