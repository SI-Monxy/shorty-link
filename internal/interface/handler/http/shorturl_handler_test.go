package http

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	entity "shorty-link/internal/domain/shorturl"
	usecase "shorty-link/internal/usecase/shorturl"
)

// ===== モック定義 =====

// モックリポジトリ（ドメイン層のShortURLを扱う）
type mockRepo struct{}

func (m *mockRepo) Save(_ *entity.ShortURL) error {
	return nil
}

func (m *mockRepo) FindByCode(code string) (*entity.ShortURL, error) {
	return &entity.ShortURL{
		Code:        code,
		OriginalURL: "https://example.com",
	}, nil
}

// モックプレゼンター（DTOをそのまま返す）
type mockPresenter struct{}

func (m *mockPresenter) PresentShorten(output *usecase.ShortenOutput) *usecase.ShortenOutput {
	return output
}

func (m *mockPresenter) PresentRedirect(output *usecase.RedirectOutput) *usecase.RedirectOutput {
	return output
}

// ===== テスト関数 =====

// POST /api/v1/shorten のテスト
func TestShortenAPI(t *testing.T) {
	gin.SetMode(gin.TestMode)

	repo := &mockRepo{}
	pres := &mockPresenter{}
	uc := usecase.NewInteractor(repo, pres, "http://localhost:8080")
	h := NewShortURLHandler(uc)

	r := gin.Default()
	h.RegisterRoutes(r)

	body := `{"original_url": "https://example.com"}`
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/shorten", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "short_url")
}

// GET /:code のテスト
func TestRedirectAPI(t *testing.T) {
	gin.SetMode(gin.TestMode)

	repo := &mockRepo{}
	pres := &mockPresenter{}
	uc := usecase.NewInteractor(repo, pres, "http://localhost:8080")
	h := NewShortURLHandler(uc)

	r := gin.Default()
	h.RegisterRoutes(r)

	req, _ := http.NewRequest(http.MethodGet, "/abc123", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusFound, w.Code)
	assert.Equal(t, "https://example.com", w.Header().Get("Location"))
}
