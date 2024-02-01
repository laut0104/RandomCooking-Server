package usecase

import (
	"log"
	"strconv"

	"github.com/laut0104/RandomCooking/domain/entity"
	"github.com/laut0104/RandomCooking/domain/repository"
)

type MenuUseCase struct {
	menuRepo repository.Menu
	likeRepo repository.Like
}

func NewMenuUseCase(menuRepo repository.Menu, likeRepo repository.Like) *MenuUseCase {
	return &MenuUseCase{menuRepo: menuRepo, likeRepo: likeRepo}
}

func (u *MenuUseCase) GetMenuByID(id string) (*entity.Menu, error) {
	menu, err := u.menuRepo.FindByID(id)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return menu, nil
}

func (u *MenuUseCase) GetMenusByUserID(userID string) ([]*entity.Menu, error) {
	menus, err := u.menuRepo.FindAllByUserID(userID)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return menus, nil
}

func (u *MenuUseCase) GetMenus() ([]*entity.Menu, error) {
	menus, err := u.menuRepo.FindAll()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return menus, nil
}

func (u *MenuUseCase) AddMenu(menu *entity.Menu) error {
	if err := u.menuRepo.Store(menu); err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (u *MenuUseCase) UpdateMenu(menu *entity.Menu) error {
	if err := u.menuRepo.Update(menu); err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (u *MenuUseCase) DeleteMenu(id, userID string) error {
	if err := u.menuRepo.Delete(id, userID); err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (u *MenuUseCase) ExploreMenus(userID string) ([]*entity.Menu, error) {
	menus, err := u.menuRepo.FindAllNotByUserID(userID)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	likes, err := u.likeRepo.FindAllByUserID(userID)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	exploreMenuList := make([]*entity.Menu, 0, len(menus))
	menuLikesMap := make(map[string]struct{}, len(likes))

	for _, like := range likes {
		menuLikesMap[strconv.Itoa(like.MenuID)] = struct{}{}
	}
	for _, menu := range menus {
		if _, ok := menuLikesMap[menu.ID]; !ok {
			exploreMenuList = append(exploreMenuList, menu)
		}
	}

	return exploreMenuList, nil
}
