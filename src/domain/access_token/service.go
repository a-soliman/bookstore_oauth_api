package access_token

import (
	"strings"

	"github.com/a-soliman/bookstore_oauth_api/src/utils/errors"
)

// Repository the Repository interface
type Repository interface {
	GetByID(string) (*AccessToken, *errors.RestErr)
}

// Service the service interface
type Service interface {
	GetByID(string) (*AccessToken, *errors.RestErr)
}

type service struct {
	repository Repository
}

// NewService Creates a new service instance
func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

// GetByID gets access token by id
func (s *service) GetByID(accessTokenID string) (*AccessToken, *errors.RestErr) {
	accessTokenID = strings.TrimSpace(accessTokenID)
	if len(accessTokenID) == 0 {
		return nil, errors.NewBadRequestError("invalid access token id")
	}
	return s.repository.GetByID(accessTokenID)
}
