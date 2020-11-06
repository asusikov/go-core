package index

import (
	"goondex/web"
	"strings"
)

type Index struct {
	tokens map[string][]int
}

func (ind *Index) Add(page web.Page) {
	tokens := extractTokens(page.Title)
	for _, token := range tokens {
		ind.tokens[token] = append(ind.tokens[token], page.Id)
	}
}

func (ind *Index) Search(query string) []int {
	return ind.tokens[query]
}

func New() *Index {
	return &Index{
		tokens: make(map[string][]int),
	}
}

func extractTokens(word string) []string {
	return strings.Split(word, " ")
}
