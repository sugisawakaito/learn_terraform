package service

import (
	"server/model"
	"server/repository"
	"server/infrastructure/dao"
)

type (
	messengerService interface {
		CreateUser() (*model.User, error)
	}

	MessengerService struct{
		userRepository repository.UserRepository
	}
)

func NewMessengerService() MessengerService {
	return &messengerService{
		userRepository: dao.NewUser(),
	}
}

func (s *MessengerService) CreateUser() (*model.User, error) {
	return nil, nil
}
