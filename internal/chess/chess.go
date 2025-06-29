package chess

import (
	"fmt"
	"github.com/digitalsquid7/cli-chess/internal/command"
	"github.com/digitalsquid7/cli-chess/internal/gamestate"
	"github.com/digitalsquid7/cli-chess/internal/screenupdater"
	"time"
)

type Game struct {
	commandExecutor *command.Executor
	gameState       *gamestate.GameState
	screenUpdater   *screenupdater.ScreenUpdater
}

func NewGame() *Game {
	gameState := gamestate.New()

	return &Game{
		commandExecutor: command.NewExecutor(gameState),
		gameState:       gameState,
		screenUpdater:   screenupdater.New(gameState),
	}
}

func (g *Game) Play() error {
	if err := g.screenUpdater.Update(); err != nil {
		return fmt.Errorf("screenUpdater.Update: %w", err)
	}

	for range time.Tick(time.Second / 15) {
		end, err := g.commandExecutor.Execute()
		if end {
			return nil
		}

		if err != nil {
			return fmt.Errorf("commandExecutor.Execute: %w", err)
		}

		if err = g.screenUpdater.Update(); err != nil {
			return fmt.Errorf("screenUpdater.Update: %w", err)
		}
	}

	return nil
}
