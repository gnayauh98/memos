package blocks

import (
	"regexp"
)

const (
	TodoRegexp = `^-\s\[([XI ])\]\s(.*)`
)

func FindTodoItemIndex(text []byte) (Indexes, bool) {
	re, _ := regexp.Compile(TodoRegexp)

	matches := re.FindSubmatchIndex(text)

	if len(matches) == 6 {
		return Indexes{
			Indexes: matches,
			Matches: matches[2:],
			Type:    TodoList,
		}, true
	}
	return Indexes{}, false
}

func GetTodoStatus(text []byte, matches []int) bool {
	return string(text[matches[2]:matches[3]]) == "X"
}

func GetTodoContent(text []byte, matches []int) []byte {
	return text[matches[4]:matches[5]]
}
