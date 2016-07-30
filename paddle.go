package pong

type Side int

const (
	Left = iota
	Right
)

type Paddle struct {
	height int
	width  int
	side   Side
	row    int
	window Window
	input  chan Direction
	column int
	min    int
	max    int
}

func (paddle Paddle) Up() {
	paddle.input <- Up
}

func (paddle Paddle) Down() {
	paddle.input <- Down
}

func (paddle Paddle) Face() int {
	if paddle.side == Left {
		return paddle.column + paddle.width
	} else {
		return paddle.column
	}
}

func (paddle Paddle) Top() int {
	return paddle.row
}

func (paddle Paddle) Bottom() int {
	return paddle.row + paddle.height
}

func (paddle *Paddle) Update() {
	for direction := range paddle.input {
		paddle.row = direction.Change(paddle.row, paddle.min, paddle.max)
	}
}

func (paddle Paddle) Collide() {
}

func (paddle Paddle) Draw() Window {
	paddle.window.ColorOn(1)
	paddle.window.Move(paddle.row, paddle.column)
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
		column = 5
	} else {
		column = maxColumn - width - 5
	}
	return Paddle{
		height: height,
		width:  width,
		side:   side,
		window: ui.NewWindow(height, width),
		input:  make(chan Direction, 10),
		row:    maxRow / 2,
		column: column,
		min:    0,
		max:    maxRow - height,
	}
}
