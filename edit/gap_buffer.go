package edit

import "fmt"

type GapBuffer struct {
	buffer []rune
	size int
	index int
}

func NewGapBuffer(s string) *GapBuffer {
	b := make([]rune, 256)
	for i, c := range s {
		b[i] = c
	}

	return &GapBuffer{
		buffer: b,
		size: 256-len(s),
		index: len(s),
	}
}

func (gb *GapBuffer) Insert(rune) {
}

func (gb *GapBuffer) Render() {
	a := string(gb.buffer[:gb.index])
	fmt.Println(a)
}

