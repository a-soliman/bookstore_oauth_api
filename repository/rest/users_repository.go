package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/federicoleon/golang-restclient/rest"
	"github.com/joho/godotenv"

	"github.com/a-soliman/bookstore_oauth_api/domain/users"
	"github.com/a-soliman/bookstore_utils-go/rest_errors"
)

const (
	usersMicroserviceURLVar = "USERS_MICROSERVICE_URL"
)

var (
	usersMicroserviceURL = goDotEnvVariable(usersMicroserviceURLVar)
	usersRestClient      = rest.RequestBuilder{
		BaseURL: usersMicroserviceURL,
		Timeout: 2 * time.Second,
	}
)

// UsersRepository interface
type UsersRepository interface {
	LoginUser(string, string) (*users.User, rest_errors.RestErr)
}

type usersRepository struct{}

// New returns a new RestUsersRepository
func New() UsersRepository {
	return &usersRepository{}
}

func (r *usersRepository) LoginUser(email, password string) (*users.User, rest_errors.RestErr) {
	request := users.UserLoginRequest{
		Email:    email,
		Password: password,
	}
	response := usersRestClient.Post("/users/login", request)
	if response == nil || response.Response == nil {
		return nil, rest_errors.NewInternalServerError("invalid restclient response while trying to login user", nil)
	}
	if response.StatusCode > 299 {
		restErr, err := rest_errors.NewRestErrorFromBytes(response.Bytes())
		if err != nil {
			fmt.Println(err)
			return nil, rest_errors.NewInternalServerError("invalid error interface while trying to login user", err)
		}
		return nil, restErr
	}
	var user users.User
	if err := json.Unmarshal(response.Bytes(), &user); err != nil {
		return nil, rest_errors.NewInternalServerError("error while trying to unmarshal users response to login user", err)
	}
	return &user, nil
}

// TODO: this is duplicated and must go to utils instead.
func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
