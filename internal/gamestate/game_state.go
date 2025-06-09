package gamestate

import (
	"github.com/digitalsquid7/cli-chess/internal/character"
	"github.com/digitalsquid7/cli-chess/internal/gamestate/board"
	"github.com/digitalsquid7/cli-chess/internal/gamestate/coordinate"
	"github.com/digitalsquid7/cli-chess/internal/gamestate/gamecolour"
	"github.com/digitalsquid7/cli-chess/internal/gamestate/movefinder"
	"slices"
)

type GameState struct {
	Board      board.Board
	Cursor     coordinate.Coordinate
	Selected   *coordinate.Coordinate
	MoveFinder *movefinder.MoveFinder
	ValidMoves []coordinate.Coordinate
	PlayerTurn gamecolour.Colour
}

func New() *GameState {
	gameState := &GameState{
		Board:      board.New(),
		Cursor:     coordinate.Coordinate{0, 2},
		PlayerTurn: gamecolour.White,
	}

	gameState.MoveFinder = movefinder.New(gameState.Board)
	return gameState
}

func (g *GameState) MoveLeft() {
	g.moveCursor(g.Cursor.Left(1))
}

func (g *GameState) MoveRight() {
	g.moveCursor(g.Cursor.Right(1))
}

func (g *GameState) MoveUp() {
	g.moveCursor(g.Cursor.Up(1))
}

func (g *GameState) MoveDown() {
	g.moveCursor(g.Cursor.Down(1))
}

func (g *GameState) Select() {
	if g.Selected == nil {
		g.selectPiece()
		return
	}

	if slices.Contains(g.ValidMoves, g.Cursor) {
		g.movePiece()
		return
	}

	if *g.Selected == g.Cursor {
		g.clearSelection()
	}
}

func (g *GameState) moveCursor(coor coordinate.Coordinate) {
	if insideBoard(coor) {
		g.Cursor = coor
	}
}

func (g *GameState) selectPiece() {
	piece, _ := g.Board.GetPiece(g.Cursor)
	if piece != nil && piece.Colour == g.PlayerTurn {
		g.Selected = ref(coordinate.Make(g.Cursor.X(), g.Cursor.Y()))
		g.ValidMoves = g.MoveFinder.FindMoves(g.Board[g.Cursor.Y()][g.Cursor.X()])
	}
}

func (g *GameState) movePiece() {
	selected, _ := g.Board.GetPiece(*g.Selected)
	destination, _ := g.Board.GetPiece(g.Cursor)

	g.Board.SetPiece(g.Cursor, selected)

	if g.castling(selected, destination) {
		g.Board.SetPiece(*g.Selected, destination)
	} else {
		g.Board.SetPiece(*g.Selected, nil)
	}

	if g.queening(selected) {
		selected.Character = character.Queen
	}

	g.clearSelection()

	if g.PlayerTurn == gamecolour.White {
		g.PlayerTurn = gamecolour.Black
	} else {
		g.PlayerTurn = gamecolour.White
	}
}

func (g *GameState) clearSelection() {
	g.ValidMoves = nil
	g.Selected = nil
}

func (g *GameState) castling(selected, destination *board.Piece) bool {
	return destination != nil &&
		((selected.Character == character.Rook && destination.Character == character.King) ||
			(selected.Character == character.King && destination.Character == character.Rook))
}

func (g *GameState) queening(selected *board.Piece) bool {
	return selected.Character == character.Pawn && (g.Cursor.Y() == 0 || g.Cursor.Y() == 7)
}

func insideBoard(coor coordinate.Coordinate) bool {
	return coor.X() >= 0 && coor.X() <= 7 && coor.Y() >= 0 && coor.Y() <= 7
}

func ref[T any](val T) *T {
	return &val
}
