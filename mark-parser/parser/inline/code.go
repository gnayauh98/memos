package inline

import "regexp"

const (
	CodeRegexp = "^`(.*?)`"
)

func FindCodeIndex(text []byte) (Indexes, bool) {
	re, _ := regexp.Compile(CodeRegexp)

	matches := re.FindSubmatchIndex(text)

	if len(matches) == 4 {
		return Indexes{
			Indexes: matches,
			Type:    Code,
		}, true
	}

	return Indexes{}, false
}
