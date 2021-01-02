package http

import (
	"net/http"

	atDomain "github.com/a-soliman/bookstore_oauth_api/src/domain/access_token"
	"github.com/a-soliman/bookstore_oauth_api/src/services/access_token_service"
	"github.com/a-soliman/bookstore_oauth_api/src/utils/errors"
	"github.com/gin-gonic/gin"
)

// AccessTokenHandler the handler interface
type AccessTokenHandler interface {
	GetByID(*gin.Context)
	Create(*gin.Context)
}

type accessTokenHandler struct {
	service access_token_service.Service
}

// NewHandler returns a new AccessTokenHandler
func NewHandler(service access_token_service.Service) AccessTokenHandler {
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

func (handler *accessTokenHandler) Create(c *gin.Context) {
	var request atDomain.AccessTokenRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	accessToken, err := handler.service.Create(request)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusCreated, accessToken)
}
