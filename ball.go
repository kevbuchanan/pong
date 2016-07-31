package pong

import (
	"time"
)

type Ball struct {
	window Window
	maxX   int
	maxY   int
	x      int
	vx     float64
	y      int
	vy     float64
}

func (ball *Ball) isCollidingLeft(paddle *Paddle) bool {
	return ball.x <= paddle.Face() && ball.y <= paddle.Bottom() && ball.y >= paddle.Top()
}

func (ball *Ball) isCollidingRight(paddle *Paddle) bool {
	return ball.x >= paddle.Face() && ball.y <= paddle.Bottom() && ball.y >= paddle.Top()
}

func (ball *Ball) addSpin(paddle *Paddle) {
	if paddle.IsMovingUp() {
		if ball.vy == 0 {
			ball.vy = -1
		}
		if ball.vy < 0 {
			ball.vy = ball.vy * 1.5
		} else {
			ball.vy = ball.vy * 0.5
		}
	}
	if paddle.IsMovingDown() {
		if ball.vy == 0 {
			ball.vy = 1
		}
		if ball.vy < 0 {
			ball.vy = ball.vy * 0.5
		} else {
			ball.vy = ball.vy * 1.5
		}
	}
}

func (ball *Ball) Update(leftPaddle *Paddle, rightPaddle *Paddle) {
	ticks := time.NewTicker(time.Second / 16)

	for range ticks.C {
		ball.x = ball.x + int(1*ball.vx)
		ball.y = ball.y + int(1*ball.vy)
		if ball.isCollidingLeft(leftPaddle) {
			ball.x = leftPaddle.Face()
			ball.vx = -ball.vx
			ball.addSpin(leftPaddle)
		}
		if ball.isCollidingRight(rightPaddle) {
			ball.x = rightPaddle.Face()
			ball.vx = -ball.vx
			ball.addSpin(rightPaddle)
		}
		if ball.x > ball.maxX {
			ball.x = ball.maxX
			ball.vx = -ball.vx
		}
		if ball.x < 0 {
			ball.x = 0
			ball.vx = -ball.vx
		}
		if ball.y > ball.maxY {
			ball.y = ball.maxY
			ball.vy = -ball.vy
		}
		if ball.y < 0 {
			ball.y = 0
			ball.vy = -ball.vy
		}
	}
}

func (ball Ball) Draw() Window {
	ball.window.Move(ball.y, ball.x)
	ball.window.Print(0, 0, "o")
	return ball.window
}

func NewBall(ui Ui) Ball {
	maxRow, maxCol := ui.MaxRowAndColumn()
	return Ball{
		window: ui.NewWindow(1, 1),
		maxX:   maxCol,
		maxY:   maxRow,
		x:      maxCol / 2,
		vx:     2,
		y:      maxRow / 2,
		vy:     0,
	}
}
