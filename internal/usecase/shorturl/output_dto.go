package shorturl

// 出力DTO（レスポンスとして返す構造体）
type ShortenOutput struct {
	ShortURL string `json:"short_url"`
}
