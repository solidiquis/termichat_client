package structs

import (
	"sort"

	win "github.com/solidiquis/termichat_client/internal/terminal"
)

type User struct {
	Name   string
	Online bool
}

func SortUsers(users []*User) []*User {
	sort.Slice(users, func(i, j int) bool {
		return len(users[i].Name) > len(users[j].Name)
	})

	return users
}

func (u *User) OnlineStatus() string {
	if u.Online {
		return win.FgColor(win.FG_CYAN, "\u25CF")
	}

	return "\u25CB"
}
