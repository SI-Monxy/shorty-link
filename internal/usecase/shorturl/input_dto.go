package shorturl

// 入力DTO（リクエストを受け取るための構造体）
type ShortenInput struct {
	OriginalURL string `json:"original_url" example:"https://example.com"`
}

type RedirectInput struct {
	Code string `json:"code" example:"abc123"`
}
