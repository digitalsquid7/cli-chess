package chess

import (
	"fmt"
	"github.com/digitalsquid7/cli-chess/internal/gamestate"
	"github.com/digitalsquid7/cli-chess/internal/screenupdater"
)

type Game struct{}

func NewGame() *Game {
	return &Game{}
}

func (g *Game) Play() {
	// Activate alternate screen mode
	fmt.Print("\033[?1049h")
	defer fmt.Print("\033[?1049l")

	fmt.Print("\033[2J\033[H")
	//cmd := exec.Command("clear")
	//cmd.Stdout = os.Stdout
	//cmd.Run()

	gameState := gamestate.NewGameState()
	screenUpdater := screenupdater.New(gameState)
	screenUpdater.Update()
}
