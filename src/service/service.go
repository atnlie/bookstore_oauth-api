package service

import (
	"amiera/src/domain/access_token"
	"amiera/src/utils/utils_errors"
)

type Repository interface {
	GetById(int64) (*access_token.AccessToken, *utils_errors.RestErr)
	GetAT() ([]access_token.AccessToken, *utils_errors.RestErr)
}

type Service interface {
	GetById(int64) (*access_token.AccessToken, *utils_errors.RestErr)
	GetAT() ([]access_token.AccessToken, *utils_errors.RestErr)
}

type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) GetById(atId int64) (*access_token.AccessToken, *utils_errors.RestErr)  {
	if atId < 0 {
		return nil, utils_errors.CustomBadRequestError("invalid access token id")
	}

	accessToken, err := s.repository.GetById(atId)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}

func (s *service) GetAT() ([]access_token.AccessToken, *utils_errors.RestErr)  {
	accessToken, err := s.repository.GetAT()
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}