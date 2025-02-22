package ansi

import "fmt"

type Colour string

const (
	Red           Colour = "\033[31m"
	Green         Colour = "\033[32m"
	Yellow        Colour = "\033[33m"
	Blue          Colour = "\033[34m"
	Magenta       Colour = "\033[35m"
	Cyan          Colour = "\033[36m"
	Gray          Colour = "\033[37m"
	White         Colour = "\033[97m"
	BrightRed     Colour = "\033[31;1m"
	BrightGreen   Colour = "\033[32;1m"
	BrightYellow  Colour = "\033[33;1m"
	BrightBlue    Colour = "\033[34;1m"
	BrightMagenta Colour = "\033[35;1m"
	BrightCyan    Colour = "\033[36;1m"
	BrightGray    Colour = "\033[37;1m"
	BrightWhite   Colour = "\033[97;1m"
)

const (
	EnableAlternativeScreen  = "\033[?1049h"
	DisableAlternativeScreen = "\033[?1049l"
	RefreshScreen            = "\u001B[2J\u001B[H"
)

func CreateANSIText(text string, color Colour) string {
	return fmt.Sprint(color, text, "\033[0m")
}
