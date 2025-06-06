package shorturl

type ShortUrl struct {
	Id        uint64 `gorm:"primaryKey"`
	Code      string
	TargetUrl string
}
