package board

import (
	"github.com/digitalsquid7/cli-chess/internal/character"
	"github.com/digitalsquid7/cli-chess/internal/gamestate/coordinate"
	"github.com/digitalsquid7/cli-chess/internal/gamestate/direction"
	"github.com/digitalsquid7/cli-chess/internal/gamestate/gamecolour"
)

type Board [][]*Piece

func New() Board {
	board := Board{
		{
			&Piece{Colour: gamecolour.Black, Character: character.Rook},
			&Piece{Colour: gamecolour.Black, Character: character.Knight},
			&Piece{Colour: gamecolour.Black, Character: character.Bishop},
			&Piece{Colour: gamecolour.Black, Character: character.Queen},
			&Piece{Colour: gamecolour.Black, Character: character.King},
			&Piece{Colour: gamecolour.Black, Character: character.Bishop},
			&Piece{Colour: gamecolour.Black, Character: character.Knight},
			&Piece{Colour: gamecolour.Black, Character: character.Rook},
		},
		{
			&Piece{Colour: gamecolour.Black, Character: character.Pawn, Direction: direction.Down},
			&Piece{Colour: gamecolour.Black, Character: character.Pawn, Direction: direction.Down},
			&Piece{Colour: gamecolour.Black, Character: character.Pawn, Direction: direction.Down},
			&Piece{Colour: gamecolour.Black, Character: character.Pawn, Direction: direction.Down},
			&Piece{Colour: gamecolour.Black, Character: character.Pawn, Direction: direction.Down},
			&Piece{Colour: gamecolour.Black, Character: character.Pawn, Direction: direction.Down},
			&Piece{Colour: gamecolour.Black, Character: character.Pawn, Direction: direction.Down},
			&Piece{Colour: gamecolour.Black, Character: character.Pawn, Direction: direction.Down},
		},
		{nil, nil, nil, nil, nil, nil, nil, nil},
		{nil, nil, nil, nil, nil, nil, nil, nil},
		{nil, nil, nil, nil, nil, nil, nil, nil},
		{nil, nil, nil, nil, nil, nil, nil, nil},
		{
			&Piece{Colour: gamecolour.White, Character: character.Pawn, Direction: direction.Up},
			&Piece{Colour: gamecolour.White, Character: character.Pawn, Direction: direction.Up},
			&Piece{Colour: gamecolour.White, Character: character.Pawn, Direction: direction.Up},
			&Piece{Colour: gamecolour.White, Character: character.Pawn, Direction: direction.Up},
			&Piece{Colour: gamecolour.White, Character: character.Pawn, Direction: direction.Up},
			&Piece{Colour: gamecolour.White, Character: character.Pawn, Direction: direction.Up},
			&Piece{Colour: gamecolour.White, Character: character.Pawn, Direction: direction.Up},
			&Piece{Colour: gamecolour.White, Character: character.Pawn, Direction: direction.Up},
		},
		{
			&Piece{Colour: gamecolour.White, Character: character.Rook},
			&Piece{Colour: gamecolour.White, Character: character.Knight},
			&Piece{Colour: gamecolour.White, Character: character.Bishop},
			&Piece{Colour: gamecolour.White, Character: character.Queen},
			&Piece{Colour: gamecolour.White, Character: character.King},
			&Piece{Colour: gamecolour.White, Character: character.Bishop},
			&Piece{Colour: gamecolour.White, Character: character.Knight},
			&Piece{Colour: gamecolour.White, Character: character.Rook},
		},
	}

	for y := range board {
		for x := range board[y] {
			if board[y][x] != nil {
				board[y][x].Coor = coordinate.Make(x, y)
			}
		}
	}

	return board
}

func (b Board) GetPiece(coor coordinate.Coordinate) (*Piece, bool) {
	if !b.insideBoard(coor) {
		return nil, false
	}
	return b[coor.Y()][coor.X()], true
}

func (b Board) SetPiece(coor coordinate.Coordinate, piece *Piece) {
	if piece != nil {
		piece.Coor = coor
		piece.Moved = true
	}

	b[coor.Y()][coor.X()] = piece
}

func (b Board) insideBoard(coor coordinate.Coordinate) bool {
	return coor.X() >= 0 && coor.X() <= 7 && coor.Y() >= 0 && coor.Y() <= 7
}
