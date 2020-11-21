// Реализует бинарное дерево
package bitree

import (
	"errors"
	"fmt"
)

const (
	Left = iota + 1
	Equal
	Right
)

// Элемент в бинарном дереве
type TreeNode struct {
	Value interface{}
	Left  *TreeNode
	Right *TreeNode
}

type compareFnType = func(interface{}, interface{}) int

var ErrNotFound = errors.New("Value not found")

// Добавить элемент в дерево
func Add(node *TreeNode, val interface{}, compareFn compareFnType) (el *TreeNode) {
	initFn := func(value interface{}) *TreeNode {
		return &TreeNode{
			Value: value,
		}
	}
	switch compareFn(node.Value, val) {
	case Left:
		if node.Left != nil {
			el = Add(node.Left, val, compareFn)
		} else {
			el = initFn(val)
			node.Left = el
		}
	case Right:
		if node.Right != nil {
			el = Add(node.Right, val, compareFn)
		} else {
			el = initFn(val)
			node.Right = el
		}
	case Equal:
		node.Value = val
		el = node
	}
	return el
}

func Search(node *TreeNode, val interface{}, compareFn compareFnType) (el *TreeNode, err error) {
	if node == nil {
		return nil, fmt.Errorf("%w: %v", ErrNotFound, val)
	}
	switch compareFn(node.Value, val) {
	case Left:
		el, err = Search(node.Left, val, compareFn)
	case Right:
		el, err = Search(node.Right, val, compareFn)
	case Equal:
		el = node
	}
	return el, err
}
