package handler

import (
	"github.com/laut0104/RandomCooking/domain/entity"
	usecase "github.com/laut0104/RandomCooking/usecase/interactor"
	_ "github.com/lib/pq"

	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type LikeHandler struct {
	likeUC *usecase.LikeUseCase
}

func NewLikeHandler(likeUC *usecase.LikeUseCase) *LikeHandler {
	return &LikeHandler{likeUC: likeUC}
}

func (h *LikeHandler) GetLikesMenuByUserID(c echo.Context) error {
	userID := c.Param("uid")
	likesMenu, err := h.likeUC.GetLikesMenuByUserID(userID)
	if err != nil {
		log.Println(err)
		return err
	}

	return c.JSON(http.StatusOK, likesMenu)
}

func (h *LikeHandler) GetLikeByUniqueKey(c echo.Context) error {
	userID := c.Param("uid")
	menuID := c.Param("mid")
	like, err := h.likeUC.GetLikeByUniqueKey(userID, menuID)
	if err != nil {
		log.Println(err)
		return err
	}

	return c.JSON(http.StatusOK, like)
}

func (h *LikeHandler) AddLike(c echo.Context) error {
	var like *entity.Like

	if err := c.Bind(&like); err != nil {
		log.Println(err)
		return err
	}

	if err := h.likeUC.AddLike(like); err != nil {
		log.Println(err)
		return err
	}

	return c.JSON(http.StatusOK, like)
}

func (h *LikeHandler) DeleteLike(c echo.Context) error {
	id := c.Param("id")

	if err := h.likeUC.DeleteLike(id); err != nil {
		log.Println(err)
		return err
	}

	return c.NoContent(http.StatusOK)
}
