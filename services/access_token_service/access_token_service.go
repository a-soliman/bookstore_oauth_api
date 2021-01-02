package access_token_service

import (
	"strings"

	"github.com/a-soliman/bookstore_oauth_api/domain/access_token"
	"github.com/a-soliman/bookstore_oauth_api/repository/db"
	"github.com/a-soliman/bookstore_oauth_api/repository/rest"
	"github.com/a-soliman/bookstore_oauth_api/utils/errors"
)

// Service the service interface
type Service interface {
	GetByID(string) (*access_token.AccessToken, *errors.RestErr)
	Create(access_token.AccessTokenRequest) (*access_token.AccessToken, *errors.RestErr)
	UpdateExpirationTime(access_token.AccessToken) *errors.RestErr
}

type service struct {
	restUserRepo rest.UsersRepository
	dbRepo       db.Repository
}

// NewService returns a new service interface
func NewService(usersRepo rest.UsersRepository, dbRepo db.Repository) Service {
	return &service{
		restUserRepo: usersRepo,
		dbRepo:       dbRepo,
	}
}

func (s *service) GetByID(accessTokenID string) (*access_token.AccessToken, *errors.RestErr) {
	accessTokenID = strings.TrimSpace(accessTokenID)
	if accessTokenID == "" {
		return nil, errors.NewBadRequestError("invalid access token id")
	}
	accessToken, err := s.dbRepo.GetByID(accessTokenID)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}

func (s *service) Create(request access_token.AccessTokenRequest) (*access_token.AccessToken, *errors.RestErr) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	// TODO: support both grant types: client_credentials and password

	// Authenticate the user against the Users API:
	user, err := s.restUserRepo.LoginUser(request.Username, request.Password)
	if err != nil {
		return nil, err
	}

	// Generate a new access token:
	at := access_token.GetNewAccessToken(user.ID)
	at.Generate()

	// Save the new access token in Cassandra
	if err := s.dbRepo.Create(at); err != nil {
		return nil, err
	}
	return &at, nil
}

func (s *service) UpdateExpirationTime(at access_token.AccessToken) *errors.RestErr {
	if err := at.Validate(); err != nil {
		return err
	}
	return s.dbRepo.UpdateExpirationTime(at)
}
