package shorturl

// 出力DTO（レスポンスとして返す構造体）
type ShortenOutput struct {
	ShortURL string `json:"short_url" example:"http://localhost:8080/abc123"`
}

type RedirectOutput struct {
	OriginalURL string `json:"original_url" example:"https://example.com"`
}
