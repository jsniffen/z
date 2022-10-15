package main

import (
	"log"
	"os"
	"z/edit"

	"github.com/nsf/termbox-go"
)

func main() {
	f, _ := os.OpenFile("log", os.O_RDWR|os.O_CREATE, 0777)
	defer f.Close()
	log.SetOutput(f)

	if err := termbox.Init(); err != nil {
		panic(err)
	}
	defer termbox.Close()

	pt := edit.NewPieceTable("hello, world\njulian\n1\ntesting")
	for {
		pt.Render(5, 5, 10, 10, termbox.ColorBlack, termbox.ColorWhite)
		termbox.Flush()

		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if ev.Key == termbox.KeyBackspace || ev.Key == termbox.KeyBackspace2 {
				pt.Delete()
			} else if ev.Key == termbox.KeyArrowLeft {
				pt.MoveCursorLeft()
			} else if ev.Key == termbox.KeyArrowRight {
				pt.MoveCursorRight()
			} else if ev.Key == termbox.KeyArrowUp {
				pt.MoveCursorUp()
			} else if ev.Key == termbox.KeyArrowDown {
				pt.MoveCursorDown()
			} else if ev.Key == termbox.KeyCtrlQ {
				return
			} else {
				pt.Insert(byte(ev.Ch))
			}
		}
	}
}
