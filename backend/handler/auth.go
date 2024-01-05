package handler

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	usecase "github.com/laut0104/RandomCooking/usecase/interactor"
	_ "github.com/lib/pq"
)

type AuthHandler struct {
	authUC *usecase.AuthUseCase
}

// NewAuthHandler はAuthHandlerのポインタを生成する関数です。
func NewAuthHandler(authUC *usecase.AuthUseCase) *AuthHandler {
	return &AuthHandler{authUC: authUC}
}

func (h *AuthHandler) Login(c echo.Context) error {
	code := c.QueryParam("access_token")

	token, err := h.authUC.Login(code)
	if err != nil {
		log.Println(err)
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": token,
	})
}
