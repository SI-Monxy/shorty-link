package shorturl

import "errors"

// リクエストバリデーション専用クラス
func ValidateInput(input ShortenInput) error {
	if input.OriginalURL == "" {
		return errors.New("original_url is required")
	}
	return nil
}
