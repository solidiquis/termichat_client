package terminal

import (
	"fmt"
)

const (
	CURSOR_POSITION_TEMPLATE = "\x1B[%d;%df"
	CURSOR_DOWN_TEMPLATE     = "\x1B[%dB"
	CURSOR_HOME              = "\x1B[H"

	ERASE_CURRENT_LINE = "\x1B[1J"

	CLEAR_SCREEN = "\x1B[2J"
	RESET_ATTRS  = "\x1B[m"

	FG_COLOR_TEMPLATE = "\x1B[%dm%s\x1B[0m"
	FG_BLACK          = 30
	FG_RED            = 31
	FG_GREEN          = 32
	FG_YELLOW         = 33
	FG_BLUE           = 34
	FG_MAGENTA        = 35
	FG_CYAN           = 36
	FG_WHITE          = 37
)

func FgColor(color int, text string) string {
	if color < 30 || color > 37 {
		color = FG_WHITE
	}

	return fmt.Sprintf(FG_COLOR_TEMPLATE, color, text)
}

func SetCursorPosition(x, y int) {
	fmt.Printf(
		fmt.Sprintf(CURSOR_POSITION_TEMPLATE, y, x),
	)
}

func CursorDown(y int) {
	fmt.Printf(CURSOR_DOWN_TEMPLATE, y)
}

func ClearScreen() {
	fmt.Print(CLEAR_SCREEN)
}

func ResetAttributes() {
	fmt.Println(RESET_ATTRS)
}

func CursorHome() {
	fmt.Println(CURSOR_HOME)
}

func EraseCurrentLine() {
	fmt.Print(ERASE_CURRENT_LINE)
}
