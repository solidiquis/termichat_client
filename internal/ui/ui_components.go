package ui

import (
	"bufio"
	"fmt"
	"os"

	st "github.com/solidiquis/termichat_client/internal/structs"
	win "github.com/solidiquis/termichat_client/internal/terminal"
)

func UserList(termWidth int, users []*st.User) {
	users = st.SortUsers(users)

	maxLength := 0
	for _, user := range users {
		if current := len(user.Name); maxLength < current {
			maxLength = current
		}
	}

	// offset by 4 for status circle
	offsetX := termWidth - maxLength - 4
	offsetY := 1

	win.SetCursorPosition(offsetX, 0)

	for _, user := range users {
		fmt.Print(fmt.Sprintf(
			"%s - %s",
			(*user).OnlineStatus(),
			(*user).Name,
		))
		offsetY++

		win.SetCursorPosition(offsetX, offsetY)
	}
	fmt.Println()
}

func MessageContainer(termHeight int, messages []string) {
	for i, j := len(messages)-1, 0; i >= 0; i, j = i-1, j+1 {
		win.SetCursorPosition(0, termHeight-i)
		fmt.Print(messages[j])
	}
}

func Prompt(termHeight int, chatroom *st.Chatroom) {
	win.SetCursorPosition(0, termHeight)
	fmt.Print(win.FgColor(win.FG_RED, " \u25BA "))
	r := bufio.NewReader(os.Stdin)
	text, _ := r.ReadString('\n')
	win.EraseCurrentLine()
	text = processEmojis(string(text))
	text = fmt.Sprintf("%s: %s", win.FgColor(win.FG_GREEN, "Yuna"), text)
	chatroom.PushMessage(text)
	MessageContainer(termHeight, chatroom.Messages)
}
