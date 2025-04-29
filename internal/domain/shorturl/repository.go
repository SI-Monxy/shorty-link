package shorturl

// Repositoryインターフェース
type Repository interface {
	Save(s *ShortURL) error
	FindByCode(code string) (*ShortURL, error)
}
