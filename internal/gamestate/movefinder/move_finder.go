package movefinder

import (
	"github.com/digitalsquid7/cli-chess/internal/character"
	"github.com/digitalsquid7/cli-chess/internal/gamestate/board"
	"github.com/digitalsquid7/cli-chess/internal/gamestate/coordinate"
	"github.com/digitalsquid7/cli-chess/internal/gamestate/movefinder/bishop"
	"github.com/digitalsquid7/cli-chess/internal/gamestate/movefinder/king"
	"github.com/digitalsquid7/cli-chess/internal/gamestate/movefinder/knight"
	"github.com/digitalsquid7/cli-chess/internal/gamestate/movefinder/pawn"
	"github.com/digitalsquid7/cli-chess/internal/gamestate/movefinder/queen"
	"github.com/digitalsquid7/cli-chess/internal/gamestate/movefinder/rook"
)

type PieceMoveFinder interface {
	FindMoves(piece *board.Piece) []coordinate.Coordinate
}

type MoveFinder struct {
	moveFinders map[character.Character]PieceMoveFinder
	board       board.Board
}

func New(board board.Board) *MoveFinder {
	return &MoveFinder{
		board: board,
		moveFinders: map[character.Character]PieceMoveFinder{
			character.Pawn:   pawn.NewMoveFinder(board),
			character.Rook:   rook.NewMoveFinder(board),
			character.Knight: knight.NewMoveFinder(board),
			character.Bishop: bishop.NewMoveFinder(board),
			character.Queen:  queen.NewMoveFinder(board),
			character.King:   king.NewMoveFinder(board),
		},
	}
}

func (m *MoveFinder) FindMoves(piece *board.Piece) []coordinate.Coordinate {
	if piece == nil {
		return nil
	}

	moveFinder := m.moveFinders[piece.Character]
	if moveFinder == nil {
		return nil
	}

	return moveFinder.FindMoves(piece)
}
