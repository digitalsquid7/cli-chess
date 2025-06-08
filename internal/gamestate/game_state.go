package gamestate

import (
	"github.com/digitalsquid7/cli-chess/internal/gamestate/board"
	"github.com/digitalsquid7/cli-chess/internal/gamestate/coordinate"
	"github.com/digitalsquid7/cli-chess/internal/gamestate/movefinder"
)

type GameState struct {
	Board         board.Board
	Cursor        coordinate.Coordinate
	SelectedPiece *coordinate.Coordinate
	MoveFinder    *movefinder.MoveFinder
	ValidMoves    []coordinate.Coordinate
}

func NewGameState() *GameState {
	gameState := &GameState{
		Board:  board.New(),
		Cursor: coordinate.Coordinate{0, 2},
	}

	gameState.MoveFinder = movefinder.New(gameState.Board)
	return gameState
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

func (g *GameState) SelectPiece() {
	if g.SelectedPiece == nil {
		g.SelectedPiece = ref(coordinate.Make(g.Cursor.X(), g.Cursor.Y()))
		g.ValidMoves = g.MoveFinder.FindMoves(g.Cursor)
		return
	}

	if *g.SelectedPiece != g.Cursor {
		*g.SelectedPiece = g.Cursor
		g.ValidMoves = g.MoveFinder.FindMoves(g.Cursor)
		return
	}

	g.ValidMoves = nil
	g.SelectedPiece = nil
}

func (g *GameState) insideBoard(coor coordinate.Coordinate) bool {
	return coor.X() >= 0 && coor.X() <= 7 && coor.Y() >= 0 && coor.Y() <= 7
}

func ref[T any](val T) *T {
	return &val
}
