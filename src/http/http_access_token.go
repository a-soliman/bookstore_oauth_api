package http

import (
	"net/http"

	"github.com/a-soliman/bookstore_oauth_api/src/domain/access_token"
	"github.com/gin-gonic/gin"
)

// AccessTokenHandler the handler interface
type AccessTokenHandler interface {
	GetByID(*gin.Context)
}

type accessTokenHandler struct {
	service access_token.Service
}

// NewHandler returns a new AccessTokenHandler
func NewHandler(service access_token.Service) AccessTokenHandler {
	return &accessTokenHandler{
		service: service,
	}
}

func (handler *accessTokenHandler) GetByID(c *gin.Context) {
	accessToken, err := handler.service.GetByID(c.Param("access_token_id"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, accessToken)
}
