package main

import (
	"github.com/nsf/termbox-go"
	"log"
	"os"
	"z/edit"
)

type Buffer struct {
	pt     *edit.PieceTable
	cursor int
}

func NewBuffer() *Buffer {
	return &Buffer{
		pt:     edit.NewPieceTable("hello"),
		cursor: 0,
	}
}

func (b *Buffer) Insert(char byte) {
	b.pt.Insert(char, b.cursor)
	b.cursor += 1
}

func (b *Buffer) Delete() {
	b.pt.Delete(b.cursor)
	b.cursor -= 1
}

func (b *Buffer) Render() {
	if b.cursor >= 0 {
		termbox.SetCursor(b.cursor, 0)
	}
	x, y := 0, 0
	for _, c := range b.pt.String() {
		termbox.SetCell(x, y, c, termbox.ColorRed, termbox.ColorDefault)
		x += 1
	}
}

func main() {
	f, _ := os.OpenFile("log", os.O_RDWR|os.O_CREATE, 0777)
	defer f.Close()
	log.SetOutput(f)

	if err := termbox.Init(); err != nil {
		panic(err)
	}
	defer termbox.Close()

	buffer := NewBuffer()
	for {
		buffer.Render()
		termbox.Flush()

		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if ev.Key == termbox.KeyBackspace || ev.Key == termbox.KeyBackspace2 {
				buffer.Delete()
			} else if ev.Key == termbox.KeyCtrlQ {
				return
			} else {
				buffer.Insert(byte(ev.Ch))
			}
		}
	}
}
