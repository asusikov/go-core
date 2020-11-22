package bitree

import (
	"fmt"
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
		name  string
		value int
		want  string
	}{
		{
			name:  "Добавление правой ветви к корню",
			value: 7,
			want:  " 5 3 7",
		},
		{
			name:  "Добавление правой ветви к ветви",
			value: 4,
			want:  " 5 3 4 7",
		},
		{
			name:  "Добавление левой ветви к ветви",
			value: 2,
			want:  " 5 3 2 4 7",
		},
		{
			name:  "Добавление существующего значения",
			value: 3,
			want:  " 5 3 2 4 7",
		},
	}

	serializeValueFn := func(value interface{}) string {
		return fmt.Sprintf(" %v", value.(int))
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
			got := Serialize(root, serializeValueFn)
			if test.want != got {
				t.Errorf("ожидали %s, получили %s", test.want, got)
			}
		})
	}
}

func TestSerialize(t *testing.T) {
	root := &TreeNode{
		Value: "d",
		Left: &TreeNode{
			Value: "b",
			Right: &TreeNode{
				Value: "c",
			},
		},
		Right: &TreeNode{
			Value: "e",
		},
	}
	serializeValueFn := func(value interface{}) string {
		return " " + value.(string)
	}
	want := " d b c e"
	got := Serialize(root, serializeValueFn)
	if want != got {
		t.Errorf("ожидали %s, получили %s", want, got)
	}
}
