package engine

import (
	"goondex/goondex"
	"testing"
)

type StubCrawler struct {
}

func (st *StubCrawler) Scan(string, int) ([]goondex.Page, error) {
	links := []goondex.Page{
		goondex.Page{Id: 1, Url: "yandex.ru", Title: "Яндекс"},
		goondex.Page{Id: 2, Url: "google.com", Title: "Google"},
	}
	return links, nil
}

func Test_Search(t *testing.T) {
	eng := Engine{
		crawler: &StubCrawler{},
	}
	eng.Scan("example.com")
	want := "Яндекс"
	result := eng.Search("Яндекс")
	if len(result) != 1 {
		t.Fatalf("длина результата не равна 1")
	}
	got := result[0].Title
	if got != want {
		t.Fatalf("для ключа yandex.ru ждали %s, получили %s", want, got)
	}
}
