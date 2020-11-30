package webpages

import "hash/fnv"

// Описание страницы
type Page struct {
	URL   string
	Title string
	ID    uint32
}

func New(title string, url string) *Page {
	fnvHash := fnv.New32a()
	_, _ = fnvHash.Write([]byte(title))
	return &Page{
		ID:    fnvHash.Sum32(),
		Title: title,
		URL:   url,
	}
}
