package main

import (
	"strings"

	"github.com/nsf/termbox-go"
)

type Entry struct {
	add    bool
	start  int
	length int
}

type PieceTable struct {
	originalBuffer []byte
	addBuffer      []byte
	table          []Entry
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

func New(init string) *PieceTable {
	pt := &PieceTable{
		originalBuffer: []byte(init),
		addBuffer:      make([]byte, 0),
		table:          make([]Entry, 0),
	}

	pt.table = append(pt.table, Entry{
		add:    false,
		start:  0,
		length: len(init),
	})

	return pt
}

// TODO(Julian): Fix this, it doesn't work...
func (pt *PieceTable) Delete(pos int) {
	currentPos := 0
	for i, entry := range pt.table {
		if currentPos == pos {
			entry.start += 1
			entry.length -= 1
		} else if pos == currentPos+entry.length - 1{
			entry.length -= 1
		} else if pos > currentPos && pos < currentPos+entry.length {
			newLength := pos - currentPos - 1
			pt.table = append(pt.table, Entry{})

			for j := len(pt.table) - 1; j-1 > i; j -= 1 {
				pt.table[j] = pt.table[j-1]
			}

			pt.table[i+1] = Entry{
				add:    entry.add,
				start:  entry.start + newLength,
				length: entry.length - newLength,
			}

			pt.table[i].length = newLength
			return
		}
	}
}

func (pt *PieceTable) Insert(s string, pos int) {
	newEntry := Entry{
		add:    true,
		start:  len(pt.addBuffer),
		length: len(s),
	}

	pt.addBuffer = append(pt.addBuffer, []byte(s)...)

	currentPos := 0
	for i, entry := range pt.table {
		if currentPos == pos {
			pt.table = append(pt.table[:i+1], pt.table[i:]...)
			pt.table[i] = newEntry
			return
		} else if pos > currentPos && pos < currentPos+entry.length {
			newLength := pos - currentPos
			pt.table = append(pt.table, Entry{}, Entry{})

			for j := len(pt.table) - 1; j-2 > i; j -= 1 {
				pt.table[j] = pt.table[j-2]
			}

			pt.table[i+1] = newEntry
			pt.table[i+2] = Entry{
				add:    entry.add,
				start:  entry.start + newLength,
				length: entry.length - newLength,
			}

			pt.table[i].length = newLength
			return
		}

		currentPos += entry.length
	}
}

type Buffer struct {
	pt *PieceTable
	cursor int
}

func NewBuffer() *Buffer {
	return &Buffer{
		pt: New(""),
		cursor: 0,
	}
}

func (b *Buffer) Insert(s string) {
	b.pt.Insert(s, b.cursor)
	b.cursor += 1
}

func (b *Buffer) Delete() {
	b.pt.Delete(b.cursor)
	b.cursor -= 1
}

func (b *Buffer) Render() {
	termbox.SetCursor(b.cursor, 0)
	x, y := 0, 0
	for _, c := range b.pt.String() {
		termbox.SetCell(x, y, c, termbox.ColorRed, termbox.ColorDefault)
		x += 1
	}
}


func main() {
	if err := termbox.Init(); err != nil {
		panic(err)
	}
	defer termbox.Close()

	termbox.SetInputMode(termbox.InputMouse)

	buffer := NewBuffer()
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if ev.Key == termbox.KeyBackspace {
				buffer.Delete()
			} else {
				buffer.Insert(string(ev.Ch))
			}
			break
			
		case termbox.EventMouse:
			if ev.Key == termbox.MouseLeft {
				return
			}
		}

		buffer.Render()

		termbox.Flush()
	}
}
