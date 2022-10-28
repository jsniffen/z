package main

import (
	"github.com/nsf/termbox-go"

    "z/gui"
)

func main() {
	if err := termbox.Init(); err != nil {
		panic(err)
	}
	defer termbox.Close()

	window := gui.NewWindow(0, 0, 100, 100, termbox.ColorBlack, termbox.ColorWhite)
	for {
		window.Render()
		termbox.Flush()

		switch e := termbox.PollEvent(); e.Type {
		case termbox.EventKey:
			if e.Key == termbox.KeyCtrlQ {
				return
			} else {
				window.HandleEvent(e)
			}
		}
	}
}
