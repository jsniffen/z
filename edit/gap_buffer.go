package edit

const InitialSize = 256 

type GapBuffer struct {
	buffer   []rune
	gapsize  int
	gapstart int
}

func NewGapBuffer() *GapBuffer {
	return &GapBuffer{
		buffer:   make([]rune, InitialSize),
		gapsize:  InitialSize,
		gapstart: 0,
	}
}

func (gb *GapBuffer) Insert(r rune) {
	if gb.gapsize > 0 {
		gb.buffer[gb.gapstart] = r
		gb.gapstart += 1
		gb.gapsize -= 1
	} else {
		gb.Resize()
	}
}

func (gb *GapBuffer) Delete() (rune, bool) {
	if gb.gapstart > 0 {
		gb.gapstart -= 1
		gb.gapsize += 1
        return gb.buffer[gb.gapstart], true
	}
    return ' ', false
}

func (gb *GapBuffer) Move(gapstart int) {
	if gapstart < 0 || gapstart+gb.gapsize > len(gb.buffer) {
		panic("invalid gapstart")
	}

	if gb.gapstart == gapstart {
		return
	}

	buffer := make([]rune, len(gb.buffer))
	i := 0
	for _, c := range gb.String() {
		if i == gapstart {
			i += gb.gapsize
		}
		buffer[i] = c
		i += 1
	}
	gb.gapstart = gapstart
	gb.buffer = buffer
}

func (gb *GapBuffer) Resize() {
	if gb.gapsize > 0 {
		panic("resizing with gapsize > 0")
	}

	gapsize := InitialSize

	buffer := make([]rune, len(gb.buffer)+gapsize)
	i := 0
	for _, c := range gb.String() {
		if i == gb.gapstart {
			i += gapsize
		}
		buffer[i] = c
		i += 1
	}
	gb.gapsize = gapsize
    gb.buffer = buffer
}

func (gb *GapBuffer) String() string {
	return string(gb.buffer[:gb.gapstart]) + string(gb.buffer[gb.gapstart+gb.gapsize:])
}
