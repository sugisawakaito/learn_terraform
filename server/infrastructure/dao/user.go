package dao

import (
	"server/model"
)

type (
	userDAO interface {
		CreateUser() (*model.User, error)
	}
	UserDAO struct{}
)

func NewUser() userDAO {
	return &UserDAO{}
}

func (d *UserDAO) CreateUser() (*model.User, error) {
	return nil, nil
}
