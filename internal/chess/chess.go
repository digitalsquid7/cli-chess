package chess

import (
	"fmt"
	"github.com/digitalsquid7/cli-chess/internal/ansi"
	"github.com/digitalsquid7/cli-chess/internal/command"
	"github.com/digitalsquid7/cli-chess/internal/gamestate"
	"github.com/digitalsquid7/cli-chess/internal/screenupdater"
	term "github.com/nsf/termbox-go"
	"time"
)

type Game struct{}

func NewGame() *Game {
	return &Game{}
}

func (g *Game) Play() error {
	fmt.Print(ansi.EnableAlternativeScreen)
	defer fmt.Print(ansi.DisableAlternativeScreen)

	err := term.Init()
	if err != nil {
		return fmt.Errorf("term.Init: %w", err)
	}
	defer term.Close()

	gameState := gamestate.NewGameState()
	screenUpdater := screenupdater.New(gameState)
	commandExecutor := command.NewExecutor(gameState)

	for range time.Tick(time.Second / 15) {
		end, err := commandExecutor.Execute()
		if end {
			return nil
		}

		if err != nil {
			return fmt.Errorf("commandExecutor.Execute: %w", err)
		}

		if err = screenUpdater.Update(); err != nil {
			return fmt.Errorf("screenUpdater.Update: %w", err)
		}
	}

	return nil
}
