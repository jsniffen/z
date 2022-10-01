package edit

import "testing"

func TestGapBufferInsert(t *testing.T) {
	gb := NewGapBuffer()

	tests := []struct{
		b byte
		want string
	}{
		{'a', "a"},
		{'b', "ab"},
		{'c', "abc"},
		{'d', "abcd"},
	}

	for _, test := range tests {
		gb.Insert(test.b)
		want, got := test.want, gb.String()
		if want != got {
			t.Errorf("want: %s, got: %s", want, got)
		}
	}
}

func TestGapBufferDelete(t *testing.T) {
	w := "testing"

	gb := NewGapBuffer()
	for _, c := range w {
		gb.Insert(byte(c))
	}

	for i := 0; i < len(w); i += 1 {
		gb.Delete()
		want, got := w[:len(w)-i-1], gb.String()
		if want != got {
			t.Errorf("want: %s, got: %s", want, got)
		}
	}
}
