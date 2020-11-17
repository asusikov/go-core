package storage

import (
	"goondex/webpages"
	"reflect"
	"testing"
)

func TestInsert(t *testing.T) {
	storage := Storage{
		pages: []webpages.Page{},
	}
	page1 := webpages.Page{ID: 1}
	page2 := webpages.Page{ID: 2}
	page3 := webpages.Page{ID: 3}
	storage.Insert(page2)
	storage.Insert(page1)
	storage.Insert(page3)
	want := []webpages.Page{
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
		pages: []webpages.Page{
			webpages.Page{ID: 3},
			webpages.Page{ID: 2},
			webpages.Page{ID: 1},
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
