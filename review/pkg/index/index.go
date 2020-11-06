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
		lowerToken := strings.ToLower(token)
		ind.tokens[lowerToken] = append(ind.tokens[lowerToken], page.Id)
	}
}

func (ind *Index) Search(query string) []int {
	return ind.tokens[strings.ToLower(query)]
}

func New() *Index {
	return &Index{
		tokens: make(map[string][]int),
	}
}

func extractTokens(word string) []string {
	return strings.Split(word, " ")
}
