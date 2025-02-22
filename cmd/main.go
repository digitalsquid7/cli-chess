package main

import (
	"github.com/digitalsquid7/cli-chess/internal/chess"
)

func main() {
	game := chess.NewGame()
	game.Play()
}
