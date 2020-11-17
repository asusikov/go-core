package storage

import (
	"fmt"
	"goondex/webpages"
	"sort"
)

type Interface interface {
	Insert(page webpages.Page)
	Find(id int) (*webpages.Page, error)
}

type Storage struct {
	pages []webpages.Page
}

func (st *Storage) Insert(page webpages.Page) {
	pos := findPos(st.pages, page.ID)
	st.pages = append(st.pages, webpages.Page{})
	copy(st.pages[pos+1:], st.pages[pos:])
	st.pages[pos] = page
}

func (st *Storage) Find(id int) (*webpages.Page, error) {
	pos := findPos(st.pages, id)
	if pos < len(st.pages) {
		return &st.pages[pos], nil
	} else {
		return nil, fmt.Errorf("страница %d не найдена", id)
	}
}

func New() *Storage {
	return &Storage{
		pages: []webpages.Page{},
	}
}

func findPos(pages []webpages.Page, id int) int {
	return sort.Search(len(pages), func(i int) bool { return pages[i].ID <= id })
}
