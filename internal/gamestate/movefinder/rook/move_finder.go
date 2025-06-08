package rook

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

	moves = append(moves, m.findMoves(coor, up)...)
	moves = append(moves, m.findMoves(coor, down)...)
	moves = append(moves, m.findMoves(coor, left)...)
	moves = append(moves, m.findMoves(coor, right)...)

	return moves
}

func (m *MoveFinder) findMoves(coor coordinate.Coordinate, move func(coordinate.Coordinate, int) coordinate.Coordinate) []coordinate.Coordinate {
	var moves []coordinate.Coordinate
	selected, _ := m.board.GetPiece(coor)

	incr := 0

	for {
		incr++
		curr, ok := m.board.GetPiece(move(coor, incr))
		if !ok {
			break
		}

		if curr == nil {
			moves = append(moves, move(coor, incr))
			continue
		}

		if curr.Colour != selected.Colour {
			moves = append(moves, move(coor, incr))
		}

		break
	}

	return moves
}

func up(coor coordinate.Coordinate, amount int) coordinate.Coordinate {
	return coor.Up(amount)
}

func down(coor coordinate.Coordinate, amount int) coordinate.Coordinate {
	return coor.Down(amount)
}

func left(coor coordinate.Coordinate, amount int) coordinate.Coordinate {
	return coor.Left(amount)
}

func right(coor coordinate.Coordinate, amount int) coordinate.Coordinate {
	return coor.Right(amount)
}
