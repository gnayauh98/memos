package blocks

import "regexp"

const (
	ParagraphRegexp = `^(.*)`
)

func FindParagraphIndex(text []byte) (Indexes, bool) {
	re, _ := regexp.Compile(ParagraphRegexp)
	matches := re.FindSubmatchIndex(text)

	if len(matches) == 4 {
		return Indexes{
			Indexes: matches,
			Matches: matches[2:],
			Type:    Paragraph,
		}, true
	}
	return Indexes{}, false
}

func GetParagraphContent(text []byte, matches []int) []byte {
	return text[matches[2]:matches[3]]
}
