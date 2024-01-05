package usecase

import (
	"log"

	"github.com/laut0104/RandomCooking/domain/entity"
	"github.com/laut0104/RandomCooking/domain/repository"
)

type MenuUseCase struct {
	menuRepo repository.Menu
}

func NewMenuUseCase(menuRepo repository.Menu) *MenuUseCase {
	return &MenuUseCase{menuRepo: menuRepo}
}

func (u *MenuUseCase) GetMenuByID(id, userID string) (*entity.Menu, error) {
	menu, err := u.menuRepo.FindByID(id, userID)
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
