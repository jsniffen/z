package edit

import "strings"

type GapBuffer struct {
	size int 
	gap int 
	gapsize int 
	data []byte
}

func NewGapBuffer() *GapBuffer {
	return &GapBuffer{
		size: 256,
		gap: 0,
		gapsize: 256,
		data: make([]byte, 256),
	}
}

func (gb *GapBuffer) String() string {
	var sb strings.Builder
	sb.Write(gb.data[:gb.gap])
	sb.Write(gb.data[gb.gap+gb.gapsize:])
	return sb.String()
}

func (gb* GapBuffer) Insert(b byte) {
	gb.data[gb.gap] = b
	gb.gapsize -= 1
	gb.gap += 1
}

func (gb* GapBuffer) Delete() {
	gb.gap -=1 
	gb.gapsize += 1
}
