package middleware

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/remnv/go-boiler/internal/model"
	"github.com/remnv/go-boiler/internal/web/response"
)

const (
	JWT_DATA        = "JWT_DATA"
	E_DECODE_CLAIMS = "ERR_JWT_FALSE"
)

type JWTData struct {
	model.User
}

func TokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearer := getBearerAuth(c.Request)
		if bearer == nil {
			c.Abort()
			response.Error(c, response.ErrUnauthorized, "")
			return
		}
		claim, err := decodeJwtData(*bearer)
		if err != nil || (claim.ID == "" && len(claim.Scope) == 0) {
			c.Abort()
			response.Error(c, response.ErrUnauthorized, "")
			return
		}
		c.Set(JWT_DATA, claim)
		c.Next()
	}
}

func GetJWTData(c *gin.Context) (model.User, error) {
	claim, err := c.Get(JWT_DATA)
	if !err {
		return model.User{}, errors.New(response.E_NOT_FOUND)
	}
	return claim.(model.User), nil
}

func getBearerAuth(r *http.Request) *string {
	authHeader := r.Header.Get("Authorization")
	authForm := r.Form.Get("code")
	if authHeader == "" && authForm == "" {
		return nil
	}
	token := authForm
	if authHeader != "" {
		s := strings.SplitN(authHeader, " ", 2)
		if (len(s) != 2 || strings.ToLower(s[0]) != "bearer") && token == "" {
			return nil
		}
		//Use authorization header token only if token type is bearer else query string access token would be returned
		if len(s) > 0 && strings.ToLower(s[0]) == "bearer" {
			token = s[1]
		}
	}
	return &token
}

func decodeJwtData(token string) (model.User, error) {
	var claim JWTData
	tokenSplit := strings.Split(token, ".")
	if len(tokenSplit) < 2 {
		return model.User{}, errors.New(E_DECODE_CLAIMS)
	}
	token = tokenSplit[1]
	tokenDec, err := base64.RawStdEncoding.DecodeString(token)
	if err != nil {
		return model.User{}, err
	}
	err = json.Unmarshal(tokenDec, &claim)
	if err != nil {
		return model.User{}, err
	}
	return claim.User, nil
}
