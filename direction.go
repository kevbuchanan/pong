package pong

type Direction int

const (
	Up   Direction = 1
	Down Direction = -1
)

func (direction Direction) Change(position int, min int, max int) int {
	if (direction == Up && position < max) || (direction == Down && position > min) {
		return position + int(direction)
	} else {
		return position
	}
}
