package playfield

import(
    "time"
    "github.com/alfonsodev/gotris/render"
)

const (
  emptyBlock = 0 
  wallBlock = 9
)

func New() Playfield{
  var p Playfield
  p.init(10, 10) 
  return p
}

type Playfield struct {
    Matrix [][]int8
    Width int8
    Height int8 
    Color uint16
    PosX int8
    PosY int8
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

func (p *Playfield)init (givenWidth uint8, givenHeight uint8) {
  p.Matrix = make([][]int8, 10)
  for y:= givenHeight-1; y >= 0; y-- {
    p.Matrix[y] = make([]int8, 10)
    for x:= givenWidth-1; x >= 0; x-- {
      p.Matrix[y][x] = 1
      // if y == (givenHeight- 1) ||  y == 0 || x == (givenWidth - 1)  {
      //   p.Matrix[y][x] = 1 
      // } else {
      //   p.Matrix[y][x] = 0 
      // }
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
            time.Sleep(100*time.Millisecond)
        }
    }
}
