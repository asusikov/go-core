package engine

import (
	"goondex/engine/storage"
	"goondex/index"
	"goondex/web"
	"testing"
)

type StubCrawler struct {
}

func (st *StubCrawler) Scan(string, int) ([]web.Page, error) {
	links := []web.Page{
		web.Page{ID: 1, URL: "yandex.ru", Title: "Яндекс"},
		web.Page{ID: 2, URL: "google.com", Title: "Google"},
	}
	return links, nil
}

func TestSearch(t *testing.T) {
	eng := Engine{
		crawler: &StubCrawler{},
		index:   index.New(),
		storage: storage.New(),
	}
	eng.Scan("example.com")
	want := "Яндекс"
	result, err := eng.Search("Яндекс")
	if err != nil {
		t.Fatalf("поиск вернул ошибку")
	}
	if len(result) != 1 {
		t.Fatalf("длина результата не равна 1")
	}
	got := result[0].Title
	if got != want {
		t.Fatalf("для ключа yandex.ru ждали %s, получили %s", want, got)
	}
}
