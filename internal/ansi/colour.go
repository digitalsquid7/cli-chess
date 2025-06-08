package ansi

type Colour string

const (
	Black            = ""
	Red              = "\033[31m"
	Green            = "\033[32m"
	BackgroundGreen  = "\033[42m"
	Yellow           = "\033[33m"
	BackgroundYellow = "\033[43m"
	Blue             = "\033[34m"
	BackgroundBlue   = "\033[44m"
	Magenta          = "\033[35m"
	Cyan             = "\033[36m"
	Gray             = "\033[37m"
	White            = "\033[97m"
	DarkGrey         = "\033[90m"

	BrightRed     = "\033[31;1m"
	BrightGreen   = "\033[32;1m"
	BrightYellow  = "\033[33;1m"
	BrightBlue    = "\033[34;1m"
	BrightMagenta = "\033[35;1m"
	BrightCyan    = "\033[36;1m"
	BrightGray    = "\033[37;1m"
	BrightWhite   = "\033[97;1m"
)

const (
	EnableAlternativeScreen  = "\033[?1049h"
	DisableAlternativeScreen = "\033[?1049l"
	RefreshScreen            = "\u001B[2J\u001B[H"
	Reset                    = "\033[0m"
	Bold                     = "\033[1m"
)
