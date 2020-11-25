package list

import (
	"testing"
)

func TestNew(t *testing.T) {
	l := New()
	t.Logf("%+v", l)
}

func TestList_Push(t *testing.T) {
	l := New()
	l.Push(Elem{Val: 2})
	l.Push(Elem{Val: 1})
	got := l.String()
	want := "1 2"
	if got != want {
		t.Fatalf("получили %s, ожидалось %s", got, want)
	}
}

func TestLis_Pop(t *testing.T) {
	tests := []struct {
		name     string
		elements []Elem
		want     string
	}{
		{
			name: "Удаление из длинного списка",
			elements: []Elem{
				Elem{Val: 3},
				Elem{Val: 2},
				Elem{Val: 1},
			},
			want: "2 3",
		},
		{
			name: "Удаление из списка из одного элемента",
			elements: []Elem{
				Elem{Val: 1},
			},
			want: "",
		},
		{
			name:     "Удаление из пустого списка",
			elements: []Elem{},
			want:     "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			l := New()
			for _, el := range test.elements {
				l.Push(el)
			}
			got := l.Pop().String()
			if got != test.want {
				t.Fatalf("получили %s, ожидалось %s", got, test.want)
			}
		})
	}
}

func TestLis_Reverse(t *testing.T) {
	tests := []struct {
		name     string
		elements []Elem
		want     string
	}{
		{
			name: "Разворот длинного списка",
			elements: []Elem{
				Elem{Val: 3},
				Elem{Val: 2},
				Elem{Val: 1},
			},
			want: "3 2 1",
		},
		{
			name: "Разворот списка из одного элемента",
			elements: []Elem{
				Elem{Val: 1},
			},
			want: "1",
		},
		{
			name:     "Разворот пустого списка",
			elements: []Elem{},
			want:     "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			l := New()
			for _, el := range test.elements {
				l.Push(el)
			}
			got := l.Reverse().String()
			if got != test.want {
				t.Fatalf("получили %s, ожидалось %s", got, test.want)
			}
		})
	}
}
