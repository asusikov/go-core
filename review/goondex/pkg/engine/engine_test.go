package engine

import (
	"errors"
	"goondex/web"
	"testing"
)

type StubStorage struct{}

func (st *StubStorage) Find(id int) (*web.Page, error) {
	if id == 1 {
		return &web.Page{ID: 1, URL: "yandex.ru", Title: "Яндекс"}, nil
	} else {
		return nil, errors.New("not found")
	}
}
func (st *StubStorage) Insert(page web.Page) {}

type StubIndex struct{}

func (si *StubIndex) Search(query string) []int {
	if query == "Яндекс" {
		return []int{1}
	} else {
		return []int{}
	}
}
func (si *StubIndex) Add(web.Page) {}

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
