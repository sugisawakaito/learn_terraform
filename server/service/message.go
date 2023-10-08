package service

import (
	"context"
	"server/infrastructure/dao"
	"server/model"
	"server/repository"

	"github.com/bluele/go-timecop"
	"golang.org/x/xerrors"
	"gorm.io/gorm"
)

type (
	MessengerService interface {
		CreateUser(ctx context.Context, tx *gorm.DB, msg string) error
	}

	messengerService struct {
		userRepository repository.UserRepository
	}
)

func NewMessengerService() MessengerService {
	return &messengerService{
		userRepository: dao.NewUser(),
	}
}

func (s *messengerService) CreateUser(ctx context.Context, tx *gorm.DB, msg string) error {
	now := timecop.Now()
	user := &model.User{
		Email:     msg,
		FullName:  msg,
		CreatedAt: now,
	}
	if err := s.userRepository.CreateUser(ctx, tx, user); err != nil {
		return xerrors.Errorf("failed to create user: %w", err)
	}
	return nil
}
