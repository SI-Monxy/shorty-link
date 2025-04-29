package shorturl

import (
	"fmt"
	"math/rand"
	domain "shorty-link/internal/domain/shorturl"
	"time"
)

// Interactor（ビジネスロジック実装）
type Interactor struct {
	repo domain.Repository
	pres OutputPort
}

func NewInteractor(r domain.Repository, p OutputPort) *Interactor {
	return &Interactor{repo: r, pres: p}
}

func (i *Interactor) Execute(input ShortenInput) (*ShortenOutput, error) {
	code := generateCode(6)
	entity := &domain.ShortURL{
		Code:        code,
		OriginalURL: input.OriginalURL,
	}
	if err := i.repo.Save(entity); err != nil {
		return nil, err
	}
	output := &ShortenOutput{ShortURL: fmt.Sprintf("http://localhost:8080/%s", code)}
	return i.pres.Present(output), nil
}

func generateCode(length int) string {
	rand.Seed(time.Now().UnixNano())
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
