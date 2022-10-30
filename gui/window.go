package gui

import (
	"fmt"
	"strings"
	"z/edit"

	"github.com/nsf/termbox-go"
)

var debugMessage = ""

func debug(s string) {
	debugMessage = s + "                      "
}

func min(a, b int) int {
	if a < b { return a
	}
	return b
}

type Window struct {
	gb     *edit.GapBuffer
	active bool
	cx     int
	cy     int
	x0     int
	y0     int
	w      int
	h      int
	fg     termbox.Attribute
	bg     termbox.Attribute
	lines  []string
}

func NewWindow(x, y, w, h int, fg, bg termbox.Attribute) *Window {
	gb := edit.NewGapBuffer()
	return &Window{
		active: true,
		gb:     gb,
		cx:     0,
		cy:     0,
		x0:     x,
		y0:     y,
		w:      w,
		h:      h,
		fg:     fg,
		bg:     bg,
		lines:  make([]string, 1),
	}
}

func (w *Window) calculateLines() {
	w.lines = strings.Split(w.gb.String(), "\n")
}

func (w *Window) updateCursor() {
	i := 0
	for y := 0; y < w.cy; y += 1 {
		i += len(w.lines[y]) + 1
	}
	w.gb.Move(i + w.cx)
	debug(fmt.Sprintf("moving to %d", i+w.cx))
}

func (w *Window) SetSize(width, height int) {
	w.w = width
	w.h = height
}

func (w *Window) HandleEvent(e termbox.Event) {
	switch e.Key {
	case termbox.KeyArrowDown:
		if w.cy < w.h && w.cy < len(w.lines)-1 {
			w.cy += 1
			w.cx = min(w.cx, len(w.lines[w.cy]))
			w.updateCursor()
		}
	case termbox.KeyArrowUp:
		if w.cy > 0 {
			w.cy -= 1
			w.cx = min(w.cx, len(w.lines[w.cy]))
			w.updateCursor()
		}
	case termbox.KeyArrowLeft:
		if w.cx > 0 {
			w.cx -= 1
			w.updateCursor()
		} else if w.cy > 0 {
			w.cy -= 1
			w.cx = len(w.lines[w.cy])
			w.updateCursor()
		}
	case termbox.KeyArrowRight:
		if w.cx < w.w-1 && w.cx < len(w.lines[w.cy]) {
			w.cx += 1
			w.updateCursor()
		}
	case termbox.KeyBackspace:
		fallthrough
	case termbox.KeyBackspace2:
		r, ok := w.gb.Delete()
		if ok {
			if r == '\n' {
				w.cy -= 1
				w.cx = len(w.lines[w.cy])
			} else {
				w.cx -= 1
			}
		}
	case termbox.KeyEnter:
		w.gb.Insert(rune('\n'))
		w.cy += 1
		w.cx = 0
	default:
		w.gb.Insert(rune(e.Ch))
		w.cx += 1
	}
	w.calculateLines()
}

func (w *Window) Render() {
	for y := 0; y < w.h; y += 1 {
		for x := 0; x < w.w; x += 1 {
			c := ' '
			if y < len(w.lines) {
				if x < len(w.lines[y]) {
					c = rune(w.lines[y][x])
				}
			}
			termbox.SetCell(x+w.x0, y+w.y0, c, w.fg, w.bg)
		}
	}
	termbox.SetCursor(w.cx+w.x0, w.cy+w.y0)
}
