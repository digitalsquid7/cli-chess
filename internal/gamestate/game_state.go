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
	g.checkCastling()
	g.checkQueening()

	selected, _ := g.Board.GetPiece(*g.Selected)
	g.Board.SetPiece(g.Cursor, selected)
	g.Board.SetPiece(*g.Selected, nil)

	g.clearSelection()
	g.switchPlayerTurn()
}

func (g *GameState) clearSelection() {
	g.ValidMoves = nil
	g.Selected = nil
}

// checkCastling checks if the selected and destination pieces are a rook and king in a valid castling
// position and updates the board.
func (g *GameState) checkCastling() {
	selected, _ := g.Board.GetPiece(*g.Selected)
	destination, _ := g.Board.GetPiece(g.Cursor)

	if selected == nil || destination == nil ||
		!((selected.Character == character.Rook && destination.Character == character.King) ||
			(selected.Character == character.King && destination.Character == character.Rook)) {
		return
	}

	var rook, king *board.Piece

	if selected.Character == character.Rook {
		rook = selected
		king = destination
	} else {
		rook = destination
		king = selected
	}

	if rook.Coor.X() == 0 {
		g.Board.SetPiece(coordinate.Make(3, rook.Coor.Y()), rook)
		g.Board.SetPiece(coordinate.Make(2, rook.Coor.Y()), king)
	} else {
		g.Board.SetPiece(coordinate.Make(5, rook.Coor.Y()), rook)
		g.Board.SetPiece(coordinate.Make(6, rook.Coor.Y()), king)
	}

	g.Board.SetPiece(g.Cursor, nil)
	g.Board.SetPiece(*g.Selected, nil)
}

// checkQueening promotes a pawn to a queen if it reaches the opposing end of the board.
func (g *GameState) checkQueening() {
	selected, _ := g.Board.GetPiece(*g.Selected)
	if selected != nil && selected.Character == character.Pawn && (g.Cursor.Y() == 0 || g.Cursor.Y() == 7) {
		selected.Character = character.Queen
	}
}

func (g *GameState) switchPlayerTurn() {
	if g.PlayerTurn == gamecolour.White {
		g.PlayerTurn = gamecolour.Black
	} else {
		g.PlayerTurn = gamecolour.White
	}
}

func insideBoard(coor coordinate.Coordinate) bool {
	return coor.X() >= 0 && coor.X() <= 7 && coor.Y() >= 0 && coor.Y() <= 7
}

func ref[T any](val T) *T {
	return &val
}
