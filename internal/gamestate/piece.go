package gamestate

import (
	"github.com/digitalsquid7/cli-chess/internal/character"
	"github.com/digitalsquid7/cli-chess/internal/colour"
	"github.com/digitalsquid7/cli-chess/internal/gamestate/direction"
)

type Piece struct {
	Character character.Character
	Colour    colour.Colour
	Direction direction.Direction
}
