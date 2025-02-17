package gamestate

import (
	"github.com/digitalsquid7/cli-chess/internal/character"
	"github.com/digitalsquid7/cli-chess/internal/gamestate/direction"
)

type Board [][]*Piece

func NewBoard() Board {
	return Board{
		{
			&Piece{Character: character.BlackRook},
			&Piece{Character: character.BlackKnight},
			&Piece{Character: character.BlackBishop},
			&Piece{Character: character.BlackKing},
			&Piece{Character: character.BlackQueen},
			&Piece{Character: character.BlackBishop},
			&Piece{Character: character.BlackKnight},
			&Piece{Character: character.BlackRook},
		},
		{
			&Piece{Character: character.BlackPawn, Direction: direction.Down},
			&Piece{Character: character.BlackPawn, Direction: direction.Down},
			&Piece{Character: character.BlackPawn, Direction: direction.Down},
			&Piece{Character: character.BlackPawn, Direction: direction.Down},
			&Piece{Character: character.BlackPawn, Direction: direction.Down},
			&Piece{Character: character.BlackPawn, Direction: direction.Down},
			&Piece{Character: character.BlackPawn, Direction: direction.Down},
			&Piece{Character: character.BlackPawn, Direction: direction.Down},
		},
		{}, {}, {}, {},
		{
			&Piece{Character: character.WhitePawn, Direction: direction.Up},
			&Piece{Character: character.WhitePawn, Direction: direction.Up},
			&Piece{Character: character.WhitePawn, Direction: direction.Up},
			&Piece{Character: character.WhitePawn, Direction: direction.Up},
			&Piece{Character: character.WhitePawn, Direction: direction.Up},
			&Piece{Character: character.WhitePawn, Direction: direction.Up},
			&Piece{Character: character.WhitePawn, Direction: direction.Up},
			&Piece{Character: character.WhitePawn, Direction: direction.Up},
		},
		{
			&Piece{Character: character.WhiteRook},
			&Piece{Character: character.WhiteKnight},
			&Piece{Character: character.WhiteBishop},
			&Piece{Character: character.WhiteKing},
			&Piece{Character: character.WhiteQueen},
			&Piece{Character: character.WhiteBishop},
			&Piece{Character: character.WhiteKnight},
			&Piece{Character: character.WhiteRook},
		},
	}
}
