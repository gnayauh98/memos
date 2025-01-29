package inline

import "regexp"

const (
	TagRegexp = `^#(.*)`
)

func FindTagIndex(text []byte) (Indexes, bool) {
	re, _ := regexp.Compile(TagRegexp)

	matches := re.FindSubmatchIndex(text)

	if len(matches) == 4 {
		return Indexes{
			Indexes: matches,
			Type:    Tag,
		}, true
	}

	return Indexes{}, false
}

// func FindTagContent(text []byte, matches []int)
