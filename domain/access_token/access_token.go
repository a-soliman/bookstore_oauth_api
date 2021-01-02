package access_token

import (
	"fmt"
	"strings"
	"time"

	"github.com/a-soliman/bookstore_utils-go/crypto_utils"
	"github.com/a-soliman/bookstore_utils-go/rest_errors"
)

const (
	expirationTime = 24
)

// AccessToken struct
type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserID      int64  `json:"user_id"`
	ClientID    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
}

// GetNewAccessToken returns a new access token
func GetNewAccessToken(userID int64) AccessToken {
	return AccessToken{
		UserID:  userID,
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

// IsExpired checks if the access token is expired
func (at *AccessToken) IsExpired() bool {
	now := time.Now().UTC()
	expirationTime := time.Unix(at.Expires, 0)
	return now.After(expirationTime)
}

// Generate generates the access token
func (at *AccessToken) Generate() {
	at.AccessToken = crypto_utils.GetMd5(fmt.Sprintf("at-%d-%d-ran", at.UserID, at.Expires))
}

// Validate validates the access token
func (at *AccessToken) Validate() *rest_errors.RestErr {
	at.AccessToken = strings.TrimSpace(at.AccessToken)
	if at.AccessToken == "" {
		return rest_errors.NewBadRequestError("invalid access token id")
	}
	if at.UserID <= 0 {
		return rest_errors.NewBadRequestError("invalid user id")
	}
	if at.ClientID <= 0 {
		return rest_errors.NewBadRequestError("invalid client id")
	}
	if at.Expires <= 0 {
		return rest_errors.NewBadRequestError("invalid expiration time")
	}
	return nil
}
