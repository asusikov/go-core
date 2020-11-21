package engine

import (
	"goondex/index"
	"goondex/webpages"
	"goondex/webpages/storage"
)

// Поисковый движок
type Engine struct {
	index   index.Interface
	storage storage.Interface
}

// Поиск ссылки по слову
func (eng *Engine) Search(query string) ([]webpages.Page, error) {
	result := []webpages.Page{}
	for _, id := range eng.index.Search(query) {
		page, err := eng.storage.Find(id)
		if err != nil {
			return nil, err
		}
		result = append(result, page)
	}
	return result, nil
}

func New(index index.Interface, storage storage.Interface) *Engine {
	return &Engine{
		index:   index,
		storage: storage,
	}
}
