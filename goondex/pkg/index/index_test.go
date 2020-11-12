package index

import (
	"goondex/web"
	"reflect"
	"testing"
)

func TestSearch(t *testing.T) {
	ind := New()
	ind.Add(web.Page{ID: 1, URL: "yandex.ru", Title: "Яндекс"})
	ind.Add(web.Page{ID: 2, URL: "google.com", Title: "Google"})
	want := []int{1}
	got := ind.Search("Яндекс")
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("для запроса Яндекс ждали %v, получили %v", want, got)
	}
}
