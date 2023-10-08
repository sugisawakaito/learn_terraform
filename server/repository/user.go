package repository

import (
	"context"
	"server/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(ctx context.Context, tx *gorm.DB, user *model.User) error
}
