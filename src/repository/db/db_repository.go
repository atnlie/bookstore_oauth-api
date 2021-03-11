package db

import (
	"amiera/src/clients/mongodb"
	"amiera/src/domain/access_token"
	"amiera/src/utils/utils_errors"
)

type DbRepository interface {
	GetAT() ([]access_token.AccessToken, *utils_errors.RestErr)
	GetById(int64) (*access_token.AccessToken, *utils_errors.RestErr)
	GetOptionById(int64) ([]access_token.AccessToken, *utils_errors.RestErr)
}


type dbRepository struct {
}

func NewRepository() DbRepository {
	return &dbRepository{}
}

func (r *dbRepository) GetAT() ([]access_token.AccessToken, *utils_errors.RestErr) {
	at, err := mongodb.GetAllAccessToken()
	if err != nil {
		return nil, utils_errors.CustomInternalServerError("Data not found")
	}

	return at, nil
}

func (r *dbRepository) GetById(acId int64) (*access_token.AccessToken, *utils_errors.RestErr) {
	at, err := mongodb.GetAccessTokenById(acId)
	if err != nil {
		return nil, err
	}
	return at, nil
}

func (r *dbRepository) GetOptionById(filter int64) ([]access_token.AccessToken, *utils_errors.RestErr)  {
	at, err := mongodb.GetOptionAccessTokenById(filter)
	if err != nil {
		return nil, err
	}
	return at, nil
}