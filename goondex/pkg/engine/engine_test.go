package engine

import (
	"testing"
)

type StubScanner struct {
}

func (st *StubScanner) Scan(string, int) (map[string]string, error) {
	links := map[string]string{
		"yandex.ru":  "Яндекс",
		"google.com": "Google",
	}
	return links, nil
}

func Test_Search(t *testing.T) {
	eng := Engine{
		scanner: &StubScanner{},
	}
	eng.Scan("example.com")
	want := "Яндекс"
	got := eng.Search("yandex")
	if len(got) != 1 {
		t.Fatalf("длина хэша не равна 1")
	}
	if got["yandex.ru"] != want {
		t.Fatalf("для ключа yandex.ru ждали %s, получили %s", want, got["yandex.ru"])
	}
}
