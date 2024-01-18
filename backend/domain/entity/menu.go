package entity

type (
	// Menu は料理の内容を表します。
	Menu struct {
		ID          string   `json:"id" param:"id"`
		UserID      int      `json:"userid" param:"uid"`
		MenuName    string   `json:"menuname" param:"menuname"`
		ImageUrl    string   `json:"imageurl" param:"imageurl"`
		Ingredients []string `json:"ingredients" param:"ingredients"`
		Quantities  []string `json:"quantities" param:"quantities"`
		Recipes     []string `json:"recipes" param:"recipes"`
	}
)

func NewMenu(userID int, menuName string, imageUrl string, ingredients []string, quantities []string, recipes []string) *Menu {
	return &Menu{
		UserID:      userID,
		MenuName:    menuName,
		ImageUrl:    imageUrl,
		Ingredients: ingredients,
		Quantities:  quantities,
		Recipes:     recipes,
	}
}
