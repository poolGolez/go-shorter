package shorturl

import (
	"crypto/rand"
	"math/big"
)

func CreateShortenedUrl(url string) (*ShortUrl, error) {
	const codeLength = 6
	code, err := generateCode(codeLength)
	if err != nil {
		return nil, err
	}

	shortUrl := &ShortUrl{
		Code:      code,
		TargetUrl: url,
	}
	return shortUrl, nil
}

func generateCode(length int) (string, error) {
	const letters = "abcdefghijklmnopqrstuvwxyz0123456789"
	code := ""
	for i := 0; i < length; i++ {
		index, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", err
		}
		code = code + string(letters[index.Int64()])
	}

	return code, nil
}
