package shorturl

// InputPort（ユースケースインターフェース）
type InputPort interface {
	Shorten(input ShortenInput) (*ShortenOutput, error)
	Redirect(input RedirectInput) (*RedirectOutput, error)
}
