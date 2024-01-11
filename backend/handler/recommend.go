package handler

import (
	usecase "github.com/laut0104/RandomCooking/usecase/interactor"
	_ "github.com/lib/pq"
	"github.com/line/line-bot-sdk-go/linebot"

	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type RecommendHandler struct {
	userUC      *usecase.UserUseCase
	recommendUC *usecase.RecommendUseCase
	lineUC      *usecase.LineUseCase
}

func NewRecommendHandler(userUC *usecase.UserUseCase, recommendUC *usecase.RecommendUseCase, lineUC *usecase.LineUseCase) *RecommendHandler {
	return &RecommendHandler{
		userUC:      userUC,
		recommendUC: recommendUC,
		lineUC:      lineUC,
	}
}

func (h *RecommendHandler) RecommendMenu(c echo.Context) error {
	uid := c.Param("uid")

	var recommendedList []string
	flexMessage, err := h.recommendUC.RecommendMenu(uid, recommendedList)
	if err != nil {
		log.Println(err)
		return err
	}

	user, err := h.userUC.GetUserByID(uid)
	if err != nil {
		log.Println(err)
		return err
	}

	if flexMessage == nil {
		errMsg := "おすすめできるメニューがありません。メニューを登録してください。"
		if err = h.lineUC.SendMessage(user.LineUserID, linebot.NewTextMessage(errMsg)); err != nil {
			log.Print(err)
			return err
		}
		return nil
	}

	if err = h.lineUC.SendFlexMessage(user.LineUserID, flexMessage); err != nil {
		log.Println(err)
		return err
	}

	return c.NoContent(http.StatusOK)
}
