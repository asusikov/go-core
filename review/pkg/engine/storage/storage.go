package storage

import "goondex/web"

type Interface interface {
	Add(page web.Page)
	Find(id int) web.Page
}

type Storage struct {
	pages map[int]web.Page
}

func (st *Storage) Add(page web.Page) {
	st.pages[page.Id] = page
}

func (st *Storage) Find(id int) web.Page {
	return st.pages[id]
}

func New() *Storage {
	return &Storage{
		pages: make(map[int]web.Page),
	}
}
