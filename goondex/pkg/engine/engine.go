package engine

import (
	"strings"
)

// Engine - поисковый движок
type Engine struct {
	links map[string]string
}

func New() *Engine {
	return &Engine{}
}

// Index - индексировать хэш с ссылками
func (eng *Engine) Index(links map[string]string) {
	eng.links = links
}

// Search - поиск ссылок
func (eng *Engine) Search(query string) map[string]string {
	result := make(map[string]string)
	for k, v := range eng.links {
		if strings.Contains(k, query) || strings.Contains(v, query) {
			result[k] = v
		}
	}
	return result
}
