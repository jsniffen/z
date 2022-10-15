package main

import (
	"z/edit"

	"github.com/nsf/termbox-go"
)

func main() {
	if err := termbox.Init(); err != nil {
		panic(err)
	}
	defer termbox.Close()

	gb := edit.NewGapBuffer("hello")
	for {
		gb.Render()
		termbox.Flush()

		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if ev.Key == termbox.KeyCtrlQ {
				return
			} else {
				gb.Insert(rune(ev.Ch))
			}
		}
	}
}
