package inline

import (
	"regexp"
)

const (
	LinkRegexp = `^\[(.*)\]\((.*)\)`
)

func FindLinkIndex(text []byte) (Indexes, bool) {
	re, _ := regexp.Compile(LinkRegexp)

	// log.Printf("%s", text)

	matches := re.FindSubmatchIndex(text)

	if len(matches) == 6 {
		return Indexes{
			Indexes: matches,
			Type:    Link,
			Matches: matches[2:],
		}, true
	}

	return Indexes{}, false
}
