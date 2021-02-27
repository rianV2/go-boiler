package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	ErrorResponse struct {
		Code     string      `json:"code,omitempty"`
		Error    string      `json:"error,omitempty"`
		Message  string      `json:"error_message,omitempty"`
		Payload  interface{} `json:"payload,omitempty"`
		HttpCode int         `json:"-"`
	}

	SuccessResponse struct {
		Success bool `json:"success" default:"true"`
	}

	Message struct {
		Success bool        `json:"success"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}
)

const (
	E_DUPLICATE_KEY = "duplicate_key"
	E_NOT_FOUND     = "not_found"
	E_UNAUTHORIZED  = "unauthorized"
	E_BAD_REQUEST   = "bad_request"
	E_SERVER_ERROR  = "server_error"
)

var (
	SuccessOK = SuccessResponse{
		Success: true,
	}

	ErrNotFound = ErrorResponse{
		Error:    E_NOT_FOUND,
		Message:  "Entry not found",
		HttpCode: http.StatusNotFound,
	}
	ErrBadRequest = ErrorResponse{
		Error:    E_BAD_REQUEST,
		Message:  "Bad request",
		HttpCode: http.StatusBadRequest,
	}
	ErrUnauthorized = ErrorResponse{
		Error:    E_UNAUTHORIZED,
		Message:  "Unauthorized, please login",
		HttpCode: http.StatusUnauthorized,
	}
	ErrDuplicate = ErrorResponse{
		Error:    E_DUPLICATE_KEY,
		Message:  "Created value already exists",
		HttpCode: http.StatusConflict,
	}
	ErrValidation = ErrorResponse{
		Error:    E_BAD_REQUEST,
		Message:  "Invalid parameters or payload",
		HttpCode: http.StatusUnprocessableEntity,
	}
	ErrServerError = ErrorResponse{
		Error:    E_SERVER_ERROR,
		Message:  "Something bad happened",
		HttpCode: http.StatusInternalServerError,
	}
)

func Error(c *gin.Context, err ErrorResponse, msg string) {
	ErrorWithPayload(c, err, msg, nil)
}

func ErrorWithPayload(c *gin.Context, err ErrorResponse, msg string, payload interface{}) {
	c.Writer.Header().Del("content-type")
	if msg != "" {
		err.Message = msg
	}
	if payload != nil {
		err.Payload = payload
	}
	status := http.StatusBadRequest
	if err.HttpCode != 0 {
		status = err.HttpCode
	}
	c.JSON(status, err)
}
func Success(c *gin.Context) {
	SuccessWithPayload(c, nil)
}

func SuccessWithPayload(c *gin.Context, payload interface{}) {
	if payload != nil {
		c.JSON(http.StatusOK, payload)
		return
	}
	c.JSON(http.StatusOK, SuccessOK)
}
