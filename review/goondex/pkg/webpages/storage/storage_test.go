package storage

import (
	"fmt"
	"goondex/bitree"
	"goondex/webpages"
	"testing"
)

func TestInsert(t *testing.T) {
	initRoot := func() *bitree.TreeNode {
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

// func TestInsert(t *testing.T) {
// 	storage := Storage{
// 		pages: []webpages.Page{},
// 	}
// 	page1 := webpages.Page{ID: 1}

// 	page2 := webpages.Page{ID: 2}
// 	page3 := webpages.Page{ID: 3}
// 	storage.Insert(page2)
// 	storage.Insert(page1)
// 	storage.Insert(page3)
// 	want := []webpages.Page{
// 		page3,
// 		page2,
// 		page1,
// 	}
// 	got := storage.pages
// 	if !reflect.DeepEqual(want, got) {
// 		t.Fatalf("ждали %v, получили %v", want, got)
// 	}
// }

// func TestFind(t *testing.T) {
// 	storage := Storage{
// 		pages: []webpages.Page{
// 			webpages.Page{ID: 3},
// 			webpages.Page{ID: 2},
// 			webpages.Page{ID: 1},
// 		},
// 	}
// 	want := &storage.pages[1]
// 	got, err := storage.Find(2)
// 	if err != nil {
// 		t.Fatalf("поиск завершился с ошибкой")
// 	}
// 	if !reflect.DeepEqual(want, got) {
// 		t.Fatalf("ждали %v, получили %v", want, got)
// 	}
// }
