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

func (m *MoveFinder) FindMoves(knight *board.Piece) []coordinate.Coordinate {
	var moves []coordinate.Coordinate

	targets := []coordinate.Coordinate{
		knight.Coor.Up(1).Right(2),
		knight.Coor.Up(1).Left(2),
		knight.Coor.Right(1).Up(2),
		knight.Coor.Right(1).Down(2),
		knight.Coor.Down(1).Right(2),
		knight.Coor.Down(1).Left(2),
		knight.Coor.Left(1).Up(2),
		knight.Coor.Left(1).Down(2),
	}

	for _, target := range targets {
		piece, ok := m.board.GetPiece(target)
		if ok && (piece == nil || piece.Colour != knight.Colour) {
			moves = append(moves, target)
		}
	}

	return moves
}
