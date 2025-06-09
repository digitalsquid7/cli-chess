package queen

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

func (m *MoveFinder) FindMoves(queen *board.Piece) []coordinate.Coordinate {
	var moves []coordinate.Coordinate

	moves = append(moves, m.findMoves(queen, coordinate.Up)...)
	moves = append(moves, m.findMoves(queen, coordinate.Down)...)
	moves = append(moves, m.findMoves(queen, coordinate.Right)...)
	moves = append(moves, m.findMoves(queen, coordinate.Left)...)
	moves = append(moves, m.findMoves(queen, coordinate.TopRight)...)
	moves = append(moves, m.findMoves(queen, coordinate.TopLeft)...)
	moves = append(moves, m.findMoves(queen, coordinate.BottomRight)...)
	moves = append(moves, m.findMoves(queen, coordinate.BottomLeft)...)

	return moves
}

func (m *MoveFinder) findMoves(queen *board.Piece, move func(coordinate.Coordinate, int) coordinate.Coordinate) []coordinate.Coordinate {
	var moves []coordinate.Coordinate

	incr := 0

	for {
		incr++
		curr, ok := m.board.GetPiece(move(queen.Coor, incr))
		if !ok {
			break
		}

		if curr == nil {
			moves = append(moves, move(queen.Coor, incr))
			continue
		}

		if curr.Colour != queen.Colour {
			moves = append(moves, move(queen.Coor, incr))
		}

		break
	}

	return moves
}
