// Реализует бинарное дерево
package bitree

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

// Добавить элемент в дерево
func Add(node *TreeNode, val interface{}, compareFn func(interface{}, interface{}) int) (el *TreeNode) {
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
