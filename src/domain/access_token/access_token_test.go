package access_token

import (
	"github.com/stretchr/testify/assert"

	"testing"
	"time"
)

func TestAccessTokenConstans(t *testing.T) {
	assert.EqualValues(t, 24, ExpirationTime, "Expiration time should be 24")
}

func TestGetNewAccessToken(t *testing.T) {
	at := GetNewAccessToken()
	assert.False(t, at.isExpired(), "branch access token should not be expired")
	assert.EqualValues(t, "", at.AccessToken, "new access token not have to defined")
	assert.True(t, at.UserId == 0, "new access token should have an associated user id")
}

func TestAccessTokenIsExpired(t *testing.T) {
	at := AccessToken{}
	assert.True(t, at.isExpired(), "empty access token by default")

	at.Expires = time.Now().UTC().Add(3 * time.Hour).Unix()
	assert.False(t, at.isExpired(),"access token expiring 3 hours" )
}