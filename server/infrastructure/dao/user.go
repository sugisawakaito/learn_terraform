package dao

import (
	"context"
	"server/model"

	// "github.com/bluele/go-timecop"
	"golang.org/x/xerrors"
	"gorm.io/gorm"
)

type (
	userDAO interface {
		CreateUser(ctx context.Context, tx *gorm.DB, user *model.User) error
	}
	UserDAO struct{}
)

func NewUser() userDAO {
	return &UserDAO{}
}

func (d *UserDAO) CreateUser(ctx context.Context, tx *gorm.DB, user *model.User) error {
	if err := tx.Create(user).Error; err != nil {
		return xerrors.Errorf("failed to create user: %w", err)
	}
	return nil
}
