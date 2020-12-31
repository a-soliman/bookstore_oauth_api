package db

import (
	"github.com/a-soliman/bookstore_oauth_api/src/clients/cassandra"
	"github.com/a-soliman/bookstore_oauth_api/src/domain/access_token"
	"github.com/a-soliman/bookstore_oauth_api/src/utils/errors"
)

// Repository the dbRepo interface
type Repository interface {
	GetByID(string) (*access_token.AccessToken, *errors.RestErr)
}

type repository struct {
}

// New returns a new Repository interface
func New() Repository {
	return &repository{}
}

func (r *repository) GetByID(accessTokenID string) (*access_token.AccessToken, *errors.RestErr) {
	session, err := cassandra.GetSession()
	if err != nil {
		panic(err)
	}
	defer session.Close()
	return nil, errors.NewInternalServerError("database connection not implemented")
}
