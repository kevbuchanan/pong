package pong

type Direction int

const (
	Up      Direction = -1
	Down    Direction = 1
	Stopped Direction = 0
)

func (direction Direction) Change(position int, min int, max int) int {
	if (direction == Down && position < max) || (direction == Up && position > min) {
		return position + int(direction)
	} else {
		return position
	}
}
