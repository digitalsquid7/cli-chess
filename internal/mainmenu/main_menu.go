package mainmenu

import (
	"fmt"
	"github.com/digitalsquid7/cli-chess/internal/ansi"
	term "github.com/nsf/termbox-go"
	"slices"
	"time"
)

const ChessTitle = `
 ██████╗██╗  ██╗███████╗███████╗███████╗
██╔════╝██║  ██║██╔════╝██╔════╝██╔════╝
██║     ███████║█████╗  ███████╗███████╗
██║     ██╔══██║██╔══╝  ╚════██║╚════██║
╚██████╗██║  ██║███████╗███████║███████║
 ╚═════╝╚═╝  ╚═╝╚══════╝╚══════╝╚══════╝

`

type Option string

const (
	LocalPlayerVSPlayer      Option = "Local PvP"
	LocalPlayerVSEnvironment Option = "Local PvE"
	Exit                     Option = "Exit"
)

type MainMenu struct {
	options   []Option
	selected  int
	confirmed bool
}

func New() *MainMenu {
	return &MainMenu{
		options: []Option{
			LocalPlayerVSPlayer,
			LocalPlayerVSEnvironment,
			Exit,
		},
	}
}

func (m *MainMenu) SelectOption() (Option, error) {
	for range time.Tick(time.Second / 15) {
		if m.confirmed {
			break
		}

		err := m.executeCommands()
		if err != nil {
			return Exit, err
		}

		if err = m.updateScreen(); err != nil {
			return Exit, err
		}
	}

	return m.options[m.selected], nil
}

func (m *MainMenu) executeCommands() error {
	event := term.PollEvent()

	if event.Type == term.EventError {
		return event.Err
	}

	if event.Type != term.EventKey {
		return nil
	}

	switch event.Key {
	case term.KeyArrowUp:
		if m.selected != 0 {
			m.selected--
		}
	case term.KeyArrowDown:
		if m.selected < len(m.options)-1 {
			m.selected++
		}
	case term.KeyEnter:
		m.confirmed = true
	case term.KeyCtrlC:
		m.selected = slices.Index(m.options, Exit)
		m.confirmed = true
	default:
	}

	return term.Sync()
}

func (m *MainMenu) updateScreen() error {
	output := ansi.RefreshScreen + ChessTitle

	for _, option := range m.options {
		if option == m.options[m.selected] {
			output += fmt.Sprint(ansi.BackgroundYellow, string(option), ansi.Reset) + "\n"
		} else {
			output += string(option) + "\n"
		}
	}

	_, err := fmt.Print(output)
	return err
}
