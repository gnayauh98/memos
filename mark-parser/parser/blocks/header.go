package blocks

import (
	"regexp"
)

const (
	HeaderRegexp = `^(#+)\s(.*)`
)

func FindHeaderIndex(text []byte) (Indexes, bool) {
	re, _ := regexp.Compile(HeaderRegexp)

	matches := re.FindSubmatchIndex(text)

	if len(matches) == 6 {
		return Indexes{
			Indexes: matches,
			Matches: matches[2:],
			Type:    Header,
		}, true
	}

	return Indexes{}, false
}

func GetHeaderNum(text []byte, matches []int) int {
	return len(string(text[matches[2]:matches[3]]))
}

func GetHeaderContent(text []byte, matches []int) []byte {
	return text[matches[4]:matches[5]]
}
