package index

import (
	"goondex/webpages"
	"strings"
)

type Index struct {
	index map[string][]int
}

func (ind *Index) Add(page webpages.Page) {
	tokens := extractTokens(page.Title)
	for _, token := range tokens {
		lowerToken := strings.ToLower(token)
		ind.index[lowerToken] = append(ind.index[lowerToken], page.ID)
	}
}

func (ind *Index) Search(query string) []int {
	return ind.index[strings.ToLower(query)]
}

func New() *Index {
	return &Index{
		index: make(map[string][]int),
	}
}

func extractTokens(word string) []string {
	return strings.Split(word, " ")
}
