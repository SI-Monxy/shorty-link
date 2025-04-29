package main

import (
	"log"
	"shorty-link/internal/config"
	"shorty-link/internal/infrastructure/db/mysql"
	handler "shorty-link/internal/interface/handler/http"
	"shorty-link/internal/presenter"
	"shorty-link/internal/usecase/shorturl"

	"github.com/gin-gonic/gin"
)

func main() {
	// 設定読み込み
	db := config.NewDB()

	r := gin.Default()

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
