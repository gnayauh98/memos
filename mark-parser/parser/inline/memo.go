package inline

import "regexp"

const (
	MemoRegexp = `^@\[(.*)\]`
)

func FindMemoIndex(texts []byte) (Indexes, bool) {
	re, _ := regexp.Compile(MemoRegexp)

	matches := re.FindSubmatchIndex(texts)

	if len(matches) == 4 {
		return Indexes{
			Indexes: matches,
			Type:    MemoLink,
			Matches: matches[2:],
		}, true
	}
	return Indexes{}, false
}
