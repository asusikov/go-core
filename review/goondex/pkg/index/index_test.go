package index

import (
	"goondex/webpages"
	"reflect"
	"testing"
)

func TestSearch(t *testing.T) {
	ind := New()
	ind.Add(webpages.Page{ID: 1, URL: "yandex.ru", Title: "Яндекс"})
	ind.Add(webpages.Page{ID: 2, URL: "google.com", Title: "Google"})
	want := []uint32{1}
	got := ind.Search("Яндекс")
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("для запроса Яндекс ждали %v, получили %v", want, got)
	}
}
