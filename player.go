package pong

type Player struct {
	UpKey   rune
	DownKey rune
	Paddle  Paddle
}

func NewPlayer(up rune, down rune, paddle Paddle) Player {
	return Player{
		UpKey:   up,
		DownKey: down,
		Paddle:  paddle,
	}
}
