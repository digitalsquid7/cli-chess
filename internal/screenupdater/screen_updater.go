package screenupdater

import (
	"bytes"
	"fmt"
	"github.com/digitalsquid7/cli-chess/internal/ansi"
	"github.com/digitalsquid7/cli-chess/internal/gamestate"
	"github.com/digitalsquid7/cli-chess/internal/gamestate/coordinate"
	"github.com/digitalsquid7/cli-chess/internal/gamestate/gamecolour"
	"slices"
)

type ScreenUpdater struct {
	gameState *gamestate.GameState
}

func New(gameState *gamestate.GameState) *ScreenUpdater {
	return &ScreenUpdater{gameState: gameState}
}

func (s *ScreenUpdater) Update() error {
	board := s.createBoard()

	if err := s.printGame(board); err != nil {
		return fmt.Errorf("printGame: %w", err)
	}

	return nil
}

func (s *ScreenUpdater) createBoard() [][]string {
	board := [][]string{
		{"╔", "═══", "╤", "═══", "╤", "═══", "╤", "═══", "╤", "═══", "╤", "═══", "╤", "═══", "╤", "═══", "╗", "╮"},
		{"║", "   ", "│", "   ", "│", "   ", "│", "   ", "│", "   ", "│", "   ", "│", "   ", "│", "   ", "║", "8"},
		{"╟", "═══", "┼", "═══", "┼", "═══", "┼", "═══", "┼", "═══", "┼", "═══", "┼", "═══", "┼", "═══", "╢", "┊"},
		{"║", "   ", "│", "   ", "│", "   ", "│", "   ", "│", "   ", "│", "   ", "│", "   ", "│", "   ", "║", "7"},
		{"╟", "═══", "┼", "═══", "┼", "═══", "┼", "═══", "┼", "═══", "┼", "═══", "┼", "═══", "┼", "═══", "╢", "┊"},
		{"║", "   ", "│", "   ", "│", "   ", "│", "   ", "│", "   ", "│", "   ", "│", "   ", "│", "   ", "║", "6"},
		{"╟", "═══", "┼", "═══", "┼", "═══", "┼", "═══", "┼", "═══", "┼", "═══", "┼", "═══", "┼", "═══", "╢", "┊"},
		{"║", "   ", "│", "   ", "│", "   ", "│", "   ", "│", "   ", "│", "   ", "│", "   ", "│", "   ", "║", "5"},
		{"╟", "═══", "┼", "═══", "┼", "═══", "┼", "═══", "┼", "═══", "┼", "═══", "┼", "═══", "┼", "═══", "╢", "┊"},
		{"║", "   ", "│", "   ", "│", "   ", "│", "   ", "│", "   ", "│", "   ", "│", "   ", "│", "   ", "║", "4"},
		{"╟", "═══", "┼", "═══", "┼", "═══", "┼", "═══", "┼", "═══", "┼", "═══", "┼", "═══", "┼", "═══", "╢", "┊"},
		{"║", "   ", "│", "   ", "│", "   ", "│", "   ", "│", "   ", "│", "   ", "│", "   ", "│", "   ", "║", "3"},
		{"╟", "═══", "┼", "═══", "┼", "═══", "┼", "═══", "┼", "═══", "┼", "═══", "┼", "═══", "┼", "═══", "╢", "┊"},
		{"║", "   ", "│", "   ", "│", "   ", "│", "   ", "│", "   ", "│", "   ", "│", "   ", "│", "   ", "║", "2"},
		{"╟", "═══", "┼", "═══", "┼", "═══", "┼", "═══", "┼", "═══", "┼", "═══", "┼", "═══", "┼", "═══", "╢", "┊"},
		{"║", "   ", "│", "   ", "│", "   ", "│", "   ", "│", "   ", "│", "   ", "│", "   ", "│", "   ", "║", "1"},
		{"╚", "═══", "╧", "═══", "╧", "═══", "╧", "═══", "╧", "═══", "╧", "═══", "╧", "═══", "╧", "═══", "╝", "┊"},
		{"╰", "┈a┈", "│", "┈b┈", "│", "┈c┈", "│", "┈d┈", "│", "┈e┈", "│", "┈f┈", "│", "┈g┈", "│", "┈h┈", "┈", "╯"},
	}

	for y := range s.gameState.Board {
		for x := range s.gameState.Board[y] {
			board[y*2+1][x*2+1] = fmt.Sprint(s.getBackgroundColour(x, y), s.getForegroundColour(x, y), s.getBold(x, y), s.getCharacter(x, y), ansi.Reset)
		}
	}

	return board
}

func (s *ScreenUpdater) printGame(board [][]string) error {
	buffer := bytes.Buffer{}

	_, err := fmt.Fprint(&buffer, ansi.RefreshScreen)
	if err != nil {
		return err
	}

	for i := range board {
		for j := range board[i] {
			if _, err = fmt.Fprint(&buffer, board[i][j]); err != nil {
				return err
			}
		}

		if _, err = fmt.Fprint(&buffer, "\n"); err != nil {
			return err
		}
	}

	_, err = fmt.Print(buffer.String())
	return err
}

func (s *ScreenUpdater) getForegroundColour(x, y int) gamecolour.Colour {
	if s.gameState.Cursor.Y() == y && s.gameState.Cursor.X() == x && s.gameState.Board[y][x] != nil {
		return ansi.Yellow
	}

	if slices.Contains(s.gameState.ValidMoves, coordinate.Make(x, y)) {
		return ansi.Green
	}

	if s.gameState.Board[y][x] != nil {
		if s.gameState.Board[y][x].Colour == gamecolour.Black {
			return ansi.Black
		}
		return ansi.DarkGrey
	}

	return ""
}

func (s *ScreenUpdater) getBackgroundColour(x, y int) gamecolour.Colour {
	if s.gameState.Cursor.Y() == y && s.gameState.Cursor.X() == x && s.gameState.Board[y][x] == nil {
		return ansi.BackgroundYellow
	}

	return ""
}

func (s *ScreenUpdater) getCharacter(x, y int) string {
	if s.gameState.Board[y][x] != nil {
		return " " + string(s.gameState.Board[y][x].Character) + " "
	}

	if slices.Contains(s.gameState.ValidMoves, coordinate.Make(x, y)) {
		return "░░░"
	}

	return "   "
}

func (s *ScreenUpdater) getBold(x, y int) string {
	if s.gameState.Selected != nil && s.gameState.Selected.Y() == y && s.gameState.Selected.X() == x {
		return ansi.Bold
	}

	return ""
}
