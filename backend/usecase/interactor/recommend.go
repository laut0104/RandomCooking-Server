package usecase

import (
	"log"
	"math/rand"
	"time"

	"github.com/laut0104/RandomCooking/domain/entity"
	"github.com/laut0104/RandomCooking/domain/repository"
	"github.com/line/line-bot-sdk-go/linebot"
)

type RecommendUseCase struct {
	menuRepo repository.Menu
}

func NewRecommendUseCase(menuRepo repository.Menu) *RecommendUseCase {
	return &RecommendUseCase{menuRepo: menuRepo}
}

func (u *RecommendUseCase) RecommendMenu(userID string, recommendedList []string) (*linebot.FlexMessage, error) {
	menus, err := u.menuRepo.FindAllByUserID(userID)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println(menus)
	recommendMenuMap := make(map[string]struct{}, len(menus))
	recommendedMenu := ""
	for _, menuID := range recommendedList {
		recommendMenuMap[menuID] = struct{}{}
		recommendedMenu = recommendedMenu + menuID + ","
	}
	menuList := make([]*entity.Menu, 0, len(menus))
	for _, menu := range menus {
		if _, ok := recommendMenuMap[menu.ID]; !ok {
			menuList = append(menuList, menu)
		}
	}
	log.Println(menuList)
	if len(menuList) == 0 {
		return nil, nil
	}
	rand.NewSource(time.Now().UnixNano())
	rnd := rand.Intn(len(menuList))

	link := "https://liff.line.me/1660690567-wegZZboy/menu/" + menuList[rnd].ID
	recommendedMenu = recommendedMenu + menuList[rnd].ID

	json := `{
		"type": "bubble",
		"header": {
			"type": "box",
			"layout": "vertical",
			"contents": [
				{
					"type": "text",
					"text": "今日のメニュー",
					"size": "xl",
					"margin": "none",
					"style": "normal",
					"align": "center"
				}
			],
			"spacing": "none",
			"margin": "none",
			"height": "60px"
		},
		"body": {
			"type": "box",
			"layout": "vertical",
			"contents": [
				{
					"type": "text",
					"weight": "bold",
					"size": "xl",
					"text": "` + menuList[rnd].MenuName +
		`"}`
	if menuList[rnd].ImageUrl != "" {
		imageBody := `,{
			"type": "image",
			"url": "` + menuList[rnd].ImageUrl + `",
			"size": "full"
		}`
		json = json + imageBody
	}
	json = json +
		`]
		},
		"footer": {
			"type": "box",
			"layout": "vertical",
			"contents": [
				{
					"type": "box",
					"layout": "horizontal",
					"contents": [
						{
							"type": "button",
							"action": {
								"type": "postback",
								"label": "他のは？",
								"data": "` + recommendedMenu + `"
							},
							"flex": 1,
							"style": "secondary",
							"margin": "md",
							"gravity": "center"
						},
						{
							"type": "separator",
							"color": "#FFFFFF",
							"margin": "10px"
						},
						{
							"type": "button",
							"action": {
								"type": "uri",
								"label": "詳細を見る",
								"uri": "` + link + `"
							},
							"flex": 1,
							"style": "secondary",
							"margin": "none",
							"gravity": "center"
						}
					],
					"height": "60px"
				}
			]
		}
	}`
	jsonData := []byte(json)
	container, err := linebot.UnmarshalFlexMessageJSON(jsonData)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	flexMessage := linebot.NewFlexMessage("alt text", container)

	return flexMessage, nil
}
