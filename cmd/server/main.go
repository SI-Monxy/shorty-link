package main

// @title shorty-link API
// @version 1.0
// @description URL短縮サービスのAPIドキュメント
// @host localhost:8080
// @BasePath /

import (
	"log"
	"shorty-link/internal/config"
	"shorty-link/internal/infrastructure/db/mysql"
	handler "shorty-link/internal/interface/handler/http"
	"shorty-link/internal/presenter"
	"shorty-link/internal/usecase/shorturl"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "shorty-link/docs"
)

func main() {
	// 設定読み込み
	db := config.NewDB()

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(200, "OK")
	})

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// DI (依存性注入)
	repo := mysql.NewShortURLRepository(db)
	pres := presenter.NewShortURLPresenter()
	uc := shorturl.NewInteractor(repo, pres)
	h := handler.NewShortURLHandler(uc)

	// ルーティング登録
	h.RegisterRoutes(r)

	log.Println("Server starting at :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
