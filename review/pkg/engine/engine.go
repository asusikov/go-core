package engine

import (
	"goondex/crawler"
	"goondex/engine/storage"
	"goondex/index"
	"goondex/web"
)

// Поисковый движок
type Engine struct {
	crawler crawler.Interface
	index   index.Interface
	storage storage.Interface
}

// Сканировать сайт
func (eng *Engine) Scan(url string) error {
	const depth = 2
	pages, err := eng.crawler.Scan(url, depth)
	if err != nil {
		return err
	}
	for _, page := range pages {
		eng.storage.Insert(page)
		eng.index.Add(page)
	}
	return nil
}

// Поиск ссылки по слову
func (eng *Engine) Search(query string) ([]web.Page, error) {
	result := []web.Page{}
	for _, id := range eng.index.Search(query) {
		page, err := eng.storage.Find(id)
		if err != nil {
			return nil, err
		}
		result = append(result, *page)
	}
	return result, nil
}

func New() *Engine {
	return &Engine{
		crawler: crawler.New(),
		index:   index.New(),
		storage: storage.New(),
	}
}
