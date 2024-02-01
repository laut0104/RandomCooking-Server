package repository

import "github.com/laut0104/RandomCooking/domain/entity"

type Menu interface {
	FindByID(id string) (*entity.Menu, error)
	FindAllByUserID(userID string) ([]*entity.Menu, error)
	FindAllNotByUserID(userID string) ([]*entity.Menu, error)
	FindAll() ([]*entity.Menu, error)
	Store(menu *entity.Menu) error
	Update(menu *entity.Menu) error
	Delete(id, userID string) error
}
