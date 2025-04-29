package presenter_test

import (
	"shorty-link/internal/presenter"
	"shorty-link/internal/usecase/shorturl"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPresentShorten(t *testing.T) {
	p := presenter.NewShortURLPresenter()

	input := &shorturl.ShortenOutput{
		ShortURL: "http://localhost:8080/abc123",
	}

	output := p.PresentShorten(input)

	assert.Equal(t, input.ShortURL, output.ShortURL)
}

func TestPresentRedirect(t *testing.T) {
	p := presenter.NewShortURLPresenter()

	input := &shorturl.RedirectOutput{
		OriginalURL: "https://example.com",
	}

	output := p.PresentRedirect(input)

	assert.Equal(t, input.OriginalURL, output.OriginalURL)
}
