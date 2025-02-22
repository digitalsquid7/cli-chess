package gamestate

import "github.com/digitalsquid7/cli-chess/internal/gamestate/coordinate"

type GameState struct {
	Board  Board
	Cursor coordinate.Coordinate
}

func NewGameState() *GameState {
	return &GameState{
		Board:  NewBoard(),
		Cursor: coordinate.Coordinate{0, 2},
	}
}

func (g *GameState) MoveLeft() {
	coor := g.Cursor.Left(1)

	if g.insideBoard(coor) {
		g.Cursor = coor
	}
}

func (g *GameState) MoveRight() {
	coor := g.Cursor.Right(1)

	if g.insideBoard(coor) {
		g.Cursor = coor
	}
}

func (g *GameState) MoveUp() {
	coor := g.Cursor.Up(1)

	if g.insideBoard(coor) {
		g.Cursor = coor
	}
}

func (g *GameState) MoveDown() {
	coor := g.Cursor.Down(1)

	if g.insideBoard(coor) {
		g.Cursor = coor
	}
}

func (g *GameState) insideBoard(coor coordinate.Coordinate) bool {
	return coor.X() >= 0 && coor.X() <= 7 && coor.Y() >= 0 && coor.Y() <= 7
}
