package database

import (
	"database/sql"
	"log"

	"github.com/laut0104/RandomCooking/domain/entity"
	"github.com/laut0104/RandomCooking/domain/repository"
)

var _ repository.User = &UserRepository{}

// UserRepository は repository.UserRepository を満たす構造体です
type UserRepository struct {
	db *sql.DB
}

// NewUserRepository はUserRepositoryのポインタを生成する関数です
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

// FindByID は指定されたIDを持つユーザをDBから取得します
func (r *UserRepository) FindByID(id string) (*entity.User, error) {
	var dto userDTO
	if err := r.db.QueryRow(`SELECT * FROM users where id=$1`, id).Scan(&dto.ID, &dto.LineUserID, &dto.UserName); err != nil {
		log.Println(err)
		return nil, err
	}

	return &entity.User{
		ID:         dto.ID,
		LineUserID: dto.LineUserID,
		UserName:   dto.UserName,
	}, nil
}

// FindByLineUserID はlineUserIDを持つユーザを取得します。
func (r *UserRepository) FindByLineUserID(lineUserID string) (*entity.User, error) {
	var dto userDTO

	if err := r.db.QueryRow(`SELECT * FROM users where lineuserid=$1`, lineUserID).Scan(&dto.ID, &dto.LineUserID, &dto.UserName); err != nil {
		log.Println(err)
		return nil, err
	}

	return &entity.User{
		ID:         dto.ID,
		LineUserID: dto.LineUserID,
		UserName:   dto.UserName,
	}, nil
}

// Store はユーザを新規保存します。
func (r *UserRepository) Store(user *entity.User) error {
	dto := &userDTO{
		LineUserID: user.LineUserID,
		UserName:   user.UserName,
	}

	// 重複の場合、Insertしない
	_, err := r.db.Exec(`INSERT INTO users (lineuserid, username) VALUES($1, $2) ON CONFLICT DO NOTHING`, dto.LineUserID, dto.UserName)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

// Update はユーザの情報を更新します。
func (r *UserRepository) Update(user *entity.User) error {
	dto := &userDTO{
		ID:         user.ID,
		LineUserID: user.LineUserID,
		UserName:   user.UserName,
	}

	_, err := r.db.Exec(`UPDATE users SET(lineuserid, username)=($1, $2) WHERE id=$3`, dto.LineUserID, dto.UserName, dto.ID)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (r *UserRepository) Delete(id string) error {
	_, err := r.db.Exec(`DELETE FROM users WHERE id=$1`, id)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

type userDTO struct {
	ID         string
	LineUserID string
	UserName   string
}
