package shorturl

var records = make(map[string]ShortUrl)

func Save(shortUrl *ShortUrl) {
	records[shortUrl.Code] = *shortUrl
}

func FindByCode(code string) ShortUrl {
	return records[code]
}
