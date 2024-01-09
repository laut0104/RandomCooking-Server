package usecase

import (
	"context"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"os"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

type StorageUseCase struct {
}

func NewStorageUseCase() *StorageUseCase {
	return &StorageUseCase{}
}

func (u *StorageUseCase) UploadToGCS(file multipart.File, fileName string) (string, error) {
	credentialFilePath := "./credentials.json"
	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithCredentialsFile(credentialFilePath))
	bucketName := os.Getenv("BUCKET_NAME")
	if err != nil {
		log.Println("cannot create client : %w", err)
		return "", err
	}

	// // バケットオブジェクトの作成
	bucket := client.Bucket(bucketName)
	// バケット内のアップロード先のオブジェクトを作成
	obj := bucket.Object(fileName)
	// ファイルをアップロード
	wc := obj.NewWriter(ctx)
	if _, err := io.Copy(wc, file); err != nil {
		log.Println(err)
		return "", nil
	}
	if err := wc.Close(); err != nil {
		log.Println(err)
		return "", nil
	}

	log.Println("file upload success")
	uploadFilePath := fmt.Sprintf("https://storage.googleapis.com/%s/%s", bucketName, fileName)

	return uploadFilePath, nil
}
