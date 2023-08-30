package users

import (
	"go_ddd_example/internal/domain/entity"
	"go_ddd_example/internal/domain/repository"
)

type UserRepositoryImpl struct {
	querier Querier
}

func NewUserRepository(querier Querier) repository.UserRepository {
	return &UserRepositoryImpl{
		querier: querier,
	}
}

func (repo *UserRepositoryImpl) Create(user *entity.User) error {
	return nil
}

func (repo *UserRepositoryImpl) Update(user *entity.User) error {
	return nil
}

func (repo *UserRepositoryImpl) Delete(userID int64) error {
	return nil
}

func (repo *UserRepositoryImpl) GetByID(userID int64) (*entity.User, error) {
	return nil, nil
}

func (repo *UserRepositoryImpl) GetByUsername(username string) (*entity.User, error) {
	return nil, nil
}
