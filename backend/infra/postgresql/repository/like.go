package database

import (
	"database/sql"
	"log"

	"github.com/laut0104/RandomCooking/domain/entity"
	"github.com/laut0104/RandomCooking/domain/repository"
)

var _ repository.Like = &LikeRepository{}

type LikeRepository struct {
	db *sql.DB
}

func NewLikeRepository(db *sql.DB) *LikeRepository {
	return &LikeRepository{db: db}
}

func (r *LikeRepository) FindAllByUserID(userID string) ([]*entity.Like, error) {
	rows, err := r.db.Query(`SELECT * FROM likes where userid=$1`, userID)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	menus := []*entity.Like{}
	for rows.Next() {
		dto := &entity.Like{}
		if err := rows.Scan(&dto.ID, &dto.UserID, &dto.MenuID); err != nil {
			log.Println(err)
			return nil, err
		}
		menus = append(menus, dto)
	}
	return menus, nil
}

// type likeDTO struct {
// 	ID     string `json:"id" param:"id"`
// 	UserID int    `json:"userid" param:"uid"`
// 	MenuID int    `json:"menuid" param:"menuid"`
// }
