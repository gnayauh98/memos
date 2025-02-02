package blocks

import "regexp"

const (
	CodeRegexp = "^```(.*)\n((?s).*)\n```"
)

func FindCodeIndex(text []byte) (Indexes, bool) {
	re, _ := regexp.Compile(CodeRegexp)

	matches := re.FindSubmatchIndex(text)

	if len(matches) == 6 {
		return Indexes{
			Indexes: matches,
			Matches: matches[2:],
			Type:    Code,
		}, true
	}

	return Indexes{}, false
}

func GetCodeLanguage(text []byte, matches []int) []byte {
	return text[matches[2]:matches[3]]
}

func GetCodeContent(text []byte, matches []int) []byte {
	return text[matches[4]:matches[5]]
}
