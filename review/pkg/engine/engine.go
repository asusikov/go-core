package engine

import (
	"goondex/crawler"
	"goondex/index"
	"goondex/web"
	"strings"
)

// Поисковый движок
type Engine struct {
	crawler crawler.Interface
	index   index.Interface
	links   []web.Page
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
func (eng *Engine) Search(query string) []web.Page {
	result := []web.Page{}
	for _, page := range eng.links {
		if strings.Contains(page.Title, query) {
			result = append(result, page)
		}
	}
	return result
}

func New() *Engine {
	return &Engine{
		crawler: crawler.New(),
		index:   index.New(),
	}
}
