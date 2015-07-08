package tetromino

import (
	"github.com/alfonsodev/gotris/render"
	"time"
)

func New() Tetromino {
	var t Tetromino
	t.init()
	return t
}

type Tetromino struct {
	matrix      [][][][]int8
	Type        int8
	PosX        int8
	PosY        int8
	Orientation int8
	Color       int8
}

func (t *Tetromino) GetColor() uint16 {
	return uint16(t.Type)
}

func (t *Tetromino) GetMatrix() [][]int8 {
	return t.matrix[t.Type][t.Orientation]
}

func (t *Tetromino) init() {
	t.matrix = [][][][]int8{
		[][][]int8{
			[][]int8{[]int8{0, 0, 0, 0}, []int8{1, 1, 1, 1}, []int8{0, 0, 0, 0}, []int8{0, 0, 0, 0}}, // i 6 pos 3
			[][]int8{[]int8{0, 0, 1, 0}, []int8{0, 0, 1, 0}, []int8{0, 0, 1, 0}, []int8{0, 0, 1, 0}}, // i 6 pos 1
			[][]int8{[]int8{0, 0, 0, 0}, []int8{0, 0, 0, 0}, []int8{1, 1, 1, 1}, []int8{0, 0, 0, 0}}, // i 6 pos 2
			[][]int8{[]int8{0, 1, 0, 0}, []int8{0, 1, 0, 0}, []int8{0, 1, 0, 0}, []int8{0, 1, 0, 0}}, // i 6 pos 3
		},
		[][][]int8{
			[][]int8{[]int8{1, 0, 0, 0}, []int8{1, 1, 1, 0}, []int8{0, 0, 0, 0}, []int8{0, 0, 0, 0}}, // i 6 pos 3
			[][]int8{[]int8{0, 1, 1, 0}, []int8{0, 1, 0, 0}, []int8{0, 1, 0, 0}, []int8{0, 0, 0, 0}}, // i 6 pos 1
			[][]int8{[]int8{0, 0, 0, 0}, []int8{1, 1, 1, 0}, []int8{0, 0, 1, 0}, []int8{0, 0, 0, 0}}, // i 6 pos 2
			[][]int8{[]int8{0, 1, 0, 0}, []int8{0, 1, 0, 0}, []int8{1, 1, 0, 0}, []int8{0, 0, 0, 0}}, // i 6 pos 3
		},
		[][][]int8{
			[][]int8{[]int8{0, 0, 1, 0}, []int8{1, 1, 1, 0}, []int8{0, 0, 0, 0}, []int8{0, 0, 0, 0}}, // i 6 pos 3
			[][]int8{[]int8{0, 1, 0, 0}, []int8{0, 1, 0, 0}, []int8{0, 1, 1, 0}, []int8{0, 0, 0, 0}}, // i 6 pos 1
			[][]int8{[]int8{0, 0, 0, 0}, []int8{1, 1, 1, 0}, []int8{1, 0, 0, 0}, []int8{0, 0, 0, 0}}, // i 6 pos 2
			[][]int8{[]int8{1, 1, 0, 0}, []int8{0, 1, 0, 0}, []int8{0, 1, 0, 0}, []int8{0, 0, 0, 0}}, // i 6 pos 3
		},
		[][][]int8{
			[][]int8{[]int8{0, 0, 0, 0}, []int8{0, 1, 1, 0}, []int8{0, 1, 1, 0}, []int8{0, 0, 0, 0}}, // i 6 pos 3
			[][]int8{[]int8{0, 0, 0, 0}, []int8{0, 1, 1, 0}, []int8{0, 1, 1, 0}, []int8{0, 0, 0, 0}}, // i 6 pos 1
			[][]int8{[]int8{0, 0, 0, 0}, []int8{0, 1, 1, 0}, []int8{0, 1, 1, 0}, []int8{0, 0, 0, 0}}, // i 6 pos 2
			[][]int8{[]int8{0, 0, 0, 0}, []int8{0, 1, 1, 0}, []int8{0, 1, 1, 0}, []int8{0, 0, 0, 0}}, // i 6 pos 3
		},
		[][][]int8{
			[][]int8{[]int8{0, 1, 1, 0}, []int8{1, 1, 0, 0}, []int8{0, 0, 0, 0}, []int8{0, 0, 0, 0}}, // i 6 pos 3
			[][]int8{[]int8{0, 1, 0, 0}, []int8{0, 1, 1, 0}, []int8{0, 0, 1, 0}, []int8{0, 0, 0, 0}}, // i 6 pos 1
			[][]int8{[]int8{0, 0, 0, 0}, []int8{0, 1, 1, 0}, []int8{1, 1, 0, 0}, []int8{0, 0, 0, 0}}, // i 6 pos 2
			[][]int8{[]int8{1, 0, 0, 0}, []int8{1, 1, 0, 0}, []int8{0, 1, 0, 0}, []int8{0, 0, 0, 0}}, // i 6 pos 3
		},
		[][][]int8{
			[][]int8{[]int8{0, 1, 0, 0}, []int8{1, 1, 1, 0}, []int8{0, 0, 0, 0}, []int8{0, 0, 0, 0}}, // i 6 pos 3
			[][]int8{[]int8{0, 1, 0, 0}, []int8{0, 1, 1, 0}, []int8{0, 1, 0, 0}, []int8{0, 0, 0, 0}}, // i 6 pos 1
			[][]int8{[]int8{0, 0, 0, 0}, []int8{1, 1, 1, 0}, []int8{0, 1, 0, 0}, []int8{0, 0, 0, 0}}, // i 6 pos 2
			[][]int8{[]int8{0, 1, 0, 0}, []int8{1, 1, 0, 0}, []int8{0, 1, 0, 0}, []int8{0, 0, 0, 0}}, // i 6 pos 3
		},
		[][][]int8{
			[][]int8{[]int8{1, 1, 0, 0}, []int8{0, 1, 1, 0}, []int8{0, 0, 0, 0}, []int8{0, 0, 0, 0}}, // i 6 pos 3
			[][]int8{[]int8{0, 0, 1, 0}, []int8{0, 1, 1, 0}, []int8{0, 1, 0, 0}, []int8{0, 0, 0, 0}}, // i 6 pos 1
			[][]int8{[]int8{0, 0, 0, 0}, []int8{1, 1, 0, 0}, []int8{0, 1, 1, 0}, []int8{0, 0, 0, 0}}, // i 6 pos 2
			[][]int8{[]int8{0, 1, 0, 0}, []int8{1, 1, 0, 0}, []int8{1, 0, 0, 0}, []int8{0, 0, 0, 0}}, // i 6 pos 3
		},
	}
}

//In order to meet Printable interface we need these silly getters
func (t *Tetromino) GetPosX() int8 {
	return t.PosX
}
func (t *Tetromino) GetPosY() int8 {
	return t.PosY
}

func (t *Tetromino) Show() int {
	for {
		render.Clear()
		render.PrintObject(t)
		render.Render()
		select {
		case ev := <-render.EventQueue:
			switch {
			case ev.Key == render.Enter:
				return 0
			case ev.Key == render.Esc:
				return 0
			case ev.Key == render.Down:
				if t.Orientation == 3 {
					t.Orientation = -1
				}
				t.Orientation++
			case ev.Key == render.Up:
				if t.Type == 6 {
					t.Type = -1
				}
				t.Type++
			}
		default:
			render.Render()
			time.Sleep(100 * time.Millisecond)
		}

	}
}
