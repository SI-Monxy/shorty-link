package http

import (
	"net/http"
	u "shorty-link/internal/usecase/shorturl"

	"github.com/gin-gonic/gin"
)

// GinのHandler（Controller役割）
type ShortURLHandler struct {
	usecase u.InputPort
}

func NewShortURLHandler(u u.InputPort) *ShortURLHandler {
	return &ShortURLHandler{usecase: u}
}

func (h *ShortURLHandler) Shorten(c *gin.Context) {
	var input u.ShortenInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}
	if err := u.ValidateInput(input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	output, err := h.usecase.Execute(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, output)
}

func (h *ShortURLHandler) Redirect(c *gin.Context) {
	// 簡易版：本来はリダイレクト用UseCaseを呼び出すべき
	code := c.Param("code")
	c.Redirect(http.StatusFound, "https://example.com/"+code)
}
