package app

import (
	"fmt"
	"github.com/digitalsquid7/cli-chess/internal/ansi"
	"github.com/digitalsquid7/cli-chess/internal/chess"
	"github.com/digitalsquid7/cli-chess/internal/mainmenu"
	term "github.com/nsf/termbox-go"
)

type App struct {
	mainMenu  *mainmenu.MainMenu
	chessGame *chess.Game
}

func New() *App {
	return &App{
		mainMenu:  mainmenu.New(),
		chessGame: chess.NewGame(),
	}
}

func (a *App) Run() error {
	fmt.Print(ansi.EnableAlternativeScreen)
	defer fmt.Print(ansi.DisableAlternativeScreen)

	err := term.Init()
	if err != nil {
		return fmt.Errorf("term.Init: %w", err)
	}
	defer term.Close()

	option, err := a.mainMenu.SelectOption()
	if err != nil {
		return err
	}

	if option == mainmenu.Exit {
		return nil
	}

	return a.chessGame.Play()
}
