package main

import (
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/line/line-bot-sdk-go/linebot"

	"log"

	"github.com/laut0104/RandomCooking/handler"
	"github.com/laut0104/RandomCooking/infra/postgresql"
	database "github.com/laut0104/RandomCooking/infra/postgresql/repository"
	usecase "github.com/laut0104/RandomCooking/usecase/interactor"

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

	if os.Getenv("ENV") != "PROD" {
		if err := godotenv.Load(".env.local"); err != nil {
			log.Println(err)
			log.Fatal(err)
		}
	}
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

	bot, err := linebot.New(
		os.Getenv("LINE_BOT_CHANNEL_SECRET"),
		os.Getenv("LINE_BOT_CHANNEL_TOKEN"),
	)
	if err != nil {
		log.Println(err)
		log.Fatal(err)
	}

	// TODO: ここら辺の生成周りは後でどっかでまとめたい
	likeRepo := database.NewLikeRepository(db)
	menuRepo := database.NewMenuRepository(db)
	userRepo := database.NewUserRepository(db)

	authUC := usecase.NewAuthUseCase()
	likeUC := usecase.NewLikeUseCase(likeRepo)
	lineUC := usecase.NewLineUseCase(userRepo, bot)
	menuUC := usecase.NewMenuUseCase(menuRepo, likeRepo)
	recommendUC := usecase.NewRecommendUseCase(menuRepo)
	storageUC := usecase.NewStorageUseCase()
	userUC := usecase.NewUserUseCase(userRepo)
	authHandler := handler.NewAuthHandler(authUC)
	likeHandler := handler.NewLikeHandler(likeUC)
	lineHandler := handler.NewLineHandler(userUC, lineUC, recommendUC)
	menuHandler := handler.NewMenuHandler(menuUC)
	recommendHandler := handler.NewRecommendHandler(userUC, recommendUC, lineUC)
	storageHandler := handler.NewStorageHandler(storageUC)
	userHandler := handler.NewUserHandler(userUC)

	e.POST("/callback", lineHandler.LineEvent)
	e.GET("/auth/line/callback", authHandler.Login)

	// ログインしているユーザーでなければアクセスできない
	r := e.Group("/api")
	r.Use(echojwt.JWT([]byte(os.Getenv("JWT_SECRET_KEY"))))
	r.GET("/user/:id", userHandler.GetUserByID)
	r.GET("/user", userHandler.GetUserByLineUserID)
	r.GET("/menu/:id", menuHandler.GetMenu)
	r.GET("/menus/:uid", menuHandler.GetMenusByUserID)
	r.POST("/menu/:uid", menuHandler.AddMenu)
	r.PUT("/menu/:uid/:id", menuHandler.UpdateMenu)
	r.DELETE("/menu/:uid/:id", menuHandler.DeleteMenu)
	r.POST("/image/:uid", storageHandler.UploadImage)
	r.POST("/recommend/menu/:uid", recommendHandler.RecommendMenu)
	r.GET("/explore/menu/:uid", menuHandler.ExploreMenu)
	r.GET("/like/:uid/:mid", likeHandler.GetMenuByUniqueKey)
	r.POST("/like", likeHandler.AddLike)
	r.DELETE("/like/:id", likeHandler.DeleteLike)

	// サーバーをポート番号8080で起動
	e.Logger.Fatal(e.Start(":8080"))
}
