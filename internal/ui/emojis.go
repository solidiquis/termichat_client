package ui

import (
	"regexp"
	"strings"
)

const (
	EMOJI_REGEX = `:[a-z_]+:`
)

var emojisMap = make(map[string]string)

func init() {
	emojisMap = map[string]string{
		"angry": "\U0001f620",
		"smile": "\U0001f604",
	}
}

func processEmojis(msg string) string {

	re, _ := regexp.Compile(EMOJI_REGEX)
	emojis := re.FindAllString(msg, -1)

	// No emojis
	if len(emojis) == 0 {
		return msg
	}

	emojiUnicodes := make([]string, len(emojis))
	for i, j := range emojis {
		s := strings.Trim(j, ":")
		if unicode, ok := emojisMap[s]; !ok {
			emojiUnicodes[i] = j
		} else {
			emojiUnicodes[i] = unicode
		}
	}

	for i := 0; i < len(emojis); i++ {
		if emojis[i] != emojiUnicodes[i] {
			msg = strings.Replace(msg, emojis[i], emojiUnicodes[i], -1)
		}
	}

	return msg
}
