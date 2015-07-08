package render
import(
	"github.com/nsf/termbox-go"
	"strconv"
)

var EventQueue chan termbox.Event
var Public int

type Printable interface{
	GetMatrix() [][]int8
	GetPosX() int8
	GetPosY() int8
	GetColor() uint16
}

const (
	Black = termbox.ColorBlack
	White = termbox.ColorWhite

	Red = termbox.ColorRed
	Green = termbox.ColorGreen
	Yellow = termbox.ColorYellow
  	Blue = termbox.ColorBlue
  	Magenta = termbox.ColorMagenta
    Cyan = termbox.ColorCyan

	Left = termbox.KeyArrowLeft
	Right = termbox.KeyArrowRight
	Up = termbox.KeyArrowUp
	Down = termbox.KeyArrowDown
	Space = termbox.KeySpace
	Esc = termbox.KeyEsc
	Enter = termbox.KeyEnter
)

type Window struct {
	PosX int8
	PosY int8
}

func Start() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	termbox.HideCursor()
	termbox.Clear(termbox.ColorWhite, termbox.ColorBlack)
	termbox.Flush()
}
func Clear() {
	termbox.Clear(termbox.ColorWhite, termbox.ColorBlack)
}
func Render() {
	termbox.Flush()
}
func Close() {
	defer termbox.Close()
}
func init() {
	EventQueue = make(chan termbox.Event)
	go func() {
		for {
			EventQueue <- termbox.PollEvent()
		}
	}()
}

func String(x, y int, fg, bg termbox.Attribute, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x++
	}
}
func main() {

}

func PrintMenu(selected int, options []string) {
	for i, val := range options{
		if selected == i {
			String(0, i, Red, White, strconv.Itoa(i) + "> " + val + " " + strconv.Itoa(selected))
		} else {
			String(0, i, White, Black, strconv.Itoa(i) + "  " + val)
		}
	}
}

func PrintObject(p Printable){
	matrix := p.GetMatrix()
	// colors := []termbox.Attribute{Red, Green, Yellow, Blue, Magenta, Cyan, Yellow}
	// colors[p.GetColor()]
	if len(matrix[0]) == 10 {
		panic("in the disco")
	}

 	for y := len(matrix)-1; y >= 0; y-- {
    	for x := len(matrix[0])-1; x>=0; x-- {
			if (matrix[y][x] != 0) {
				String((int(p.GetPosX()) + x)*2 , int(p.GetPosY()) + y,  White, Black, "[]")
			}
	    }
	}	
}

//  positionX = this.wPosX + leftPadding + ((posX + x) * 2 + 1);
//  positionY = this.wPosY + topPadding + posY + y;
// charm.position(positionX, positionY);
// if (matrix[y][x] != 0) {
//   if (dynamic) this.switchColor(matrix[y][x]);
//   charm.write('[]');
// } else if (dynamic) {
//   this.switchColor(8);
//   charm.write('..');
// }



