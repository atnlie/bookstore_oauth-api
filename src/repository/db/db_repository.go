package db

import (
	"amiera/src/domain/access_token"
	"amiera/src/utils/utils_errors"
)

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *utils_errors.RestErr)
}

type dbRepository struct {
}

func (r *dbRepository) GetById(acId string) (*access_token.AccessToken, *utils_errors.RestErr) {
	return nil, utils_errors.CustomInternalServerError("Db connection not implemented yet")
}

func NewRepository() DbRepository {
	return &dbRepository{}
}
