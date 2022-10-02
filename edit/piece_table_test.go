package edit

import (
	"testing"
)

func TestInsert(t *testing.T) {
	pt := NewPieceTable("test")
	tests := []struct{
		i int
		b byte
		want string
	}{
		{0, 'a', "atest"},
		{4, 'b', "atestb"},
		{3, 'c', "atecstb"},
	}
  for _, test := range tests {
    pt.Insert(test.b, test.i)
    want, got := test.want, pt.String()
    if want != got {
      t.Errorf("want: %s, got: %s", want, got)
    }
  }
}

func TestDelete(t *testing.T) {
	pt := NewPieceTable("testing")

	tests := []struct{
		i int
		want string
	}{
		{1, "tsting"},
		{3, "tstng"},
		{2, "tsng"},
		{1, "tng"},
		{1, "tg"},
		{-1, "tg"},
		{2, "tg"},
		{0, "g"},
		{0, ""},
		{0, ""},
	}

	for _, test := range tests {
		pt.Delete(test.i)
		want, got := test.want, pt.String()
		if want != got {
			t.Errorf("want: %s, got: %s", want, got)
		}
	}
}
