package http

import (
	"github.com/gin-gonic/gin"
)

// Ginルーター設定
func (h *ShortURLHandler) RegisterRoutes(r *gin.Engine) {
	r.POST("/api/v1/shorten", h.Shorten)
	r.GET("/:code", h.Redirect)
}
