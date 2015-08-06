package main

import (
	"github.com/alfonsodev/gotris/menu"
	"github.com/alfonsodev/gotris/playfield"
	"github.com/alfonsodev/gotris/render"
	"github.com/alfonsodev/gotris/tetromino"
	"github.com/alfonsodev/gotris/player"
 
)

func main() {
	// time.Sleep(time.Millisecond)

	var player1 player.Gamer
	player1 := NewHumanPlayer("alfonso")
	m := menu.NewMenu(0, []string{"Play", "Exit"})
	render.Start()
	defer render.Close()
	m.Show()
	t := tetromino.New()
	render.Clear()
	render.PrintMenu(m.Selected, m.Options)
	render.Render()
	t.Show()
	p := playfield.New()
	p.Show()

}
