package shorturl_test

import (
	entity "shorty-link/internal/domain/shorturl"
	usecase "shorty-link/internal/usecase/shorturl"
	"testing"

	"github.com/stretchr/testify/assert"
)

// モックRepo
type mockRepo struct {
	SaveFunc       func(*entity.ShortURL) error
	FindByCodeFunc func(string) (*entity.ShortURL, error)
}

func (m *mockRepo) Save(s *entity.ShortURL) error {
	return m.SaveFunc(s)
}

func (m *mockRepo) FindByCode(code string) (*entity.ShortURL, error) {
	return m.FindByCodeFunc(code)
}

// モックPresenter
type mockPresenter struct{}

func (m *mockPresenter) PresentShorten(output *usecase.ShortenOutput) *usecase.ShortenOutput {
	return output
}

func (m *mockPresenter) PresentRedirect(output *usecase.RedirectOutput) *usecase.RedirectOutput {
	return output
}

func TestInteractor_Shorten(t *testing.T) {
	mockR := &mockRepo{
		SaveFunc: func(s *entity.ShortURL) error {
			return nil
		},
	}
	mockP := &mockPresenter{}
	uc := usecase.NewInteractor(mockR, mockP)

	input := usecase.ShortenInput{OriginalURL: "https://example.com"}
	output, err := uc.Shorten(input)

	assert.NoError(t, err)
	assert.Contains(t, output.ShortURL, "http://localhost:8080/")
}

func TestInteractor_Redirect(t *testing.T) {
	mockR := &mockRepo{
		FindByCodeFunc: func(code string) (*entity.ShortURL, error) {
			return &entity.ShortURL{Code: code, OriginalURL: "https://example.com"}, nil
		},
	}
	mockP := &mockPresenter{}
	uc := usecase.NewInteractor(mockR, mockP)

	input := usecase.RedirectInput{Code: "abc123"}
	output, err := uc.Redirect(input)

	assert.NoError(t, err)
	assert.Equal(t, "https://example.com", output.OriginalURL)
}
