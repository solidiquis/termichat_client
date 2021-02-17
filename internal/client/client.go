package client

import (
	win "github.com/solidiquis/termichat_client/internal/terminal"
	ui "github.com/solidiquis/termichat_client/internal/ui"
)

func TermichatClient() {
	termWidth, termHeight, _ := win.GetTerminalDimensions()
	win.CursorHome()
	win.ClearScreen()
	win.ResetAttributes()

	for {
		ui.RenderUsers(termWidth)
		ui.Prompt(termHeight)
	}
}
