package terminal

import (
	"golang.org/x/sys/unix"
	"os"
)

func GetTerminalDimensions() (int, int, error) {
	ws, err := unix.IoctlGetWinsize(int(os.Stdout.Fd()), unix.TIOCGWINSZ)
	if err != nil {
		return 0, 0, err
	}

	return int(ws.Col), int(ws.Row), nil
}
