package access_token

import (
	"fmt"
	"time"
)

const (
	ExpirationTime = 24
)

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserId      int64  `json:"user_id"`
	ClientId    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
}

func GetNewAccessToken() AccessToken {
	return AccessToken{
		Expires: time.Now().UTC().Add(ExpirationTime * time.Hour).Unix(),
	}
}

func (at AccessToken) isExpired() bool {
	now := time.Now().UTC()
	expirationTime := time.Unix(at.Expires, 0)
	fmt.Println("ExpirationTime: ", expirationTime)

	return expirationTime.Before(now)
}