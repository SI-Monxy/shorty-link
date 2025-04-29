package presenter

import "shorty-link/internal/usecase/shorturl"

// Presenter実装（OutputPort実装）
type ShortURLPresenter struct{}

func NewShortURLPresenter() *ShortURLPresenter {
	return &ShortURLPresenter{}
}

func (p *ShortURLPresenter) PresentShorten(output *shorturl.ShortenOutput) *shorturl.ShortenOutput {
	return output
}

func (p *ShortURLPresenter) PresentRedirect(output *shorturl.RedirectOutput) *shorturl.RedirectOutput {
	return output
}
