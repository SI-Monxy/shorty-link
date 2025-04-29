package presenter

import u "shorty-link/internal/usecase/shorturl"

// Presenter実装（OutputPort実装）
type ShortURLPresenter struct{}

func NewShortURLPresenter() *ShortURLPresenter {
	return &ShortURLPresenter{}
}

func (p *ShortURLPresenter) Present(output *u.ShortenOutput) *u.ShortenOutput {
	// 本来はここで加工・整形する（今回はそのまま返す）
	return output
}
