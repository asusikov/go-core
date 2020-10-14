package factorial

import (
	"testing"
)

func Test_Calculate_whenInputNumberIsRight(t *testing.T) {
	want := 6
	got, err := Calculate(3)
	if err != nil {
		t.Fatal(err)
	}
	if got != want {
		t.Fatalf("получили %d, ожидали %d", got, want)
	}
}

func Test_Calculate_whenInputNumberIsWrong(t *testing.T) {
	_, err := Calculate(0)
	if err == nil {
		t.Fatalf("выполение завершилось без ошибки")
	}
}
