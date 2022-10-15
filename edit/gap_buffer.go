package edit

type GapBuffer struct {
	buffer []rune
	size   int
	index  int
	cx     int
	cy     int
}

func NewGapBuffer(s string) *GapBuffer {
	b := make([]rune, 256)
	for i, c := range s {
		b[i] = c
	}

	return &GapBuffer{
		buffer: b,
		size:   256 - len(s),
		index:  len(s),
		cx:     0,
		cy:     0,
	}
}

func (gb *GapBuffer) String() string {
	return string(gb.buffer[:gb.index])
}

func (gb *GapBuffer) Insert(rune) {
}
