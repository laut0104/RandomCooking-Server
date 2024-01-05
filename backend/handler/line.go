package handler

import (
	"log"
	"os"

	usecase "github.com/laut0104/RandomCooking/usecase/interactor"
	_ "github.com/lib/pq"
	"github.com/line/line-bot-sdk-go/linebot"

	"net/http"

	"github.com/labstack/echo/v4"
)

type LineHandler struct {
	lineUC *usecase.LineUseCase
}

func NewLineHandler(lineUC *usecase.LineUseCase) *LineHandler {
	return &LineHandler{lineUC: lineUC}
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
			errmsg := "正常にユーザー登録できませんでした\nブロックし、もう一度友達登録をお願いします"
			if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message)).Do(); err != nil {
				if _, err = bot.PushMessage(event.Source.UserID, linebot.NewTextMessage(errmsg)).Do(); err != nil {
					log.Print(err)
				}
				return err
			}
			if err := h.lineUC.Follow(event.Source.UserID, bot); err != nil {
				log.Println(err)
				if _, err = bot.PushMessage(event.Source.UserID, linebot.NewTextMessage(errmsg)).Do(); err != nil {
					log.Print(err)
				}
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

			// Text
			// 	case linebot.EventTypeMessage:
			// 		switch message := event.Message.(type) {
			// 		case *linebot.TextMessage:
			// 			// データベースの接続
			// 			connStr := "user=root dbname=randomcooking password=password host=postgres sslmode=disable"
			// 			db, err := sql.Open("postgres", connStr)
			// 			if err != nil {
			// 				log.Println(err)
			// 				return nil
			// 			}

			// 			rows, err := db.Query(`SELECT * FROM users where lineuserId=$1`, event.Source.UserID)
			// 			if err != nil {
			// 				log.Println(err)
			// 			}
			// 			defer rows.Close()

			// 			/*データベースに登録されていない場合*/
			// 			if !rows.Next() {
			// 				_, err := db.Exec(`INSERT INTO users (lineuserid, username) VALUES($1, $2)`, event.Source.UserID, message.Text)
			// 				if err != nil {
			// 					log.Println(err)
			// 					return nil
			// 				}
			// 			} else {
			// 				var id int
			// 				var lineuserid string
			// 				var username string
			// 				rows.Scan(&id, &lineuserid, &username)
			// 				/*メニューの登録*/
			// 				_, err := db.Exec(`INSERT INTO menus (userid, menuname, recipes) VALUES($1, $2, '{"テスト/", "メニューです/"}')`, id, message.Text)
			// 				if err != nil {
			// 					log.Println(err)
			// 					return nil
			// 				}
			// 			}

			// 			defer db.Close()

			// 			if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do(); err != nil {
			// 				log.Print(err)
			// 			}
			// 		case *linebot.StickerMessage:
			// 			replyMessage := fmt.Sprintf(
			// 				"sticker id is %s, stickerResourceType is %s", message.StickerID, message.StickerResourceType)
			// 			if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMessage)).Do(); err != nil {
			// 				log.Print(err)
			// 			}
			// 		}
		}
	}
	return c.NoContent(http.StatusOK)
}
