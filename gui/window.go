package gui

import (
	"strings"
	"z/edit"

	"github.com/nsf/termbox-go"
)

type Window struct {
	gb *edit.GapBuffer
	cx int
	cy int
	x0 int
	y0 int
	w  int
	h  int
}

func NewWindow() *Window {
	gb := edit.NewGapBuffer("hello\n\nhello\nhi")
	return &Window{
		gb: gb,
		cx: 0,
		cy: 0,
		x0: 0,
		y0: 0,
		w:  10,
		h:  10,
	}
}

func (w *Window) HandleEvent(e termbox.Event) {
	switch e.Key {
	case termbox.KeyArrowDown:
		if w.cy < w.h {
			w.cy += 1
		}
	case termbox.KeyArrowUp:
		if w.cy > 0 {
			w.cy -= 1
		}
	case termbox.KeyArrowLeft:
		if w.cx > 0 {
			w.cx -= 1
		}
	case termbox.KeyArrowRight:
		if w.cx < w.w-1 {
			w.cx += 1
		}
  case termbox.KeyBackspace:
  case termbox.KeyBackspace2:
   w.gb.Delete()
  default:
    w.gb.Insert(rune(e.Ch))
	}
}

func (w *Window) Render() {
	lines := strings.Split(w.gb.String(), "\n")

	for y := 0; y < w.h; y += 1 {
		for x := 0; x < w.w; x += 1 {
			c := ' '
			if y < len(lines) {
				if x < len(lines[y]) {
					c = rune(lines[y][x])
				}
			}
			termbox.SetCell(x+w.x0, y+w.y0, c, termbox.ColorGreen, termbox.ColorBlue)
		}
	}
	termbox.SetCursor(w.cx+w.x0, w.cy+w.y0)
}
