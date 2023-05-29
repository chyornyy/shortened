package models

type Link struct {
	ShortUrl    string `json:"shorturl"`
	OriginalUrl string `json:"originalurl"`
}
