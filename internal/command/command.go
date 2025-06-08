package command

import (
	"github.com/digitalsquid7/cli-chess/internal/gamestate"
	term "github.com/nsf/termbox-go"
)

type Executor struct {
	gameState *gamestate.GameState
}

func NewExecutor(gameState *gamestate.GameState) *Executor {
	return &Executor{gameState: gameState}
}

func (e *Executor) Execute() (bool, error) {
	event := term.PollEvent()

	if event.Type == term.EventError {
		return false, event.Err
	}

	if event.Type != term.EventKey {
		return false, nil
	}

	switch event.Key {
	case term.KeyArrowLeft:
		e.gameState.MoveLeft()
	case term.KeyArrowRight:
		e.gameState.MoveRight()
	case term.KeyArrowUp:
		e.gameState.MoveUp()
	case term.KeyArrowDown:
		e.gameState.MoveDown()
	case term.KeyEnter:
		e.gameState.SelectPiece()
	case term.KeyCtrlC:
		return true, nil
	default:
	}

	return false, term.Sync()
}
