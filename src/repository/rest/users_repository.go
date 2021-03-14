package rest

import (
	"amiera/src/domain/users"
	"amiera/src/utils/utils_errors"
	"encoding/json"
	"github.com/mercadolibre/golang-restclient/rest"
	"time"
)

var (
	usersRestClient = rest.RequestBuilder{
		//BaseURL: "https://api.bookstore.com",
		BaseURL: "http://localhost:8080",
		Timeout: 100 * time.Microsecond,
	}
)

type RestUsersRepository interface {
	LoginUser(string, string) (*users.User, *utils_errors.RestErr)
}

type usersRepository struct {
}

func NewRepository() RestUsersRepository {
	return &usersRepository{}
}

func (s *usersRepository) LoginUser(email string, pass string) (*users.User, *utils_errors.RestErr) {
	reqBody := users.UserLoginRequest{
		Email:    email,
		Password: pass,
	}
	response := usersRestClient.Post("/users/login", reqBody)
	if response == nil || response.Response == nil {
		return nil, utils_errors.CustomInternalServerError("invalid rest-client response when trying login")
	}

	if response.StatusCode > 299 {
		var restErr utils_errors.RestErr
		err := json.Unmarshal(response.Bytes(), &restErr)
		if err != nil {
			return nil, utils_errors.CustomInternalServerError("invalid interface error when trying to login user")
		}
		return nil, &restErr
	}

	var user users.User
	if err := json.Unmarshal(response.Bytes(), &user); err != nil {
		return nil, utils_errors.CustomInternalServerError("error when trying unmarshal users response")
	}

	return &user, nil
}
