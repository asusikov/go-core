package index

import (
	"goondex/webpages"
	"strings"
)

type Index struct {
	index map[string][]uint32
}

func (ind *Index) Add(page webpages.Page) {
	tokens := extractTokens(page.Title)
	for _, token := range tokens {
		lowerToken := strings.ToLower(token)
		ind.index[lowerToken] = append(ind.index[lowerToken], page.ID)
	}
}

func (ind *Index) Search(query string) []uint32 {
	return ind.index[strings.ToLower(query)]
}

func New() *Index {
	return &Index{
		index: make(map[string][]uint32),
	}
}

func extractTokens(word string) []string {
	return strings.Split(word, " ")
}
