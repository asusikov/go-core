package index

import (
	"goondex/web"
	"reflect"
	"testing"
)

func TestSearch(t *testing.T) {
	ind := Index{}
	ind.Add(web.Page{Id: 1, Url: "yandex.ru", Title: "Яндекс"})
	ind.Add(web.Page{Id: 2, Url: "google.com", Title: "Google"})
	want := []int{1}
	got := ind.Search("Яндекс")
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("для запроса Яндекс ждали %v, получили %v", want, got)
	}
}
