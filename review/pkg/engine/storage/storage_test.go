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
	page1 := web.Page{Id: 1}
	page2 := web.Page{Id: 2}
	page3 := web.Page{Id: 3}
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
			web.Page{Id: 3},
			web.Page{Id: 2},
			web.Page{Id: 1},
		},
	}
	want := storage.pages[1]
	got := storage.Find(2)
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("ждали %v, получили %v", want, got)
	}
}
