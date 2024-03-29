package database

import (
	"database/sql"
	"log"
	"strconv"

	"github.com/laut0104/RandomCooking/domain/entity"
	"github.com/laut0104/RandomCooking/domain/repository"
	"github.com/lib/pq"
)

var _ repository.Menu = &MenuRepository{}

type MenuRepository struct {
	db *sql.DB
}

func NewMenuRepository(db *sql.DB) *MenuRepository {
	return &MenuRepository{db: db}
}

func (r *MenuRepository) FindByID(id string) (*entity.Menu, error) {
	var dto menuDTO
	if err := r.db.QueryRow(`SELECT * FROM menus where id=$1`, id).Scan(&dto.ID, &dto.UserID, &dto.MenuName, &dto.ImageUrl, pq.Array(&dto.Ingredients), pq.Array(&dto.Quantities), pq.Array(&dto.Recipes)); err != nil {
		log.Println(err)
		return nil, err
	}

	return &entity.Menu{
		ID:          dto.ID,
		UserID:      dto.UserID,
		MenuName:    dto.MenuName,
		ImageUrl:    dto.ImageUrl,
		Ingredients: dto.Ingredients,
		Quantities:  dto.Quantities,
		Recipes:     dto.Recipes,
	}, nil
}

type Menus struct {
	Menus []*entity.Menu `json:"menus"`
}

func (r *MenuRepository) FindAllByUserID(userID string) ([]*entity.Menu, error) {
	rows, err := r.db.Query(`SELECT * FROM menus where userid=$1`, userID)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	menus := []*entity.Menu{}
	for rows.Next() {
		// メモリ効率良くない？
		dto := &entity.Menu{}
		if err := rows.Scan(&dto.ID, &dto.UserID, &dto.MenuName, &dto.ImageUrl, pq.Array(&dto.Ingredients), pq.Array(&dto.Quantities), pq.Array(&dto.Recipes)); err != nil {
			log.Println(err)
			return nil, err
		}
		menus = append(menus, dto)
	}
	return menus, nil
}

func (r *MenuRepository) FindAllNotByUserID(userID string) ([]*entity.Menu, error) {
	rows, err := r.db.Query(`SELECT * FROM menus where userid!=$1`, userID)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	menus := []*entity.Menu{}
	for rows.Next() {
		dto := &entity.Menu{}
		if err := rows.Scan(&dto.ID, &dto.UserID, &dto.MenuName, &dto.ImageUrl, pq.Array(&dto.Ingredients), pq.Array(&dto.Quantities), pq.Array(&dto.Recipes)); err != nil {
			log.Println(err)
			return nil, err
		}
		menus = append(menus, dto)
	}
	return menus, nil
}

func (r *MenuRepository) FindAll() ([]*entity.Menu, error) {
	rows, err := r.db.Query(`SELECT * FROM menus`)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	menus := []*entity.Menu{}
	for rows.Next() {
		dto := &entity.Menu{}
		if err := rows.Scan(&dto.ID, &dto.UserID, &dto.MenuName, &dto.ImageUrl, pq.Array(&dto.Ingredients), pq.Array(&dto.Quantities), pq.Array(&dto.Recipes)); err != nil {
			log.Println(err)
			return nil, err
		}
		menus = append(menus, dto)
	}

	return menus, nil
}

func (r *MenuRepository) Store(menu *entity.Menu) error {
	dto := &menuDTO{
		UserID:      menu.UserID,
		MenuName:    menu.MenuName,
		ImageUrl:    menu.ImageUrl,
		Ingredients: menu.Ingredients,
		Quantities:  menu.Quantities,
		Recipes:     menu.Recipes,
	}

	var id int
	err := r.db.QueryRow(`INSERT INTO menus (userid, menuname, imageurl, ingredients, quantities, recipes) VALUES($1, $2, $3, $4, $5, $6) RETURNING id`, dto.UserID, dto.MenuName, dto.ImageUrl, pq.Array(dto.Ingredients), pq.Array(dto.Quantities), pq.Array(dto.Recipes)).Scan(&id)
	if err != nil {
		log.Println(err)
		return err
	}

	if err != nil {
		log.Println(err)
		return err
	}
	menu.ID = strconv.Itoa(id)

	return nil
}

func (r *MenuRepository) Update(menu *entity.Menu) error {
	dto := &menuDTO{
		ID:          menu.ID,
		UserID:      menu.UserID,
		MenuName:    menu.MenuName,
		ImageUrl:    menu.ImageUrl,
		Ingredients: menu.Ingredients,
		Quantities:  menu.Quantities,
		Recipes:     menu.Recipes,
	}

	_, err := r.db.Exec(`UPDATE menus SET(menuname, imageurl, ingredients, quantities, recipes)=($1, $2, $3, $4, $5) WHERE id=$6 AND userid=$7`, dto.MenuName, dto.ImageUrl, pq.Array(dto.Ingredients), pq.Array(dto.Quantities), pq.Array(dto.Recipes), dto.ID, menu.UserID)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (r *MenuRepository) Delete(id, userID string) error {
	_, err := r.db.Exec(`DELETE FROM menus WHERE id=$1 AND userid=$2`, id, userID)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

type menuDTO struct {
	ID          string   `json:"id" param:"id"`
	UserID      int      `json:"userid" param:"uid"`
	MenuName    string   `json:"menuname" param:"menuname"`
	ImageUrl    string   `json:"imageurl" param:"imageurl"`
	Ingredients []string `json:"ingredients" param:"ingredients"`
	Quantities  []string `json:"quantities" param:"quantities"`
	Recipes     []string `json:"recipes" param:"recipes"`
}
