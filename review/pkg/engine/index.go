package engine

import "goondex/goondex"

// Позволяет добавлять документы в индекс и искать по индексу
type Index interface {
	Add(page goondex.Page)
	Search(query string) []int
}
