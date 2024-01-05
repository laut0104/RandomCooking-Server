package entity

type (
	// Menu は料理の内容を表します。
	Menu struct {
		ID       string `json:"id" param:"id"`
		UserID   int    `json:"userid" param:"uid"`
		MenuName string `json:"menuname" param:"menuname"`
		// Materials  []string
		// Quantities []string
		Recipes []string `json:"recipes" param:"recipes"`
	}
)

func NewMenu(userID int, menuName string, recipes []string) *Menu {
	return &Menu{
		UserID:   userID,
		MenuName: menuName,
		// Materials:  materials,
		// Quantities: quantities,
		Recipes: recipes,
	}
}
