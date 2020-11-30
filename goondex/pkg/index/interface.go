package index

import "goondex/webpages"

// Позволяет добавлять документы в индекс и искать по индексу
type Interface interface {
	Add(page webpages.Page)
	Search(query string) []uint32
}
