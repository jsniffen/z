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

	w := gui.NewWindow()
	for {
		w.Render()
		termbox.Flush()

		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if ev.Key == termbox.KeyCtrlQ {
				return
			} else {
				w.HandleEvent(ev)
			}
		}
	}
}
