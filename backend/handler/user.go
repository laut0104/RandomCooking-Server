package handler

import (
	"log"

	usecase "github.com/laut0104/RandomCooking/usecase/interactor"
	_ "github.com/lib/pq"

	"net/http"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userUC *usecase.UserUseCase
}

func NewUserHandler(userUC *usecase.UserUseCase) *UserHandler {
	return &UserHandler{userUC: userUC}
}

func (h *UserHandler) GetUserByID(c echo.Context) error {
	id := c.Param("id")
	user, err := h.userUC.GetUserByID(id)
	if err != nil {
		log.Println(err)
		return err
	}

	return c.JSON(http.StatusOK, user)
}

func (h *UserHandler) GetUserByLineUserID(c echo.Context) error {
	lineid := c.QueryParam("lineuserid")
	user, err := h.userUC.GetUserByLineUserID(lineid)
	if err != nil {
		log.Println(err)
		return err
	}
	return c.JSON(http.StatusOK, user)
}
