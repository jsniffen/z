package gui

import "github.com/nsf/termbox-go"

type Modal struct {
	width  int
	height int
	x      int
	y      int
	fg     termbox.Attribute
	bg     termbox.Attribute
}

func NewModal() *Modal {
	return &Modal{
		width:  10,
		height: 10,
		x:      10,
		y:      10,
		fg:     termbox.ColorWhite,
		bg:     termbox.ColorRed,
	}
}

func (m *Modal) Render() {
	for y := m.y; y < m.y+m.height; y += 1 {
		for x := m.x; x < m.x+m.width; x += 1 {
			termbox.SetCell(x, y, ' ', m.fg, m.bg)
		}
	}
}
