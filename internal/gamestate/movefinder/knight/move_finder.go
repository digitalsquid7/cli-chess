package knight

import (
	"github.com/digitalsquid7/cli-chess/internal/gamestate/board"
	"github.com/digitalsquid7/cli-chess/internal/gamestate/coordinate"
)

type MoveFinder struct {
	board board.Board
}

func NewMoveFinder(board board.Board) *MoveFinder {
	return &MoveFinder{board: board}
}

func (m *MoveFinder) FindMoves(coor coordinate.Coordinate) []coordinate.Coordinate {
	var moves []coordinate.Coordinate
	curr, _ := m.board.GetPiece(coor)

	targets := []coordinate.Coordinate{
		coor.Up(1).Right(2),
		coor.Up(1).Left(2),
		coor.Right(1).Up(2),
		coor.Right(1).Down(2),
		coor.Down(1).Right(2),
		coor.Down(1).Left(2),
		coor.Left(1).Up(2),
		coor.Left(1).Down(2),
	}

	for _, target := range targets {
		piece, ok := m.board.GetPiece(target)
		if ok && (piece == nil || piece.Colour != curr.Colour) {
			moves = append(moves, target)
		}
	}

	return moves
}
