package access_token

import (
	"github.com/a-soliman/bookstore_oauth_api/utils/errors"
)

const (
	grantTypePassword          = "password"
	grantTypeClientCredentials = "client_credentials"
)

// AccessTokenRequest struct
type AccessTokenRequest struct {
	GrantType string `json:"grant_type"`
	Scope     string `json:"scope"`

	// Used for password grant_type
	Username string `json:"user_name"`
	Password string `json:"password"`

	// Used for client_credentials grant_type
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

// Validate validates the access token
func (atr *AccessTokenRequest) Validate() *errors.RestErr {
	switch atr.GrantType {
	case grantTypePassword:
		break

	case grantTypeClientCredentials:
		break

	default:
		return errors.NewBadRequestError("invalid grant_type paramater")
	}

	// Todo: validate parameters for each grant_type
	return nil
}
