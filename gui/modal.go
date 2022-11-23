package gui

import "github.com/nsf/termbox-go"

type ModalSize uint64

const (
	ModalSmall ModalSize = iota
	ModalMedium
	ModalLarge
)

type Modal struct {
	width  int
	height int
	x      int
	y      int
	fg     termbox.Attribute
	bg     termbox.Attribute
	size   ModalSize
}

func NewModal() *Modal {
	return &Modal{
		fg: termbox.ColorWhite,
		bg: termbox.ColorRed,
		size: ModalSmall,
	}
}

func (m *Modal) Render(w, h int) {
	var x0, x1, y0, y1 int
	switch m.size {
	case ModalSmall:
		x0 = w/4
		x1 = 3*x0
		y0 = h/2-1
		y1 = y0+2
	case ModalMedium:
		x0 = w / 8
		x1 = 7 * x0
		y0 = h / 4
		y1 = 3 * y0
	case ModalLarge:
		x0, x1 = 0, w
		y0, y1 = 0, h
	}

	for y := y0; y < y1; y += 1 {
		for x := x0; x < x1; x += 1 {
			termbox.SetCell(x, y, ' ', m.fg, m.bg)
		}
	}
}
