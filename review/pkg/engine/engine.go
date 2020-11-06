package engine

import (
	"goondex/crawler"
	"goondex/index"
	"goondex/web"
)

// Поисковый движок
type Engine struct {
	crawler crawler.Interface
	index   index.Interface
	pages   map[int]web.Page
}

// Сканировать сайт
func (eng *Engine) Scan(url string) error {
	const depth = 2
	pages, err := eng.crawler.Scan(url, depth)
	if err != nil {
		return err
	}
	for _, page := range pages {
		eng.pages[page.Id] = page
		eng.index.Add(page)
	}
	return nil
}

// Поиск ссылки по слову
func (eng *Engine) Search(query string) []web.Page {
	result := []web.Page{}
	for _, id := range eng.index.Search(query) {
		result = append(result, eng.pages[id])
	}
	return result
}

func New() *Engine {
	return &Engine{
		crawler: crawler.New(),
		index:   index.New(),
		pages:   make(map[int]web.Page),
	}
}
