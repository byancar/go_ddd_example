package user

import (
	"go_ddd_example/internal/domain/entity"
	"go_ddd_example/internal/domain/repository"
)

type UserService interface {
	Create() (*entity.User, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (s *userService) Create() (*entity.User, error) {
	return nil, nil
}
