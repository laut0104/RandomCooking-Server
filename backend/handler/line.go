package handler

import (
	"log"
	"os"
	"strings"

	usecase "github.com/laut0104/RandomCooking/usecase/interactor"
	_ "github.com/lib/pq"
	"github.com/line/line-bot-sdk-go/linebot"

	"net/http"

	"github.com/labstack/echo/v4"
)

type LineHandler struct {
	userUC      *usecase.UserUseCase
	lineUC      *usecase.LineUseCase
	recommendUC *usecase.RecommendUseCase
}

func NewLineHandler(userUC *usecase.UserUseCase, lineUC *usecase.LineUseCase, recommendUC *usecase.RecommendUseCase) *LineHandler {
	return &LineHandler{
		userUC:      userUC,
		lineUC:      lineUC,
		recommendUC: recommendUC,
	}
}

func (h *LineHandler) LineEvent(c echo.Context) error {
	bot, err := linebot.New(
		os.Getenv("LINE_BOT_CHANNEL_SECRET"),
		os.Getenv("LINE_BOT_CHANNEL_TOKEN"),
	)
	if err != nil {
		log.Println(err)
		return err
	}
	events, err := bot.ParseRequest(c.Request())
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			c.Response().WriteHeader(400)
			return c.String(400, "Hello, World!")
		} else {
			c.Response().WriteHeader(500)
			return c.String(500, "Hello, World!")
		}
	}
	for _, event := range events {
		switch event.Type {
		// Follow
		case linebot.EventTypeFollow:
			message := "友達登録ありがとう！"
			errMsg := "正常にユーザー登録できませんでした\nブロックし、もう一度友達登録をお願いします"
			if err := h.lineUC.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message)); err != nil {
				log.Println(err)
				if err = h.lineUC.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(errMsg)); err != nil {
					log.Println(err)
				}
				return err
			}
			if err := h.lineUC.Follow(event.Source.UserID); err != nil {
				log.Println(err)
				if err = h.lineUC.SendMessage(event.Source.UserID, linebot.NewTextMessage(errMsg)); err != nil {
					log.Println(err)
				}
				return err
			}

			richMenu := &linebot.RichMenu{
				Size:        linebot.RichMenuSize{Width: 2500, Height: 1686},
				Selected:    false,
				Name:        "richmenu-demo",
				ChatBarText: "Tap here",
				Areas: []linebot.AreaDetail{
					{
						Bounds: linebot.RichMenuBounds{X: 74, Y: 278, Width: 730, Height: 1130},
						Action: linebot.RichMenuAction{
							Type: "uri",
							URI:  "https://liff.line.me/1660690567-wegZZboy/menu-list",
						},
					},
					{
						Bounds: linebot.RichMenuBounds{X: 885, Y: 278, Width: 730, Height: 1130},
						Action: linebot.RichMenuAction{
							Type: "uri",
							URI:  "https://liff.line.me/1660690567-wegZZboy/like-list",
						},
					},
					{
						Bounds: linebot.RichMenuBounds{X: 1695, Y: 278, Width: 730, Height: 1130},
						Action: linebot.RichMenuAction{
							Type: "message",
							Text: "今日のメニュー何がいいかな？",
						},
					},
				},
			}

			richMenuID, err := h.lineUC.CreateRichMenu(*richMenu)
			if err != nil {
				log.Println(err)
				return err
			}
			if err = h.lineUC.SetRichMenuImage(richMenuID, os.Getenv("RICHMENU_IMG")); err != nil {
				log.Println(err)
				return err
			}
			if err = h.lineUC.SetDefaultRichMenu(richMenuID); err != nil {
				log.Println(err)
				return err
			}
			log.Println("Add user success =====")

		// Unfollow
		case linebot.EventTypeUnfollow:
			if err := h.lineUC.UnFollow(event.Source.UserID); err != nil {
				log.Println(err)
				return err
			}
			log.Println("Delete user success =====")

		case linebot.EventTypePostback:
			replyToken := event.ReplyToken
			userID := event.Source.UserID
			postBackData := event.Postback.Data
			postBackDataList := strings.Split(postBackData, ",")
			recommendedMenuList := postBackDataList[1:]

			user, err := h.userUC.GetUserByLineUserID(userID)
			if err != nil {
				log.Println(err)
				return err
			}

			var flexMessage *linebot.FlexMessage
			if postBackDataList[0] == "MyMenu" {
				flexMessage, err = h.recommendUC.RecommendMyMenu(user.ID, recommendedMenuList)
				if err != nil {
					log.Println(err)
					return err
				}
			} else if postBackDataList[0] == "MyMenuAndLikeMenu" {
				flexMessage, err = h.recommendUC.RecommendMyMenuAndLikeMenu(user.ID, recommendedMenuList)
				if err != nil {
					log.Println(err)
					return err
				}
			}

			if flexMessage == nil {
				errMsg := "おすすめできるメニューがありません。メニューを登録してください。"
				if err = h.lineUC.ReplyMessage(replyToken, linebot.NewTextMessage(errMsg)); err != nil {
					log.Print(err)
					return err
				}
				return nil
			}

			if err = h.lineUC.ReplyFlexMessage(replyToken, flexMessage); err != nil {
				log.Println(err)
				return err
			}
			// Text
		case linebot.EventTypeMessage:
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				replyToken := event.ReplyToken
				if message.Text == "今日のメニュー何がいいかな？" {
					json := `{
						"type": "bubble",
						"size": "giga",
						"body": {
							"type": "box",
							"layout": "vertical",
							"contents": [
								{
									"type": "text",
									"text": "今日のメニューは何から選ぶ？",
									"size": "md",
									"adjustMode": "shrink-to-fit",
									"align": "center"
								}
							]
						},
						"footer": {
							"type": "box",
							"layout": "vertical",
							"spacing": "sm",
							"contents": [
								{
									"type": "box",
									"layout": "horizontal",
									"contents": [
										{
											"type": "button",
											"style": "link",
											"height": "sm",
											"action": {
												"type": "message",
												"label": "自分のメニューのみ",
												"text": "自分のメニューのみ"
											},
											"color": "#000000"
										},
										{
											"type": "button",
											"style": "link",
											"height": "sm",
											"action": {
												"type": "message",
												"label": "お気に入りも含む",
												"text": "お気に入りも含む"
											},
											"color": "#000000"
										}
									],
									"spacing": "sm"
								}
							]
						}
					}`
					jsonData := []byte(json)
					container, err := linebot.UnmarshalFlexMessageJSON(jsonData)
					if err != nil {
						log.Println("error")
						log.Println(err)
						return err
					}
					flexMessage := linebot.NewFlexMessage("alt text", container)

					if err = h.lineUC.ReplyFlexMessage(replyToken, flexMessage); err != nil {
						log.Println(err)
						return err
					}

				} else if message.Text == "自分のメニューのみ" {
					// TODO: Context的なので持たせるか何かで今いる情報によって分岐させたい
					userID := event.Source.UserID

					var recommendedMenuList []string

					user, err := h.userUC.GetUserByLineUserID(userID)
					if err != nil {
						log.Println(err)
						return err
					}

					flexMessage, err := h.recommendUC.RecommendMyMenu(user.ID, recommendedMenuList)
					if err != nil {
						log.Println(err)
						return err
					}

					if flexMessage == nil {
						errMsg := "おすすめできるメニューがありません。メニューを登録してください。"
						if err = h.lineUC.ReplyMessage(replyToken, linebot.NewTextMessage(errMsg)); err != nil {
							log.Print(err)
							return err
						}
						return nil
					}
					if err = h.lineUC.ReplyFlexMessage(replyToken, flexMessage); err != nil {
						log.Println(err)
						return err
					}

				} else if message.Text == "お気に入りも含む" {
					userID := event.Source.UserID

					var recommendedMenuList []string

					user, err := h.userUC.GetUserByLineUserID(userID)
					if err != nil {
						log.Println(err)
						return err
					}

					flexMessage, err := h.recommendUC.RecommendMyMenuAndLikeMenu(user.ID, recommendedMenuList)
					if err != nil {
						log.Println(err)
						return err
					}
					if flexMessage == nil {
						errMsg := "おすすめできるメニューがありません。メニューを登録してください。"
						if err = h.lineUC.ReplyMessage(replyToken, linebot.NewTextMessage(errMsg)); err != nil {
							log.Print(err)
							return err
						}
						return nil
					}
					if err = h.lineUC.ReplyFlexMessage(replyToken, flexMessage); err != nil {
						log.Println(err)
						return err
					}

				}
			}
		}
	}
	return c.NoContent(http.StatusOK)
}
