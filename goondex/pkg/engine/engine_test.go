package engine

import (
	"testing"
)

func Test_Search(t *testing.T) {
	eng := Engine{}
	links := map[string]string{
		"yandex.ru":  "Яндекс",
		"google.com": "Google",
	}
	eng.Index(links)
	got := eng.Search("yandex")
	if len(got) != 1 {
		t.Fatalf("длина хэша не равна 1")
	}
	if got["yandex.ru"] != links["yandex.ru"] {
		t.Fatalf("для ключа yandex.ru ждали %s, получили %s", links["yandex.ru"], got["yandex.ru"])
	}
}
