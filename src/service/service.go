package service

import (
	"amiera/src/domain/access_token"
	"amiera/src/utils/utils_errors"
)

type Service interface {
	GetById(string) (*access_token.AccessToken, *utils_errors.RestErr)
}