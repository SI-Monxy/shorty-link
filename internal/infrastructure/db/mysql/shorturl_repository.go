package mysql

import (
	"shorty-link/internal/domain/shorturl"
	"shorty-link/internal/infrastructure/db/model"

	"gorm.io/gorm"
)

// MySQL用リポジトリ実装
type ShortURLRepository struct {
	db *gorm.DB
}

func NewShortURLRepository(db *gorm.DB) *ShortURLRepository {
	return &ShortURLRepository{db: db}
}

func (r *ShortURLRepository) Save(s *shorturl.ShortURL) error {
	return r.db.Create(&model.ShortURL{Code: s.Code, OriginalURL: s.OriginalURL}).Error
}

func (r *ShortURLRepository) FindByCode(code string) (*shorturl.ShortURL, error) {
	var m model.ShortURL
	if err := r.db.Where("code = ?", code).First(&m).Error; err != nil {
		return nil, err
	}
	return &shorturl.ShortURL{
		Code:        m.Code,
		OriginalURL: m.OriginalURL,
	}, nil
}
