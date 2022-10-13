package main

import (
	"github.com/nsf/termbox-go"
	"log"
	"os"
	"z/edit"
)

func main() {
	f, _ := os.OpenFile("log", os.O_RDWR|os.O_CREATE, 0777)
	defer f.Close()
	log.SetOutput(f)

	if err := termbox.Init(); err != nil {
		panic(err)
	}
	defer termbox.Close()

	pt := edit.NewPieceTable("hello, world\njulian")
	for {
    pt.Render(0, 0, 10, 10, termbox.ColorRed, termbox.ColorGreen)
		termbox.Flush()

		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if ev.Key == termbox.KeyBackspace || ev.Key == termbox.KeyBackspace2 {
				pt.Delete()

			} else if ev.Key == termbox.KeyArrowLeft {
				pt.MoveCursorLeft()
			} else if ev.Key == termbox.KeyArrowRight {
				pt.MoveCursorRight()
			} else if ev.Key == termbox.KeyCtrlQ {
				return
			} else {
				pt.Insert(byte(ev.Ch))
			}
		}
	}
}
