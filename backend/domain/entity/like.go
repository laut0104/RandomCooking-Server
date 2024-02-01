package entity

type (
	Like struct {
		ID     string `json:"id" param:"id"`
		UserID int    `json:"userid" param:"uid"`
		MenuID int    `json:"menuid" param:"menuid"`
	}
)

func NewLike(userID int, menuID int) *Like {
	return &Like{
		UserID: userID,
		MenuID: menuID,
	}
}
