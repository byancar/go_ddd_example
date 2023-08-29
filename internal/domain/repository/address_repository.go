package repository

import "go_ddd_example/internal/domain/entity"

type AddressRepository interface {
	Create(address *entity.Address) error
	Update(address *entity.Address) error
	Delete(addressID int64) error
	GetByID(addressID int64) (*entity.Address, error)
}
