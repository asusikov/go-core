package storage

import (
	"goondex/web"
	"reflect"
	"testing"
)

func TestInsert(t *testing.T) {
	storage := Storage{
		pages: []web.Page{},
	}
	page1 := web.Page{ID: 1}
	page2 := web.Page{ID: 2}
	page3 := web.Page{ID: 3}
	storage.Insert(page2)
	storage.Insert(page1)
	storage.Insert(page3)
	want := []web.Page{
		page3,
		page2,
		page1,
	}
	got := storage.pages
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("ждали %v, получили %v", want, got)
	}
}

func TestFind(t *testing.T) {
	storage := Storage{
		pages: []web.Page{
			web.Page{ID: 3},
			web.Page{ID: 2},
			web.Page{ID: 1},
		},
	}
	want := &storage.pages[1]
	got, err := storage.Find(2)
	if err != nil {
		t.Fatalf("поиск завершился с ошибкой")
	}
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("ждали %v, получили %v", want, got)
	}
}
