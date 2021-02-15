package main

import (
	"fmt"
)

const (
	CURSOR_POSITION_TEMPLATE = "\x1B[%d;%df"
	CURSOR_DOWN_TEMPLATE     = "\x1B[%dB"
	FG_COLOR_TEMPLATE        = "\x1B[%dm%s\x1B[0m"

	CLEAR_SCREEN = "\x1B[2J"
	RESET_ATTRS  = "\x1B[m"
	CURSOR_HOME  = "\x1B[H"

	FG_BLACK   = 30
	FG_RED     = 31
	FG_GREEN   = 32
	FG_YELLOW  = 33
	FG_BLUE    = 34
	FG_MAGENTA = 35
	FG_CYAN    = 36
	FG_WHITE   = 37
)

func fgColor(color int, text string) string {
	if color < 30 || color > 37 {
		color = FG_WHITE
	}

	return fmt.Sprintf(FG_COLOR_TEMPLATE, color, text)
}

func setCursorPosition(x, y int) {
	fmt.Printf(
		fmt.Sprintf(CURSOR_POSITION_TEMPLATE, y, x),
	)
}

func cursorDown(y int) {
	fmt.Printf(CURSOR_DOWN_TEMPLATE, y)
}

func clearScreen() {
	fmt.Print(CLEAR_SCREEN)
}

func resetAttributes() {
	fmt.Println(RESET_ATTRS)
}

func cursorHome() {
	fmt.Println(CURSOR_HOME)
}
