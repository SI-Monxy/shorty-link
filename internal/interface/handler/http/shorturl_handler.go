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

// Shorten godoc
// @Summary URLを短縮する
// @Tags shorturl
// @Accept json
// @Produce json
// @Param body body shorturl.ShortenInput true "入力パラメータ"
// @Success 200 {object} shorturl.ShortenOutput
// @Failure 400 {object} map[string]string
// @Router /api/v1/shorten [post]
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
	output, err := h.usecase.Shorten(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, output)
}

// Redirect godoc
// @Summary 短縮URLから元のURLへリダイレクト
// @Tags shorturl
// @Produce plain
// @Param code path string true "短縮URLコード"
// @Success 302 {string} string "リダイレクト"
// @Failure 404 {object} map[string]string
// @Router /{code} [get]
func (h *ShortURLHandler) Redirect(c *gin.Context) {
	code := c.Param("code")

	input := u.RedirectInput{Code: code}
	output, err := h.usecase.Redirect(input)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}

	c.Redirect(http.StatusFound, output.OriginalURL)
}
