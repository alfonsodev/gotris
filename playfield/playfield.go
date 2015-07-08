package playfield

import (
	"github.com/alfonsodev/gotris/render"
	"time"
)

const (
	emptyBlock = 0
	wallBlock  = 9
)

func New() Playfield {
	var p Playfield
	p.init(10, 10)
	return p
}

type Playfield struct {
	Matrix [][]int8
	Width  int8
	Height int8
	Color  uint16
	PosX   int8
	PosY   int8
}

func (p *Playfield) GetPosX() int8 {
	return p.PosX
}

func (p *Playfield) GetPosY() int8 {
	return p.PosY
}

func (p *Playfield) GetColor() uint16 {
	return p.Color
}
func (p *Playfield) GetMatrix() [][]int8 {
	return p.Matrix
}

func (p *Playfield) init(givenWidth uint8, givenHeight uint8) {
	p.Matrix = make([][]int8, givenHeight)
	for y, _ := range p.Matrix {
		p.Matrix[y] = make([]int8, givenWidth)
		for x, _ := range p.Matrix[y] {
			p.Matrix[y][x] = 1
			// fmt.Println("Hello, playground" + strconv.Itoa(i)+"-" + strconv.Itoa(j))
		}
	}
}

func (p *Playfield) Show() int {
	for {
		render.Clear()
		render.PrintObject(p)
		render.Render()
		select {
		case ev := <-render.EventQueue:
			switch {
			case ev.Key == render.Enter:
				return 0
			case ev.Key == render.Esc:
				return 0
			case ev.Key == render.Down:
				return 0
			case ev.Key == render.Up:
				return 0
			}
		default:
			render.Render()
			time.Sleep(100 * time.Millisecond)
		}
	}
}
