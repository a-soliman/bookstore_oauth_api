package db

import (
	"github.com/a-soliman/bookstore_oauth_api/clients/cassandra"
	"github.com/a-soliman/bookstore_oauth_api/domain/access_token"
	"github.com/a-soliman/bookstore_utils-go/rest_errors"
	"github.com/gocql/gocql"
)

const (
	queryGetAccessToken       = "SELECT access_token, user_id, client_id, expires FROM access_tokens WHERE access_token=?;"
	queryCreateAccessToken    = "INSERT INTO access_tokens(access_token, user_id, client_id, expires) VALUES(?, ?, ?, ?);"
	queryUpdateExpirationTime = "UPDATE access_tokens SET expires=? WHERE access_token=?;"
)

// Repository the dbRepo interface
type Repository interface {
	GetByID(string) (*access_token.AccessToken, *rest_errors.RestErr)
	Create(access_token.AccessToken) *rest_errors.RestErr
	UpdateExpirationTime(access_token.AccessToken) *rest_errors.RestErr
}

type repository struct {
}

// New returns a new Repository interface
func New() Repository {
	return &repository{}
}

func (r *repository) GetByID(accessTokenID string) (*access_token.AccessToken, *rest_errors.RestErr) {
	var result access_token.AccessToken

	if err := cassandra.GetSession().Query(queryGetAccessToken, accessTokenID).Scan(
		&result.AccessToken,
		&result.UserID,
		&result.ClientID,
		&result.Expires,
	); err != nil {
		if err == gocql.ErrNotFound {
			return nil, rest_errors.NewNotFoundError("no access token found with given id")
		}
		return nil, rest_errors.NewInternalServerError(err.Error(), err)
	}
	return &result, nil
}

func (r *repository) Create(token access_token.AccessToken) *rest_errors.RestErr {
	if err := cassandra.GetSession().Query(queryCreateAccessToken,
		token.AccessToken,
		token.UserID,
		token.ClientID,
		token.Expires,
	).Exec(); err != nil {
		return rest_errors.NewInternalServerError(err.Error(), err)
	}

	return nil
}

func (r *repository) UpdateExpirationTime(token access_token.AccessToken) *rest_errors.RestErr {
	if err := cassandra.GetSession().Query(queryUpdateExpirationTime,
		token.Expires,
		token.AccessToken,
	).Exec(); err != nil {
		return rest_errors.NewInternalServerError(err.Error(), err)
	}

	return nil
}
