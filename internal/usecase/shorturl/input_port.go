package shorturl

// InputPort（ユースケースインターフェース）
type InputPort interface {
	Execute(input ShortenInput) (*ShortenOutput, error)
}
