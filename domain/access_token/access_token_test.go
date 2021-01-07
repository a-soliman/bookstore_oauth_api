package access_token

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAccessTokenConstants(t *testing.T) {
	assert.EqualValues(t, 24, expirationTime, "expiration time should be 25 hours")
}

// func TestGetNewAccessToken(t *testing.T) {
// 	at := GetNewAccessToken()

// 	assert.False(t, at.IsExpired(), "brand new access token should not be expired")
// 	assert.EqualValues(t, "", at.AccessToken, "new access token should not have defined access token id")
// 	assert.EqualValues(t, 0, at.UserID, "new access token should not have an associated user id")
// }

func TestAccessTokenIsExpired(t *testing.T) {
	at := AccessToken{}
	assert.True(t, at.IsExpired(), "empty access token should be expired by default")

	at.Expires = time.Now().UTC().Add(3 * time.Hour).Unix()
	assert.False(t, at.IsExpired(), "access token expired three hours from now should not be expired")
}
