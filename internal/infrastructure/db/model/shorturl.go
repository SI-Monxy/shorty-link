package model

import "gorm.io/gorm"

// GORMモデル（DB保存用構造体）
type ShortURL struct {
	gorm.Model
	Code        string `gorm:"uniqueIndex;size:10;not null"`
	OriginalURL string `gorm:"type:text;not null"`
}
