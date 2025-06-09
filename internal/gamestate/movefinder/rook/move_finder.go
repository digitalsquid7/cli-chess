package rook

import (
	"github.com/digitalsquid7/cli-chess/internal/character"
	"github.com/digitalsquid7/cli-chess/internal/gamestate/board"
	"github.com/digitalsquid7/cli-chess/internal/gamestate/coordinate"
)

type MoveFinder struct {
	board board.Board
}

func NewMoveFinder(board board.Board) *MoveFinder {
	return &MoveFinder{board: board}
}

func (m *MoveFinder) FindMoves(piece *board.Piece) []coordinate.Coordinate {
	var moves []coordinate.Coordinate

	moves = append(moves, m.findMoves(piece, coordinate.Up)...)
	moves = append(moves, m.findMoves(piece, coordinate.Down)...)
	moves = append(moves, m.findMoves(piece, coordinate.Left)...)
	moves = append(moves, m.findMoves(piece, coordinate.Right)...)

	return moves
}

func (m *MoveFinder) findMoves(piece *board.Piece, move func(coordinate.Coordinate, int) coordinate.Coordinate) []coordinate.Coordinate {
	var moves []coordinate.Coordinate

	incr := 0

	for {
		incr++
		curr, ok := m.board.GetPiece(move(piece.Coor, incr))
		if !ok {
			break
		}

		if curr == nil {
			moves = append(moves, move(piece.Coor, incr))
			continue
		}

		if curr.Colour != piece.Colour {
			moves = append(moves, move(piece.Coor, incr))
		}

		break
	}

	moves = append(moves, m.findCastlingMoves(piece)...)

	return moves
}

func (m *MoveFinder) findCastlingMoves(rook *board.Piece) []coordinate.Coordinate {
	var moves []coordinate.Coordinate

	if rook.Moved {
		return moves
	}

	king, _ := m.board.GetPiece(coordinate.Make(3, rook.Coor.Y()))

	if king.Character != character.King || king.Moved {
		return moves
	}

	if (rook.Coor.X() == 0 && m.board[rook.Coor.Y()][1] == nil && m.board[rook.Coor.Y()][2] == nil) ||
		(rook.Coor.X() == 7 && m.board[rook.Coor.Y()][4] == nil && m.board[rook.Coor.Y()][5] == nil && m.board[rook.Coor.Y()][6] == nil) {
		moves = append(moves, king.Coor)
	}

	return moves
}
