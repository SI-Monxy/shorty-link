package shorturl

// OutputPort（プレゼンターインターフェース）
type OutputPort interface {
	PresentShorten(output *ShortenOutput) *ShortenOutput
	PresentRedirect(output *RedirectOutput) *RedirectOutput
}
