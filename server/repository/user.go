package repository

import (
	"server/model"
)

type UserRepository interface {
	CreateUser() (*model.User, error)
}
