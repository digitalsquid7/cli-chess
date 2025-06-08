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

func (m *MoveFinder) FindMoves(coor coordinate.Coordinate) []coordinate.Coordinate {
	var moves []coordinate.Coordinate
	curr, _ := m.board.GetPiece(coor)

	target := coor.Up(1)
	if curr.Direction == direction.Down {
		target = coor.Down(1)
	}

	piece, ok := m.board.GetPiece(target)
	if ok && piece == nil {
		moves = append(moves, target)
	}

	piece, ok = m.board.GetPiece(target.Left(1))
	if ok && piece != nil && piece.Colour != curr.Colour {
		moves = append(moves, target.Left(1))
	}

	piece, ok = m.board.GetPiece(target.Right(1))
	if ok && piece != nil && piece.Colour != curr.Colour {
		moves = append(moves, target.Right(1))
	}

	if curr.Moved {
		return moves
	}

	target = coor.Up(2)
	if curr.Direction == direction.Down {
		target = coor.Down(2)
	}

	piece, ok = m.board.GetPiece(target)
	if ok && piece == nil {
		moves = append(moves, target)
	}

	return moves
}
