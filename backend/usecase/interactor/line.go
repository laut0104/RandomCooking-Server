package usecase

import (
	"log"

	"github.com/laut0104/RandomCooking/domain/entity"
	"github.com/laut0104/RandomCooking/domain/repository"
	"github.com/line/line-bot-sdk-go/linebot"
)

type LineUseCase struct {
	userRepo repository.User
	lineBot  *linebot.Client
}

func NewLineUseCase(userRepo repository.User, lineBot *linebot.Client) *LineUseCase {
	return &LineUseCase{
		userRepo: userRepo,
		lineBot:  lineBot,
	}
}

func (u *LineUseCase) Follow(lineUserID string) error {
	log.Println("test==========")
	profile, err := u.lineBot.GetProfile(lineUserID).Do()
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

func (u *LineUseCase) CreateRichMenu(richMenu linebot.RichMenu) (string, error) {
	res, err := u.lineBot.CreateRichMenu(richMenu).Do()
	if err != nil {
		log.Println(err)
		return "", err
	}
	return res.RichMenuID, nil
}

func (u *LineUseCase) SetRichMenuImage(richMenuId string, filePath string) error {
	if _, err := u.lineBot.UploadRichMenuImage(richMenuId, filePath).Do(); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (u *LineUseCase) SetDefaultRichMenu(richMenuId string) error {
	if _, err := u.lineBot.SetDefaultRichMenu(richMenuId).Do(); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (u *LineUseCase) SendFlexMessage(lineUserID string, flexMessage *linebot.FlexMessage) error {
	if _, err := u.lineBot.PushMessage(
		lineUserID,
		flexMessage,
	).Do(); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (u *LineUseCase) SendMessage(lineUserID string, message *linebot.TextMessage) error {
	if _, err := u.lineBot.PushMessage(
		lineUserID,
		message,
	).Do(); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (u *LineUseCase) ReplyFlexMessage(replyToken string, flexMessage *linebot.FlexMessage) error {
	if _, err := u.lineBot.ReplyMessage(
		replyToken,
		flexMessage,
	).Do(); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (u *LineUseCase) ReplyMessage(replyToken string, message *linebot.TextMessage) error {
	if _, err := u.lineBot.ReplyMessage(
		replyToken,
		message,
	).Do(); err != nil {
		log.Println(err)
		return err
	}
	return nil
}
