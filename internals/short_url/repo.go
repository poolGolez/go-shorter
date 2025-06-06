package shorturl

import (
	"example.com/go-shorter/internals/gorm"
)

var records = make(map[string]ShortUrl)

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

func FindByCode(code string) ShortUrl {
	return records[code]
}
