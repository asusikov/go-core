package fibonacci

import (
	"testing"
	"reflect"
)

func Test_Calculate_whenInputNumberIsRight(t *testing.T) {
	want := []int{1,1,2,3,5}
	got, err := Calculate(5)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("получили %d, ожидали %d", got, want)
	}
}

func Test_Calculate_whenInputNumberIsWrong(t *testing.T) {
	_, err := Calculate(21)
	if err == nil {
		t.Fatalf("выполение завершилось без ошибки")
	}
}
