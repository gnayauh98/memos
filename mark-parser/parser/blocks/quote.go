package blocks

import (
	"regexp"
)

const (
	QuoteRegexp = `^> (.*)`
)

func FindQuoteIndex(text []byte) (Indexes, bool) {
	re, _ := regexp.Compile(QuoteRegexp)

	matches := re.FindSubmatchIndex(text)

	if len(matches) == 4 {
		return Indexes{
			Indexes: matches,
			Matches: matches[2:],
			Type:    Quote,
		}, true
	}
	return Indexes{}, false
}

func GetQuoteContent(text []byte, matches []int) []byte {
	return text[matches[2]:matches[3]]
}
