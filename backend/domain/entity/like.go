package entity

type (
	Like struct {
		ID     string `json:"id" param:"id"`
		UserID string `json:"userid" param:"uid"`
		MenuID string `json:"menuid" param:"menuid"`
	}
)

func NewLike(userID string, menuID string) *Like {
	return &Like{
		UserID: userID,
		MenuID: menuID,
	}
}
