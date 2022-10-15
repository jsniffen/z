package edit

import (
	"fmt"
	"strings"

	"github.com/nsf/termbox-go"
)

type Entry struct {
	add    bool
	start  int
	length int
}

func split(e Entry, i int) (Entry, Entry) {
	prev := Entry{
		add:    e.add,
		start:  e.start,
		length: i,
	}

	next := Entry{
		add:    e.add,
		start:  e.start + i,
		length: e.length - i,
	}

	return prev, next
}

func shiftRight(entries []Entry, start, n int) []Entry {
	for i := 0; i < n; i++ {
		entries = append(entries, Entry{})
	}

	for i := len(entries) - 1; i >= start+n; i-- {
		entries[i] = entries[i-n]
	}

	return entries
}

type PieceTable struct {
	originalBuffer []byte
	addBuffer      []byte
	table          []Entry
	cursor         int
	length         int
	cx             int
	cy             int
}

func (pt *PieceTable) String() string {
	var sb strings.Builder
	for _, entry := range pt.table {
		var b []byte
		if entry.add {
			b = pt.addBuffer[entry.start : entry.start+entry.length]
		} else {
			b = pt.originalBuffer[entry.start : entry.start+entry.length]
		}
		sb.Write(b)
	}
	return sb.String()
}

func NewPieceTable(init string) *PieceTable {
	pt := &PieceTable{
		originalBuffer: []byte(init),
		addBuffer:      make([]byte, 0),
		table:          make([]Entry, 0),
		cursor:         0,
		length:         len(init),
		cx:             0,
		cy:             0,
	}

	pt.table = append(pt.table, Entry{
		add:    false,
		start:  0,
		length: len(init),
	})

	return pt
}

func (pt *PieceTable) Insert(b byte) {
	newEntry := Entry{
		add:    true,
		start:  len(pt.addBuffer),
		length: 1,
	}

	pt.addBuffer = append(pt.addBuffer, b)

	start := 0
	for i, entry := range pt.table {
		end := start + entry.length

		if pt.cursor == start {
			pt.table = shiftRight(pt.table, i, 1)
			pt.table[i] = newEntry
		} else if pt.cursor == end-1 {
			pt.table = shiftRight(pt.table, i, 1)
			pt.table[i+1] = newEntry
		} else if pt.cursor > start && pt.cursor < end {
			pt.table = shiftRight(pt.table, i, 2)
			pt.table[i], pt.table[i+2] = split(entry, pt.cursor-start)
			pt.table[i+1] = newEntry
		}

		start += entry.length
	}

	pt.length += 1
	pt.MoveCursorRight()
}

func (pt *PieceTable) Delete() {
	start := 0
	for i, entry := range pt.table {
		end := start + entry.length

		if pt.cursor == start {
			pt.table[i].start += 1
			pt.table[i].length -= 1
			if pt.table[i].length == 0 {
				pt.table = append(pt.table[:i], pt.table[i+1:]...)
			}
			break
		} else if pt.cursor == end-1 {
			pt.table[i].length -= 1
			if pt.table[i].length == 0 {
				pt.table = append(pt.table[:i], pt.table[i+1:]...)
			}
			break
		} else if pt.cursor > start && pt.cursor < end {
			pt.table = shiftRight(pt.table, i, 1)
			pt.table[i], pt.table[i+1] = split(entry, pt.cursor-start)
			pt.table[i+1].start += 1
			pt.table[i+1].length -= 1
			break
		}

		if pt.table[i].length == 0 {
			pt.table = append(pt.table[:i], pt.table[i+1:]...)
		}

		start += entry.length
	}

	pt.length -= 1
	pt.MoveCursorLeft()
}

func (pt *PieceTable) MoveCursorUp() {
	if pt.cy > 0 {
		pt.cy -= 1
	}
}

func (pt *PieceTable) MoveCursorDown() {
	pt.cy += 1
}

func (pt *PieceTable) MoveCursorLeft() {
	if pt.cx > 0 {
		pt.cx -= 1
	}
}

func (pt *PieceTable) MoveCursorRight() {
	pt.cx += 1
}

func (pt *PieceTable) Render(x0, y0, w, h int, fg, bg termbox.Attribute) {
	for y := y0; y < y0+h; y += 1 {
		for x := x0; x < x0+w; x += 1 {
			termbox.SetCell(x, y, ' ', fg, bg)
		}
	}

	shiftX := max(0, pt.cx - w + 1)
	shiftY := max(0, pt.cy - h + 1)

	debug(fmt.Sprintf("%d, %d, %d, %d", shiftX, pt.cx, x0, w))

	lines := strings.Split(pt.String(), "\n")
	if len(lines) > shiftY {
		for y, line := range lines[shiftY:] {
			if len(line) > shiftX {
				for x, c := range line[shiftX:] {
					if x < w {
						termbox.SetCell(x+x0, y+y0, c, fg, bg)
					}
				}
			}
		}
	}

	cx, cy := min(pt.cx, w-1), min(pt.cy, h-1)
	termbox.SetCursor(x0+cx, y0+cy)
}

var debugLine = 0

func debug(msg ...string) {
	for i := 0; i < 100; i += 1 {
		termbox.SetCell(i, 20+debugLine, ' ', termbox.ColorGreen, termbox.ColorBlack)
	}
	x := 0
	for _, s := range msg {
		for _, c := range s {
			termbox.SetCell(x, 20+debugLine, c, termbox.ColorGreen, termbox.ColorBlack)
			x += 1
		}
		x += 1
	}
	debugLine += 1
	if debugLine > 30 {
		debugLine = 0
	}
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
