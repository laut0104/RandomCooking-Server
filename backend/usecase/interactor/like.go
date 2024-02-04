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
