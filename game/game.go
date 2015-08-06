package game

type Gamer interface {
	SetPoints(points uint16)
	GetPoints()uint16
	GetName()string
	SetName(name string)
}

func NewGame(players []Gamer) Game {
	var game Game
	game.Players = players

	return game
}

type Game struct {
	Players []Gamer
}

func (g *Game) Start(players []Gamer){

}
func (g *Game) gameStep() {

}

