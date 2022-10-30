package gui

import (
	"github.com/nsf/termbox-go"
)

type Mode uint64

const (
	ModeNormal Mode = iota
	ModeInsert
	ModeVisual
)

type Z struct {
	mode    Mode
	running bool
	windows []*Window
	width   int
	height  int
}

func NewZ() *Z {
	if err := termbox.Init(); err != nil {
		panic(err)
	}

	w, h := termbox.Size()

	scratch := NewWindow(0, 0, 0, 0, termbox.ColorWhite, termbox.ColorBlue)
	return &Z{
		running: true,
		windows: []*Window{scratch},
		width: w,
		height: h,
	}
}

func (z *Z) Run() {
	defer termbox.Close()

	for z.running {
		z.render()
		z.handleEvent()
	}
}

func (z *Z) render() {
	for _, w := range z.windows {
		w.SetSize(z.width, z.height-1)
		w.Render()
	}

    z.renderStatus()
	termbox.Flush()
}

func (z *Z) handleEvent() {
	switch e := termbox.PollEvent(); e.Type {
	case termbox.EventResize:
		z.width = e.Width
		z.height = e.Height
	case termbox.EventKey:
        if z.mode == ModeInsert {
            z.handleInsertModeEvent(e)
        } else {
            z.handleNormalModeEvent(e)
        }
	}
}

func (z *Z) handleNormalModeEvent(e termbox.Event) {
    if e.Ch == 'i' {
        z.mode = ModeInsert
        return
    }

    if e.Key == termbox.KeyCtrlQ {
        z.running = false
        return
    }
}


func (z *Z) handleInsertModeEvent(e termbox.Event) {
    if e.Key == termbox.KeyEsc {
        z.mode = ModeNormal
        return
    }

    for _, window := range z.windows {
        if window.active {
            window.HandleEvent(e)
        }
    }
}

func (z *Z) renderStatus() {
    var status []rune

    if z.mode == ModeInsert {
        status = []rune("INSERT")
    } else if z.mode == ModeVisual {
        status = []rune("VISUAL")
    } else {
        status = []rune("NORMAL")
    }

    for x := 0; x < z.width; x += 1 {
        if x < len(status) {
            termbox.SetCell(x, z.height-1, status[x], termbox.ColorBlack, termbox.ColorWhite)
        } else {
            termbox.SetCell(x, z.height-1, ' ', termbox.ColorBlack, termbox.ColorWhite)
        }
    }
}
