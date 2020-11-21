package bitree

import (
	"testing"
)

func compareFn(nodeVal interface{}, newVal interface{}) int {
	nodeInt := nodeVal.(int)
	newInt := newVal.(int)
	var res int
	switch true {
	case nodeInt < newInt:
		res = Right
	case nodeInt == newInt:
		res = Equal
	case nodeInt > newInt:
		res = Left
	}
	return res
}

func TestAdd(t *testing.T) {
	root := &TreeNode{
		Value: 5,
		Left: &TreeNode{
			Value: 3,
		},
	}

	tests := []struct {
		name   string
		value  int
		parent *TreeNode
		branch int
	}{
		{
			name:   "Добавление правой ветви к корню",
			value:  7,
			parent: root,
			branch: Right,
		},
		{
			name:   "Добавление правой ветви к ветви",
			value:  4,
			parent: root.Left,
			branch: Right,
		},
		{
			name:   "Добавление левой ветви к ветви",
			value:  2,
			parent: root.Left,
			branch: Left,
		},
		{
			name:   "Добавление существующего значения",
			value:  3,
			parent: root,
			branch: Left,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			el := Add(root, test.value, compareFn)
			if el == nil {
				t.Errorf("Элемент не создан")
			}
			if el.Value != test.value {
				t.Errorf("Значение должно было %v, получили %v", test.value, el.Value)
			}
			var node *TreeNode
			if test.branch == Left {
				node = test.parent.Left
			} else {
				node = test.parent.Right
			}
			if test.value != node.Value {
				t.Errorf("Значение вставлено не верно")
			}
		})
	}
}
