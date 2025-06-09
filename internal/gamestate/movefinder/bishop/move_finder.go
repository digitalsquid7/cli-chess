package bishop

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

func (m *MoveFinder) FindMoves(bishop *board.Piece) []coordinate.Coordinate {
	var moves []coordinate.Coordinate

	moves = append(moves, m.findMoves(bishop, coordinate.TopRight)...)
	moves = append(moves, m.findMoves(bishop, coordinate.TopLeft)...)
	moves = append(moves, m.findMoves(bishop, coordinate.BottomRight)...)
	moves = append(moves, m.findMoves(bishop, coordinate.BottomLeft)...)

	return moves
}

func (m *MoveFinder) findMoves(bishop *board.Piece, move func(coordinate.Coordinate, int) coordinate.Coordinate) []coordinate.Coordinate {
	var moves []coordinate.Coordinate

	incr := 0

	for {
		incr++
		curr, ok := m.board.GetPiece(move(bishop.Coor, incr))
		if !ok {
			break
		}

		if curr == nil {
			moves = append(moves, move(bishop.Coor, incr))
			continue
		}

		if curr.Colour != bishop.Colour {
			moves = append(moves, move(bishop.Coor, incr))
		}

		break
	}

	return moves
}
