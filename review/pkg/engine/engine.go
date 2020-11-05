package engine

import (
	"goondex"
	"spider"
	"strings"
)

// Поисковый движок
type Engine struct {
	scanner Scanner
	links   []goondex.Page
}

func New() *Engine {
	return &Engine{
		scanner: spider.New(),
	}
}

// Сканировать сайт
func (eng *Engine) Scan(url string) error {
	const depth = 2
	links, err := eng.scanner.Scan(url, depth)
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
