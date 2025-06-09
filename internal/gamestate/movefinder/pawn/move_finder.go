package pawn

import (
	"github.com/digitalsquid7/cli-chess/internal/gamestate/board"
	"github.com/digitalsquid7/cli-chess/internal/gamestate/coordinate"
	"github.com/digitalsquid7/cli-chess/internal/gamestate/direction"
)

type MoveFinder struct {
	board board.Board
}

func NewMoveFinder(board board.Board) *MoveFinder {
	return &MoveFinder{board: board}
}

func (m *MoveFinder) FindMoves(pawn *board.Piece) []coordinate.Coordinate {
	var moves []coordinate.Coordinate

	target := pawn.Coor.Up(1)
	if pawn.Direction == direction.Down {
		target = pawn.Coor.Down(1)
	}

	curr, ok := m.board.GetPiece(target)
	if ok && curr == nil {
		moves = append(moves, target)
	}

	curr, ok = m.board.GetPiece(target.Left(1))
	if ok && curr != nil && curr.Colour != pawn.Colour {
		moves = append(moves, target.Left(1))
	}

	curr, ok = m.board.GetPiece(target.Right(1))
	if ok && curr != nil && curr.Colour != pawn.Colour {
		moves = append(moves, target.Right(1))
	}

	if pawn.Moved {
		return moves
	}

	target = pawn.Coor.Up(2)
	if pawn.Direction == direction.Down {
		target = pawn.Coor.Down(2)
	}

	curr, ok = m.board.GetPiece(target)
	if ok && curr == nil {
		moves = append(moves, target)
	}

	return moves
}
