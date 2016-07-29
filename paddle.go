package pong

type Side int

const (
	Left = iota
	Right
)

type Paddle struct {
	Position int
	input    chan Direction
	window   Window
	column   int
	min      int
	max      int
}

func (paddle Paddle) Up() {
	paddle.input <- Up
}

func (paddle Paddle) Down() {
	paddle.input <- Down
}

func (paddle *Paddle) Update() {
	for direction := range paddle.input {
		paddle.Position = direction.Change(paddle.Position, paddle.min, paddle.max)
	}
}

func (paddle Paddle) Collide() {
}

func (paddle Paddle) Draw() Window {
	paddle.window.ColorOn(1)
	paddle.window.Move(paddle.Position, paddle.column)
	paddle.window.Print(0, 0, "||")
	paddle.window.Print(1, 0, "||")
	paddle.window.Print(2, 0, "||")
	paddle.window.Print(3, 0, "||")
	paddle.window.Print(4, 0, "||")
	return paddle.window
}

func NewPaddle(side Side, ui Ui) Paddle {
	height := 5
	width := 2
	maxRow, maxColumn := ui.MaxRowAndColumn()
	var column int
	if side == Left {
		column = 0
	} else {
		column = maxColumn - width
	}
	return Paddle{
		input:    make(chan Direction, 10),
		window:   ui.NewWindow(height, width),
		Position: 0,
		column:   column,
		min:      0,
		max:      maxRow - height,
	}
}
