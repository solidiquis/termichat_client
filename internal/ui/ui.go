package ui

import (
	"bufio"
	"fmt"
	"os"
	"sort"

	win "github.com/solidiquis/termichat_client/internal/terminal"
)

var messages []string

// TODO: Cache
func RenderUsers(termWidth int) {
	users := []string{
		"Yuna",
		"Tidus",
		"Aaron",
		"Wakka",
		"Lulu",
		"Riku",
		"Kimahri",
	}

	sort.Slice(users, func(i, j int) bool {
		return len(users[i]) > len(users[j])
	})

	maxLength := func() int {
		var max int
		for _, i := range users {
			current := len(i)
			if max < current {
				max = current
			}
		}
		return max
	}()

	for i, name := range users {
		var status string

		if i%2 == 0 {
			status = win.FgColor(win.FG_CYAN, "\u25CF")
		} else {
			status = "\u25CB"
		}

		users[i] = fmt.Sprintf("%s - %s", status, name)
	}

	// offset by 4 for status circle
	offsetX := termWidth - maxLength - 4
	offsetY := 1

	win.SetCursorPosition(offsetX, 0)

	for _, user := range users {
		fmt.Print(user)
		offsetY++

		win.SetCursorPosition(offsetX, offsetY)
	}
	fmt.Println()
}

func Prompt(termHeight int) {
	win.SetCursorPosition(0, termHeight)
	fmt.Print(win.FgColor(win.FG_RED, " \u25BA "))
	r := bufio.NewReader(os.Stdin)
	text, _ := r.ReadString('\n')
	win.EraseCurrentLine()
	text = processEmojis(string(text))
	text = fmt.Sprintf("%s: %s", win.FgColor(win.FG_GREEN, "Yuna"), text)
	messages = append(messages, text)
	j := 0
	for i := len(messages) - 1; i >= 0; i-- {
		win.SetCursorPosition(0, termHeight-i)
		fmt.Print(messages[j])
		j++
	}
}
