package engine

import (
	"goondex/crawler"
	"goondex/goondex"
	"strings"
)

// Поисковый движок
type Engine struct {
	crawler crawler.Interface
	links   []goondex.Page
}

func New() *Engine {
	return &Engine{
		crawler: crawler.New(),
	}
}

// Сканировать сайт
func (eng *Engine) Scan(url string) error {
	const depth = 2
	links, err := eng.crawler.Scan(url, depth)
	if err != nil {
		return err
	}
	eng.links = links
	return nil
}

// Поиск ссылки по слову
func (eng *Engine) Search(query string) []goondex.Page {
	result := []goondex.Page{}
	for _, page := range eng.links {
		if strings.Contains(page.Title, query) {
			result = append(result, page)
		}
	}
	return result
}
