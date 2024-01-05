package usecase

import (
	"log"

	"github.com/laut0104/RandomCooking/domain/entity"
	"github.com/laut0104/RandomCooking/domain/repository"
)

type UserUseCase struct {
	userRepo repository.User
}

// NewUserUseCase はUserUseCaseのポインタを生成します。
func NewUserUseCase(userRepo repository.User) *UserUseCase {
	return &UserUseCase{userRepo: userRepo}
}

func (u *UserUseCase) GetUserByID(id string) (*entity.User, error) {
	user, err := u.userRepo.FindByID(id)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return user, nil
}

func (u *UserUseCase) GetUserByLineUserID(lineUserID string) (*entity.User, error) {
	user, err := u.userRepo.FindByLineUserID(lineUserID)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return user, nil
}
