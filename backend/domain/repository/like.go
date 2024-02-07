package repository

import "github.com/laut0104/RandomCooking/domain/entity"

type LikesMenu struct {
	ID          string   `json:"id" param:"id"`
	MenuID      string   `json:"menuid" param:"menuid"`
	UserID      string   `json:"userid" param:"uid"`
	MenuName    string   `json:"menuname" param:"menuname"`
	ImageUrl    string   `json:"imageurl" param:"imageurl"`
	Ingredients []string `json:"ingredients" param:"ingredients"`
	Quantities  []string `json:"quantities" param:"quantities"`
	Recipes     []string `json:"recipes" param:"recipes"`
}

type Like interface {
	// FindByID(id, userID string) (*entity.Like, error)
	FindByUniqueKey(userID, menuID string) (*entity.Like, error)
	FindAllByUserID(userID string) ([]*entity.Like, error)
	FindLikesMenuByUserID(userID string) ([]*LikesMenu, error)
	// FindAll() ([]*entity.Like, error)
	Store(menu *entity.Like) error
	// Update(menu *entity.Like) error
	Delete(id string) error
}
