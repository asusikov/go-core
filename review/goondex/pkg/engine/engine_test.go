package engine

import (
	"errors"
	"goondex/webpages"
	"testing"
)

type StubStorage struct{}

func (st *StubStorage) Find(id uint32) (*webpages.Page, error) {
	if id == 1 {
		return &webpages.Page{ID: 1, URL: "yandex.ru", Title: "Яндекс"}, nil
	} else {
		return nil, errors.New("not found")
	}
}
func (st *StubStorage) Insert(page webpages.Page) {}

type StubIndex struct{}

func (si *StubIndex) Search(query string) []uint32 {
	if query == "Яндекс" {
		return []uint32{1}
	} else {
		return []uint32{}
	}
}
func (si *StubIndex) Add(webpages.Page) {}

func TestSearch(t *testing.T) {
	eng := Engine{
		index:   &StubIndex{},
		storage: &StubStorage{},
	}
	result, err := eng.Search("Яндекс")
	if err != nil {
		t.Fatalf("поиск вернул ошибку")
	}
	if len(result) != 1 {
		t.Fatalf("длина результата не равна 1")
	}
	want := "Яндекс"
	got := result[0].Title
	if got != want {
		t.Fatalf("для ключа yandex.ru ждали %s, получили %s", want, got)
	}
}
