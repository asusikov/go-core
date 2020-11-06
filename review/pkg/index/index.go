package index

import "goondex/web"

type Index struct{}

func (ind *Index) Add(page web.Page) {

}

func (ind *Index) Search(query string) []int {
	return []int{}
}

func New() *Index {
	return &Index{}
}
