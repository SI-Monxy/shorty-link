package shorturl

// OutputPort（プレゼンターインターフェース）
type OutputPort interface {
	Present(output *ShortenOutput) *ShortenOutput
}
