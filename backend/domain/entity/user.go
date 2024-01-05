package entity

type (
	// User はログインしているユーザを表します。
	User struct {
		ID         string
		LineUserID string
		UserName   string
	}
)

// NewUser はUserのポインタを生成する関数です。
func NewUser(lineUserID, userName string) *User {
	return &User{
		LineUserID: lineUserID,
		UserName:   userName,
	}
}
