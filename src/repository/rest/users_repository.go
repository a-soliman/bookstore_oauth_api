package rest

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/federicoleon/golang-restclient/rest"

	"github.com/a-soliman/bookstore_oauth_api/src/domain/users"
	"github.com/a-soliman/bookstore_oauth_api/src/utils/errors"
)

var (
	usersRestClient = rest.RequestBuilder{
		BaseURL: "http://localhost:8082",
		Timeout: 100 * time.Millisecond,
	}
)

// UsersRepository interface
type UsersRepository interface {
	LoginUser(string, string) (*users.User, *errors.RestErr)
}

type usersRepository struct{}

// New returns a new RestUsersRepository
func New() UsersRepository {
	return &usersRepository{}
}

func (r *usersRepository) LoginUser(email, password string) (*users.User, *errors.RestErr) {
	request := users.UserLoginRequest{
		Email:    email,
		Password: password,
	}
	response := usersRestClient.Post("/users/login", request)
	if response == nil || response.Response == nil {
		fmt.Println(response)
		return nil, errors.NewInternalServerError("invalid restclient response while trying to login user")
	}
	if response.StatusCode > 299 {
		var restErr errors.RestErr
		if err := json.Unmarshal(response.Bytes(), &restErr); err != nil {
			fmt.Println("here")
			return nil, errors.NewInternalServerError("invalid error interface while trying to login user")
		}
		return nil, &restErr
	}
	var user users.User
	if err := json.Unmarshal(response.Bytes(), &user); err != nil {
		return nil, errors.NewInternalServerError("error while trying to unmarshal users response to login user")
	}
	return &user, nil
}
