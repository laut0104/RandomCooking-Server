package repository

import "github.com/laut0104/RandomCooking/domain/entity"

type Like interface {
	// FindByID(id, userID string) (*entity.Like, error)
	FindAllByUserID(userID string) ([]*entity.Like, error)
	// FindAll() ([]*entity.Like, error)
	// Store(menu *entity.Like) error
	// Update(menu *entity.Like) error
	// Delete(id, userID string) error
}
