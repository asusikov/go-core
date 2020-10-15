package fibonacci

import (
	"reflect"
	"testing"
)

func Test_Calculate_whenInputNumberIsRight(t *testing.T) {
	want := []int{1, 1, 2, 3, 5}
	got := Calculate(5)
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("получили %d, ожидали %d", got, want)
	}
}

func Test_Calculate_whenInputNumberIs1(t *testing.T) {
	want := []int{1}
	got := Calculate(1)
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("получили %d, ожидали %d", got, want)
	}
}

func Test_Calculate_whenInputNumberIs2(t *testing.T) {
	want := []int{1, 1}
	got := Calculate(2)
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("получили %d, ожидали %d", got, want)
	}
}

func Test_Calculate_whenInputNumberIsLess0(t *testing.T) {
	want := []int{}
	got := Calculate(-1)
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("получили %d, ожидали %d", got, want)
	}
}
