package main

var (
	termWidth, termHeight int
)

func init() {
	termWidth, termHeight, _ = getTerminalDimensions()
}

func main() {
	cursorHome()
	clearScreen()
	resetAttributes()
	for {
		renderUsers()
		prompt()
	}
}
