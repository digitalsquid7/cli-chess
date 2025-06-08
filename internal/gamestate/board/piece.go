package board

import (
	"github.com/digitalsquid7/cli-chess/internal/character"
	"github.com/digitalsquid7/cli-chess/internal/gamestate/direction"
	"github.com/digitalsquid7/cli-chess/internal/gamestate/gamecolour"
)

type Piece struct {
	Character character.Character
	Colour    gamecolour.Colour
	Direction direction.Direction
}
