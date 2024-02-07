package usecase

import (
	"log"

	"github.com/laut0104/RandomCooking/domain/entity"
	"github.com/laut0104/RandomCooking/domain/repository"
)

type LikeUseCase struct {
	likeRepo repository.Like
}

func NewLikeUseCase(likeRepo repository.Like) *LikeUseCase {
	return &LikeUseCase{likeRepo: likeRepo}
}

type LikesMenu struct {
	ID     string `json:"id" param:"id"`
	UserID string `json:"userid" param:"uid"`
	MenuID string `json:"menuid" param:"menuid"`
}

func (u *LikeUseCase) GetLikesMenuByUserID(userID string) ([]*repository.LikesMenu, error) {
	likesMenu, err := u.likeRepo.FindLikesMenuByUserID(userID)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return likesMenu, nil
}

func (u *LikeUseCase) GetLikeByUniqueKey(userID, menuID string) (*entity.Like, error) {
	like, err := u.likeRepo.FindByUniqueKey(userID, menuID)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return like, nil
}

func (u *LikeUseCase) AddLike(like *entity.Like) error {
	if err := u.likeRepo.Store(like); err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (u *LikeUseCase) DeleteLike(id string) error {
	if err := u.likeRepo.Delete(id); err != nil {
		log.Println(err)
		return err
	}

	return nil
}
