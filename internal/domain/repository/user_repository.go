package repository

import "go_ddd_example/internal/domain/entity"

type UserRepository interface {
	Create(user *entity.User) error
	Update(user *entity.User) error
	Delete(userID int64) error
	GetByID(userID int64) (*entity.User, error)
	GetByUsername(username string) (*entity.User, error)
}
