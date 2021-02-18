package client

import (
	st "github.com/solidiquis/termichat_client/internal/structs"
	win "github.com/solidiquis/termichat_client/internal/terminal"
	ui "github.com/solidiquis/termichat_client/internal/ui"
)

// Dummy data for now
var (
	users    []*st.User
	chatroom *st.Chatroom
	messages []string
)

func init() {
	users = []*st.User{
		&st.User{Name: "Yuna", Online: true},
		&st.User{Name: "Tidus", Online: false},
		&st.User{Name: "Aaron", Online: true},
		&st.User{Name: "Wakka", Online: false},
		&st.User{Name: "Lulu", Online: true},
		&st.User{Name: "Riku", Online: false},
		&st.User{Name: "Kimahri", Online: true},
	}

	chatroom = &st.Chatroom{Messages: messages, Users: users}
}

// end dummy data

func TermichatClient() {
	termWidth, termHeight, _ := win.GetTerminalDimensions()
	win.CursorHome()
	win.ClearScreen()
	win.ResetAttributes()

	for {
		ui.UserList(termWidth, users)
		ui.Prompt(termHeight, chatroom)
	}
}
