package engine

import (
	"goondex/pkg/spider"
	"strings"
)

// Поисковый движок
type Engine struct {
	scanner spider.Interface
	links   map[string]string
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
func (eng *Engine) Search(query string) map[string]string {
	result := make(map[string]string)
	for k, v := range eng.links {
		if strings.Contains(k, query) || strings.Contains(v, query) {
			result[k] = v
		}
	}
	return result
}
