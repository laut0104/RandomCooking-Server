package repository

import "github.com/laut0104/RandomCooking/domain/entity"

// User はユーザの永続化を担当するリポジトリです。
type User interface {
	FindByID(id string) (*entity.User, error)
	FindByLineUserID(lineUserID string) (*entity.User, error)
	Store(user *entity.User) error
	Update(user *entity.User) error
	Delete(id string) error
}
