package shorturl

// 入力DTO（リクエストを受け取るための構造体）
type ShortenInput struct {
	OriginalURL string `json:"original_url"`
}

type RedirectInput struct {
	Code string
}
