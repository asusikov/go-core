package storage

import (
	"fmt"
	"goondex/bitree"
	"goondex/webpages"
	"testing"
)

func initRoot() *bitree.TreeNode {
	return &bitree.TreeNode{
		Value: webpages.Page{ID: 5},
		Left: &bitree.TreeNode{
			Value: webpages.Page{ID: 3},
		},
		Right: &bitree.TreeNode{
			Value: webpages.Page{ID: 6},
		},
	}
}

func TestInsert(t *testing.T) {
	tests := []struct {
		name  string
		value webpages.Page
		want  string
		root  *bitree.TreeNode
	}{
		{
			name:  "когда дерева не существует",
			value: webpages.Page{ID: 123},
			want:  " 123",
			root:  nil,
		},
		{
			name:  "когда вставляем в правую ветвь к корню",
			value: webpages.Page{ID: 7},
			want:  " 5 3 6 7",
			root:  initRoot(),
		},
		{
			name:  "когда вставляем в правую ветвь к левой ветви",
			value: webpages.Page{ID: 4},
			want:  " 5 3 4 6",
			root:  initRoot(),
		},
		{
			name:  "когда вставляем в левую ветвь к левой ветви",
			value: webpages.Page{ID: 2},
			want:  " 5 3 2 6",
			root:  initRoot(),
		},
	}

	serializeValueFn := func(val interface{}) string {
		page := val.(webpages.Page)
		return fmt.Sprintf(" %v", page.ID)
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			st := &Storage{
				root: test.root,
			}
			st.Insert(test.value)
			got := bitree.Serialize(st.root, serializeValueFn)
			if test.want != got {
				t.Errorf("ожидали %s, получили %s", test.want, got)
			}
		})
	}
}

func TestFind(t *testing.T) {
	storage := Storage{
		root: initRoot(),
	}
	want := webpages.Page{ID: 9}
	storage.root.Right.Right = &bitree.TreeNode{Value: want}
	got, err := storage.Find(11)
	if err == nil {
		t.Fatal("поиск закончился без ошибки")
	}
	got, err = storage.Find(9)
	if err != nil {
		t.Fatalf("поиск закончился с ошибкой %v", err)
	}
	if got != want {
		t.Fatalf("ожидали %v, получили %v", want, got)
	}
}
