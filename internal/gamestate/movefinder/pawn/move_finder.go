package pawn

import (
	"github.com/digitalsquid7/cli-chess/internal/gamestate/board"
	"github.com/digitalsquid7/cli-chess/internal/gamestate/coordinate"
	"github.com/digitalsquid7/cli-chess/internal/gamestate/direction"
)

type PawnMoveFinder struct {
	board board.Board
}

func NewMoveFinder(board board.Board) *PawnMoveFinder {
	return &PawnMoveFinder{board: board}
}

func (m *PawnMoveFinder) FindMoves(coor coordinate.Coordinate) []coordinate.Coordinate {
	var moves []coordinate.Coordinate
	curr, _ := m.board.GetPiece(coor)

	temp := coor.Up(1)
	if curr.Direction == direction.Down {
		temp = coor.Down(1)
	}

	piece, ok := m.board.GetPiece(temp)
	if ok && piece == nil {
		moves = append(moves, temp)
	}

	piece, ok = m.board.GetPiece(temp.Left(1))
	if ok && piece != nil && piece.Colour != curr.Colour {
		moves = append(moves, temp.Left(1))
	}

	piece, ok = m.board.GetPiece(temp.Right(1))
	if ok && piece != nil && piece.Colour != curr.Colour {
		moves = append(moves, temp.Right(1))
	}

	return moves
}
