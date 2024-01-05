package usecase

import (
	"log"

	"github.com/laut0104/RandomCooking/domain/entity"
	"github.com/laut0104/RandomCooking/domain/repository"
	"github.com/line/line-bot-sdk-go/linebot"
)

type LineUseCase struct {
	userRepo repository.User
}

func NewLineUseCase(userRepo repository.User) *LineUseCase {
	return &LineUseCase{userRepo: userRepo}
}

func (u *LineUseCase) Follow(lineUserID string, lineBot *linebot.Client) error {
	log.Println("test==========")
	profile, err := lineBot.GetProfile(lineUserID).Do()
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println(profile.DisplayName)

	user := &entity.User{
		LineUserID: lineUserID,
		UserName:   profile.DisplayName,
	}
	if err = u.userRepo.Store(user); err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (u *LineUseCase) UnFollow(lineUserID string) error {
	user, err := u.userRepo.FindByLineUserID(lineUserID)
	if err != nil {
		log.Println(err)
		return err
	}
	if err := u.userRepo.Delete(user.ID); err != nil {
		log.Println(err)
		return err
	}
	return nil
}
