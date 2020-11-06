package index

import "goondex/web"

// Позволяет добавлять документы в индекс и искать по индексу
type Interface interface {
	Add(page web.Page)
	Search(query string) []int
}
