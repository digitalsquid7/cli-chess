package king

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

func (m *MoveFinder) FindMoves(king *board.Piece) []coordinate.Coordinate {
	var moves []coordinate.Coordinate

	targets := []coordinate.Coordinate{
		king.Coor.Up(1),
		king.Coor.Right(1),
		king.Coor.Down(1),
		king.Coor.Left(1),
		king.Coor.Up(1).Right(1),
		king.Coor.Right(1).Down(1),
		king.Coor.Down(1).Left(1),
		king.Coor.Left(1).Up(1),
	}

	for _, target := range targets {
		piece, ok := m.board.GetPiece(target)
		if ok && (piece == nil || piece.Colour != piece.Colour) {
			moves = append(moves, target)
		}
	}

	moves = append(moves, m.findCastlingMoves(king)...)

	return moves
}

func (m *MoveFinder) findCastlingMoves(king *board.Piece) []coordinate.Coordinate {
	var moves []coordinate.Coordinate

	if king.Moved {
		return moves
	}

	rook, _ := m.board.GetPiece(coordinate.Make(0, king.Coor.Y()))
	if rook != nil && rook.Character == character.Rook && !rook.Moved &&
		m.board[king.Coor.Y()][1] == nil && m.board[king.Coor.Y()][2] == nil && m.board[king.Coor.Y()][3] == nil {
		moves = append(moves, rook.Coor)
	}

	rook, _ = m.board.GetPiece(coordinate.Make(7, king.Coor.Y()))
	if rook != nil && rook.Character == character.Rook && !rook.Moved &&
		m.board[king.Coor.Y()][5] == nil && m.board[king.Coor.Y()][6] == nil {
		moves = append(moves, rook.Coor)
	}

	return moves
}
