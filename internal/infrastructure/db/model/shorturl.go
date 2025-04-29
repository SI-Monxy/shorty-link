package model

// GORMモデル（DB保存用構造体）
type ShortURL struct {
	ID          uint   `gorm:"primaryKey"`
	Code        string `gorm:"uniqueIndex;size:10"`
	OriginalURL string `gorm:"type:text"`
}
