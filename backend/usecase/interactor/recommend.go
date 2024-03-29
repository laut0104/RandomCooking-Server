package usecase

import (
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/laut0104/RandomCooking/domain/entity"
	"github.com/laut0104/RandomCooking/domain/repository"
	"github.com/line/line-bot-sdk-go/linebot"
)

type RecommendUseCase struct {
	menuRepo repository.Menu
	likeRepo repository.Like
}

func NewRecommendUseCase(menuRepo repository.Menu, likeRepo repository.Like) *RecommendUseCase {
	return &RecommendUseCase{menuRepo: menuRepo, likeRepo: likeRepo}
}

func (u *RecommendUseCase) RecommendMyMenu(userID string, recommendedList []string) (*linebot.FlexMessage, error) {
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

	link := os.Getenv("LIFF_URL") + "/menu/" + menuList[rnd].ID
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
					"text": "` + menuList[rnd].MenuName + `",
					"wrap": true
					}`
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
								"data": "MyMenu,` + recommendedMenu + `",
								"displayText": "他のは？"
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
	flexMessage := linebot.NewFlexMessage("このメニューはどうですか？", container)

	return flexMessage, nil
}

func (u *RecommendUseCase) RecommendMyMenuAndLikeMenu(userID string, recommendedList []string) (*linebot.FlexMessage, error) {
	menus, err := u.menuRepo.FindAllByUserID(userID)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	likes, err := u.likeRepo.FindLikesMenuByUserID(userID)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	recommendMenuMap := make(map[string]struct{}, len(menus)+len(likes))
	recommendedMenu := ""
	for _, menuID := range recommendedList {
		recommendMenuMap[menuID] = struct{}{}
		recommendedMenu = recommendedMenu + menuID + ","
	}
	menuList := make([]*entity.Menu, 0, len(menus)+len(likes))
	for _, menu := range menus {
		if _, ok := recommendMenuMap[menu.ID]; !ok {
			menuList = append(menuList, menu)
		}
	}
	for _, like := range likes {
		if _, ok := recommendMenuMap[like.MenuID]; !ok {
			menu := &entity.Menu{
				ID:          like.MenuID,
				MenuName:    like.MenuName,
				ImageUrl:    like.ImageUrl,
				Ingredients: like.Ingredients,
				Quantities:  like.Quantities,
				Recipes:     like.Recipes,
			}
			menuList = append(menuList, menu)
		}
	}
	log.Println(menuList)
	if len(menuList) == 0 {
		return nil, nil
	}
	rand.NewSource(time.Now().UnixNano())
	rnd := rand.Intn(len(menuList))

	link := os.Getenv("LIFF_URL") + "/menu/" + menuList[rnd].ID
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
					"text": "` + menuList[rnd].MenuName + `",
					"wrap": true
					}`
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
								"data": "MyMenuAndLikeMenu,` + recommendedMenu + `",
								"displayText": "他のは？"
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
	flexMessage := linebot.NewFlexMessage("このメニューはどうですか？", container)

	return flexMessage, nil
}
