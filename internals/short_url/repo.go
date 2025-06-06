package shorturl

import (
	"example.com/go-shorter/internals/gorm"
)

func Migrate() {
	gorm.Db.AutoMigrate(&ShortUrl{})
}

func Save(shortUrl *ShortUrl) (*ShortUrl, error) {
	result := gorm.Db.Create(shortUrl)
	if err := result.Error; err != nil {
		return nil, err
	}

	return shortUrl, nil
}

func FindByCode(code string) (*ShortUrl, error) {
	var shortUrl ShortUrl
	result := gorm.Db.Where(&ShortUrl{Code: code}).Find(&shortUrl)
	if err := result.Error; err != nil {
		return nil, err
	}

	return &shortUrl, nil
}
