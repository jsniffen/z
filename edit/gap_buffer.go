package edit

type GapBuffer struct {
	buffer   []rune
	gapsize  int
	gapstart int
}

func NewGapBuffer() *GapBuffer {
	return &GapBuffer{
		buffer:   make([]rune, 256),
		gapsize:  256,
		gapstart: 0,
	}
}

func (gb *GapBuffer) String() string {
	a := gb.buffer[:gb.gapstart]
	b := gb.buffer[gb.gapstart+gb.gapsize:]
	return string(a) + string(b)
}

func (gb *GapBuffer) Resize() {
	a := gb.buffer[:gb.gapstart]
	b := gb.buffer[gb.gapstart+gb.gapsize:]
	buffer := make([]rune, len(gb.buffer)+256)
	copy(buffer, a)
	copy(buffer[:gb.gapsize+256], b)
	gb.gapsize = 256
	gb.buffer = b
}

func (gb *GapBuffer) Insert(r rune) {
	gb.buffer[gb.gapstart] = r
	gb.gapstart += 1
	gb.gapsize -= 1
}

func (gb *GapBuffer) Delete() {
	gb.gapstart -= 1
	gb.gapsize += 1
}

func (gb *GapBuffer) Move(i int) {
	a := gb.buffer[:gb.gapstart]
	b := gb.buffer[gb.gapstart+gb.gapsize:]
	buffer := make([]rune, len(gb.buffer))
	copy(buffer, a)
	copy(buffer[i+gb.gapsize:], b)
	gb.buffer = buffer
	gb.gapstart = i
}
