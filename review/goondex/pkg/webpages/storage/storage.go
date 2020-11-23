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

func compareFn(nodeVal interface{}, newVal interface{}) int {
	pageNode := nodeVal.(webpages.Page)
	pageNew := newVal.(webpages.Page)
	var res int
	switch true {
	case pageNode.ID < pageNew.ID:
		res = bitree.Right
	case pageNode.ID == pageNew.ID:
		res = bitree.Equal
	case pageNode.ID > pageNew.ID:
		res = bitree.Left
	}
	return res
}
