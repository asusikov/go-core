package storage

import (
	"goondex/bitree"
	"goondex/webpages"
)

type Interface interface {
	Insert(page webpages.Page)
	Find(id uint32) (webpages.Page, error)
}

type Storage struct {
	root *bitree.TreeNode
}

func (st *Storage) Insert(page webpages.Page) {
	if st.root == nil {
		st.root = &bitree.TreeNode{
			Value: page,
		}
	} else {
		_ = bitree.Add(st.root, page, compareFn)
	}
}

func (st *Storage) Find(id uint32) (page webpages.Page, err error) {
	p := webpages.Page{ID: id}
	node, err := bitree.Search(st.root, p, compareFn)
	if err != nil {
		return page, err
	}
	page = node.Value.(webpages.Page)
	return page, nil
}

func New() *Storage {
	return &Storage{
		root: nil,
	}
}

func compareFn(a interface{}, b interface{}) int {
	apage := a.(webpages.Page)
	bpage := b.(webpages.Page)
	var res int
	switch true {
	case apage.ID < bpage.ID:
		res = bitree.Left
	case apage.ID == bpage.ID:
		res = bitree.Equal
	case apage.ID > bpage.ID:
		res = bitree.Right
	}
	return res
}
