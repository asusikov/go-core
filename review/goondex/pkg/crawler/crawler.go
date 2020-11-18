// Package crawler реализует сканер содержимого веб-сайтов.
// Пакет позволяет получить информацию об страницах на сайте и сохранить в индекс и хранилище.
package crawler

import (
	"goondex/crawler/site"
	"goondex/engine/storage"
	"goondex/index"
)

// Scan осуществляет рекурсивный обход ссылок сайта, указанного в URL,
// с учётом глубины перехода по ссылкам, переданной в depth.
func Scan(url string, depth int, storage storage.Interface, index index.Interface) error {
	pages, err := site.Scan(url, depth)
	if err != nil {
		return err
	}
	for _, page := range pages {
		storage.Insert(page)
		index.Add(page)
	}
	return nil
}
