package service

import (
	"amiera/src/domain/access_token"
	"amiera/src/utils/utils_errors"
	"strings"
)

type Repository interface {
	GetById(string) (*access_token.AccessToken, *utils_errors.RestErr)
}

type Service interface {
	GetById(string) (*access_token.AccessToken, *utils_errors.RestErr)
}

type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) GetById(atId string) (*access_token.AccessToken, *utils_errors.RestErr)  {
	accessTokenId := strings.TrimSpace(atId)
	if len(accessTokenId) == 0 {
		return nil, utils_errors.CustomBadRequestError("invalid access token id")
	}

	accessToken, err := s.repository.GetById(atId)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}