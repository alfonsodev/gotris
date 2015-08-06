package player

import (
//	"github.com/nsf/termbox-go"
	"github.com/alfonsodev/gotris/game"
)

func NewHumanPlayer(name string) game.Gamer {
	var h Human 
	h.SetName(name)
	return h
}

type Human struct {
	points uint16
	name string
}

// func (h Human) listen() {
// 	for {
// 		select {
// 		case ev := <-render.EventQueue:
// 			switch{
// 				case ev.Key == render.Enter:
// 					h.Input <- 1
// 			}
// 		}	
// 	}
// }

func (h Human)	SetPoints(points uint16) {
	h.points = points
}
func (h Human)	GetPoints() uint16 {
	return h.points
}
func (h Human)	GetName() string {
	return h.name 
}
func (h Human)	SetName(name string) {
	h.name = name
}
