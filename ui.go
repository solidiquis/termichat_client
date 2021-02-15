package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// TODO: Cache
func renderUsers() {
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
			status = fgColor(FG_CYAN, "\u25CF")
		} else {
			status = "\u25CB"
		}

		users[i] = fmt.Sprintf("%s - %s", status, name)
	}

	// offset by 4 for status circle
	offsetX := termWidth - maxLength - 4
	offsetY := 1

	setCursorPosition(offsetX, 0)

	for _, user := range users {
		fmt.Print(user)
		offsetY++

		setCursorPosition(offsetX, offsetY)
	}
	fmt.Println()
}

func prompt() {
	//var messages []string
	setCursorPosition(0, termHeight)
	fmt.Print(fgColor(FG_RED, " \u25BA "))
	r := bufio.NewReader(os.Stdin)
	text, _ := r.ReadString('\n')
	text = processEmojis(string(text))
	text = fmt.Sprintf("%s: %s", fgColor(FG_GREEN, "Yuna"), text)
	//messages = append(messages, text)

	setCursorPosition(0, termHeight-1) // here broken
	fmt.Printf(text)
}
