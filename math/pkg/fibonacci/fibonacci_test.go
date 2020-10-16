package fibonacci

import (
	"reflect"
	"testing"
)

func Test_CalculateNumber(t *testing.T) {
	want := 13
	got := CalculateNumber(7)
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("получили %d, ожидали %d", got, want)
	}
}
