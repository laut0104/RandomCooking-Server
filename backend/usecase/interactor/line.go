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

func (u *LineUseCase) CreateRichMenu(richMenu linebot.RichMenu, lineBot *linebot.Client) (string, error) {
	res, err := lineBot.CreateRichMenu(richMenu).Do()
	if err != nil {
		log.Println(err)
		return "", err
	}
	return res.RichMenuID, nil
}

func (u *LineUseCase) SetRichMenuImage(lineBot *linebot.Client, richMenuId string, filePath string) error {
	if _, err := lineBot.UploadRichMenuImage(richMenuId, filePath).Do(); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (u *LineUseCase) SetDefaultRichMenu(lineBot *linebot.Client, richMenuId string) error {
	if _, err := lineBot.SetDefaultRichMenu(richMenuId).Do(); err != nil {
		log.Println(err)
		return err
	}
	return nil
}
