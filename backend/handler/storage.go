package handler

import (
	usecase "github.com/laut0104/RandomCooking/usecase/interactor"
	_ "github.com/lib/pq"

	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type StorageHandler struct {
	storageUC *usecase.StorageUseCase
}

func NewStorageHandler(storageUC *usecase.StorageUseCase) *StorageHandler {
	return &StorageHandler{storageUC: storageUC}
}

func (h *StorageHandler) UploadImage(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		log.Println(err)
		return err
	}

	// アップロードするファイルを開く
	src, err := file.Open()
	if err != nil {
		log.Println(err)
		return err
	}
	defer src.Close()

	fileName, err := h.storageUC.UploadToGCS(src, file.Filename)
	if err != nil {
		log.Println(err)
		return err
	}

	return c.JSON(http.StatusOK, fileName)
}
