package storage

import (
	"goondex/web"
	"sort"
)

type Interface interface {
	Insert(page web.Page)
	Find(id int) web.Page
}

type Storage struct {
	pages []web.Page
}

func (st *Storage) Insert(page web.Page) {
	pos := findPos(st.pages, page.ID)
	st.pages = append(st.pages, web.Page{})
	copy(st.pages[pos+1:], st.pages[pos:])
	st.pages[pos] = page
}

func (st *Storage) Find(id int) web.Page {
	pos := findPos(st.pages, id)
	return st.pages[pos]
}

func New() *Storage {
	return &Storage{
		pages: []web.Page{},
	}
}

func findPos(pages []web.Page, id int) int {
	return sort.Search(len(pages), func(i int) bool { return pages[i].ID <= id })
}
