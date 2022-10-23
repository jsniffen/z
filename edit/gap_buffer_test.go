package edit

import (
	"testing"
)

func TestGapBufferInsert(t *testing.T) {
	gb := NewGapBuffer()

	tests := []struct{
		insert rune
		result string
	}{
		{'h', "h",},
		{'e', "he",},
		{'l', "hel",},
		{'l', "hell",},
		{'o', "hello",},
		{'!', "hello!",},
	}

	gapstart := 0
	gapsize := 256
	for _, test := range tests {
		gb.Insert(test.insert)
		gapstart += 1
		gapsize -= 1

		want, got := test.result, gb.String()
		if want != got {
			t.Errorf("want: %s, got: %s", want, got)
		}

		if gapstart != gb.gapstart {
			t.Errorf("want: %d, got: %d", gapstart, gb.gapstart)
		}

		if gapsize != gb.gapsize {
			t.Errorf("want: %d, got: %d", gapsize, gb.gapsize)
		}
	}
}

func TestGapBufferDelete(t *testing.T) {
	gb := NewGapBuffer()

	gb.Insert('h')
	gb.Insert('e')
	gb.Insert('l')
	gb.Insert('l')
	gb.Insert('o')
	gb.Insert('!')

	tests := []string{
		"hello",
		"hell",
		"hel",
		"he",
		"h",
		"",
	}

	gapstart := 6
	gapsize := 250
	for _, test := range tests {
		gb.Delete()
		gapstart -= 1
		gapsize += 1

		want, got := test, gb.String()
		if want != got {
			t.Errorf("want: %s, got: %s", want, got)
		}

		if gapstart != gb.gapstart {
			t.Errorf("want: %d, got: %d", gapstart, gb.gapstart)
		}

		if gapsize != gb.gapsize {
			t.Errorf("want: %d, got: %d", gapsize, gb.gapsize)
		}
	}
}
