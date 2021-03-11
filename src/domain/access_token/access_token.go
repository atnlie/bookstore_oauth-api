package access_token

import (
	"fmt"
	"time"
)

const (
	ExpirationTime = 24
)

type AccessToken struct {
	AccessToken string `json:"access_token" bson:"access_token,omitempty"`
	UserId      int64  `json:"user_id" bson:"user_id,omitempty"`
	ClientId    int64  `json:"client_id" bson:"client_id,omitempty"`
	Expires     int64  `json:"expires" bson:"expires,omitempty"`
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
