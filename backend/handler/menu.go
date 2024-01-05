package handler

import (
	"github.com/laut0104/RandomCooking/domain/entity"
	usecase "github.com/laut0104/RandomCooking/usecase/interactor"
	_ "github.com/lib/pq"

	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Menus struct {
	Menus []*entity.Menu `json:"menus"`
}

type MenuHandler struct {
	menuUC *usecase.MenuUseCase
}

func NewMenuHandler(menuUC *usecase.MenuUseCase) *MenuHandler {
	return &MenuHandler{menuUC: menuUC}
}

func (h *MenuHandler) GetMenu(c echo.Context) error {
	uid := c.Param("uid")
	id := c.Param("id")

	menu, err := h.menuUC.GetMenuByID(id, uid)
	if err != nil {
		log.Println(err)
		return err
	}

	return c.JSON(http.StatusOK, menu)
}

func (h *MenuHandler) GetMenusByUserID(c echo.Context) error {
	uid := c.Param("uid")
	menus, err := h.menuUC.GetMenusByUserID(uid)
	if err != nil {
		log.Println(err)
		return err
	}

	return c.JSON(http.StatusOK, menus)
}

func (h *MenuHandler) AddMenu(c echo.Context) error {
	var menu *entity.Menu
	uid, err := strconv.Atoi(c.Param("uid"))
	if err != nil {
		log.Println(err)
		return err
	}

	if err := c.Bind(&menu); err != nil {
		log.Println(err)
		return err
	}

	menu.UserID = uid

	if err := h.menuUC.AddMenu(menu); err != nil {
		log.Println(err)
		return err
	}

	return c.JSON(http.StatusOK, menu)
}

func (h *MenuHandler) UpdateMenu(c echo.Context) error {
	var menu *entity.Menu
	uid, err := strconv.Atoi(c.Param("uid"))
	id := c.Param("id")
	if err != nil {
		log.Println(err)
		return err
	}

	if err := c.Bind(&menu); err != nil {
		log.Println(err)
		return err
	}

	menu.UserID = uid
	menu.ID = id
	if err := h.menuUC.UpdateMenu(menu); err != nil {
		log.Println(err)
		return err
	}

	return c.JSON(http.StatusOK, menu)
}

func (h *MenuHandler) DeleteMenu(c echo.Context) error {
	uid := c.Param("uid")
	id := c.Param("id")

	if err := h.menuUC.DeleteMenu(id, uid); err != nil {
		log.Println(err)
		return err
	}

	return c.NoContent(http.StatusOK)
}
