package chess

import (
	"github.com/digitalsquid7/cli-chess/internal/character"
	"github.com/digitalsquid7/cli-chess/internal/colour"
	"github.com/digitalsquid7/cli-chess/internal/gamestate"
	"github.com/digitalsquid7/cli-chess/internal/screenupdater"
)

type Char struct {
	Colour colour.Colour
	Symbol character.Character
}

type Game struct{}

func NewGame() *Game {
	return &Game{}
}

func (g *Game) Play() {
	gameState := gamestate.NewGameState()
	screenUpdater := screenupdater.New(gameState)
	screenUpdater.Update()
}
