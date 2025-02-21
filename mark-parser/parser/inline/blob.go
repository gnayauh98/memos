package inline

import "regexp"

const (
	BlobRegexp = `^~(.*?)\~`
)

func FindBlobIndex(text []byte) (Indexes, bool) {
	re, _ := regexp.Compile(BlobRegexp)

	matches := re.FindSubmatchIndex(text)

	if len(matches) == 4 {
		return Indexes{
			Indexes: matches,
			Type:    Blob,
			Matches: matches[2:],
		}, true
	}

	return Indexes{}, false
}
