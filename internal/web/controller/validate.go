package controller

import (
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/remnv/go-boiler/internal/web/response"
)

func validOrFail(c *gin.Context, v validation.Validatable) error {
	logger := getLogger(c)
	err := v.Validate()
	if err == nil {
		return nil
	}
	logger.WithError(err).Warning("Validation error")
	internalError, isInternalError := err.(validation.InternalError)
	if isInternalError {
		response.Error(c, response.ErrBadRequest, internalError.Error())
	}
	response.ErrorWithPayload(c, response.ErrValidation, "", err)
	return err
}
