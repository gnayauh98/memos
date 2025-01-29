package blocks

import "regexp"

const (
	ListRegexp = `^-\s(.*)`
)

func FindListIndex(text []byte) (Indexes, bool) {
	re, _ := regexp.Compile(ListRegexp)

	matches := re.FindSubmatchIndex(text)

	if len(matches) == 4 {
		return Indexes{
			Indexes: matches,
			Matches: matches[2:],
			Type:    List,
		}, true
	}

	return Indexes{}, false
}

func GetListContent(text []byte, matches []int) []byte {
	return text[matches[2]:matches[3]]
}
