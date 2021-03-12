package service

import (
	"amiera/src/domain/access_token"
	"amiera/src/repository/db"
	"amiera/src/utils/utils_errors"
	"fmt"
)

type Repository interface {
	GetById(int64) (*access_token.AccessToken, *utils_errors.RestErr)
	GetAT() ([]access_token.AccessToken, *utils_errors.RestErr)
	GetOptionById(int64) ([]access_token.AccessToken, *utils_errors.RestErr)
	CreateToken(access_token.AccessToken) *utils_errors.RestErr
	UpdateExpiration(access_token.AccessToken) *utils_errors.RestErr
}

type Service interface {
	GetById(int64) (*access_token.AccessToken, *utils_errors.RestErr)
	GetAT() ([]access_token.AccessToken, *utils_errors.RestErr)
	GetOptionById(int64) ([]access_token.AccessToken, *utils_errors.RestErr)
	CreateToken(access_token.AccessToken) *utils_errors.RestErr
	UpdateExpiration(access_token.AccessToken, bool) *utils_errors.RestErr
}

type service struct {
	repository Repository
}

func NewService(repo db.DbRepository) Service {
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

func (s *service) GetOptionById(filter int64) ([]access_token.AccessToken, *utils_errors.RestErr)  {
	if filter < 0 {
		return nil, utils_errors.CustomBadRequestError("invalid access token id")
	}

	accessToken, err := s.repository.GetOptionById(filter)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}

func (s *service) CreateToken(token access_token.AccessToken) *utils_errors.RestErr  {
	err := s.repository.CreateToken(token)
	if err != nil {
		return utils_errors.CustomBadRequestError("token failed to created")
	}
	return nil
}

func (s *service) UpdateExpiration(token access_token.AccessToken, isPartial bool) *utils_errors.RestErr {
	currentToken, err := s.repository.GetById(token.UserId)
	if err != nil {
		return err
	}

	if err := token.Validation(); err != nil {
		return err
	}

	if isPartial {
		if token.UserId > 0 {
			currentToken.UserId = token.UserId
		}

		if token.Expires > 0 {
			currentToken.Expires = token.Expires
		}

		if token.ClientId > 0 {
			currentToken.ClientId = token.ClientId
		}
		if token.AccessToken != "" {
			currentToken.AccessToken = token.AccessToken
		}
	} else {
		currentToken.UserId = token.UserId
		currentToken.Expires = token.Expires
		currentToken.ClientId = token.ClientId
		currentToken.AccessToken = token.AccessToken
	}

	fmt.Println("currentToken", currentToken)

	if err = s.repository.UpdateExpiration(token); err != nil {
		return utils_errors.CustomBadRequestError("can't update token expiration")
	}

	return nil
}