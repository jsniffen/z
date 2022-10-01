package main 

import (
	"z/edit"
	"github.com/nsf/termbox-go"
)

type Buffer struct {
	pt *edit.PieceTable
	cursor int
}

func NewBuffer() *Buffer {
	return &Buffer{
		pt: edit.NewPieceTable(""),
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
