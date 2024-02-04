package database

import (
	"database/sql"
	"log"
	"strconv"

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

func (r *LikeRepository) FindByUniqueKey(userID, menuID string) (*entity.Like, error) {
	var dto likeDTO
	if err := r.db.QueryRow(`SELECT * FROM likes where userid=$1 AND menuid=$2`, userID, menuID).Scan(&dto.ID, &dto.UserID, &dto.MenuID); err != nil {
		log.Println(err)
		return nil, err
	}

	return &entity.Like{
		ID:     dto.ID,
		UserID: dto.UserID,
		MenuID: dto.MenuID,
	}, nil
}

func (r *LikeRepository) FindAllByUserID(userID string) ([]*entity.Like, error) {
	rows, err := r.db.Query(`SELECT * FROM likes where userid=$1`, userID)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	likes := []*entity.Like{}
	for rows.Next() {
		dto := &entity.Like{}
		if err := rows.Scan(&dto.ID, &dto.UserID, &dto.MenuID); err != nil {
			log.Println(err)
			return nil, err
		}
		likes = append(likes, dto)
	}
	return likes, nil
}

func (r *LikeRepository) Store(like *entity.Like) error {
	dto := &likeDTO{
		UserID: like.UserID,
		MenuID: like.MenuID,
	}

	var id int
	err := r.db.QueryRow(`INSERT INTO likes (userid, menuid) VALUES($1, $2) RETURNING id`, dto.UserID, dto.MenuID).Scan(&id)
	if err != nil {
		log.Println(err)
		return err
	}

	if err != nil {
		log.Println(err)
		return err
	}
	like.ID = strconv.Itoa(id)

	return nil
}

func (r *LikeRepository) Delete(id string) error {
	_, err := r.db.Exec(`DELETE FROM likes WHERE id=$1`, id)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

type likeDTO struct {
	ID     string `json:"id" param:"id"`
	UserID string `json:"userid" param:"uid"`
	MenuID string `json:"menuid" param:"menuid"`
}
