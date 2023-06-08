package websocketsPkg

import "io"

type Game interface {
	Start(numberOfPlayers int, alertsDuration io.Writer)
	Finish(winner string)
}
