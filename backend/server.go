package main

import (
	"os"

	_ "github.com/lib/pq"

	"log"

	"github.com/laut0104/RandomCooking/handler"
	"github.com/laut0104/RandomCooking/infra/postgresql"
	database "github.com/laut0104/RandomCooking/infra/postgresql/repository"
	usecase "github.com/laut0104/RandomCooking/usecase/interactor"

	"github.com/joho/godotenv"
	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// インスタンスを作成
	e := echo.New()

	// ミドルウェアを設定
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORS())

	if err := godotenv.Load(".env"); err != nil {
		log.Println(err)
		log.Fatal(err)
	}

	db, err := postgresql.New()
	if err != nil {
		log.Println(err)
		log.Fatal(err)
	}
	defer db.Close()

	// TODO: ここら辺の生成周りは後でどっかでまとめたい
	userRepo := database.NewUserRepository(db)
	menuRepo := database.NewMenuRepository(db)

	authUC := usecase.NewAuthUseCase()
	userUC := usecase.NewUserUseCase(userRepo)
	menuUC := usecase.NewMenuUseCase(menuRepo)
	lineUC := usecase.NewLineUseCase(userRepo)
	storageUC := usecase.NewStorageUseCase()
	authHandler := handler.NewAuthHandler(authUC)
	userHandler := handler.NewUserHandler(userUC)
	menuHandler := handler.NewMenuHandler(menuUC)
	lineHandler := handler.NewLineHandler(lineUC)
	storageHandler := handler.NewStorageHandler(storageUC)

	e.POST("/callback", lineHandler.LineEvent)
	e.GET("/auth/line/callback", authHandler.Login)

	// ログインしているユーザーでなければアクセスできない
	r := e.Group("/api")
	r.Use(echojwt.JWT([]byte(os.Getenv("JWT_SECRET_KEY"))))
	r.GET("/user/:id", userHandler.GetUserByID)
	r.GET("/user", userHandler.GetUserByLineUserID)
	r.GET("/menu/:uid/:id", menuHandler.GetMenu)
	r.GET("/menus/:uid", menuHandler.GetMenusByUserID)
	r.POST("/menu/:uid", menuHandler.AddMenu)
	r.PUT("/menu/:uid/:id", menuHandler.UpdateMenu)
	r.DELETE("/menu/:uid/:id", menuHandler.DeleteMenu)
	r.POST("/image/:uid", storageHandler.UploadImage)

	// サーバーをポート番号8080で起動
	e.Logger.Fatal(e.Start(":8080"))
}
