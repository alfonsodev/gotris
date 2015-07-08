package menu

import (
	"github.com/alfonsodev/gotris/render"
	"time"
	//	"os"
)

type Menu struct {
	Selected int
	Options  []string
}

func (m *Menu) Show() int {
	for {
		render.Clear()
		render.PrintMenu(m.Selected, m.Options)
		render.Render()
		select {
		case ev := <-render.EventQueue:
			switch {
			case ev.Key == render.Enter:
				return m.Selected
			case ev.Key == render.Esc:
				return 0
			case ev.Key == render.Down:
				if m.Selected == 0 {
					m.Selected++
				} else {
					m.Selected = 0
				}
			case ev.Key == render.Up:
				if m.Selected == 0 {
					m.Selected = 1
				} else {
					m.Selected--
				}
			}
		default:
			render.PrintMenu(m.Selected, m.Options)
			render.Render()
			time.Sleep(100 * time.Millisecond)
		}

	}
}

func NewMenu(selected int, options []string) Menu {
	var m Menu
	m.Selected = selected
	m.Options = options
	return m
}
