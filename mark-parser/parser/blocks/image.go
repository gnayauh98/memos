package blocks

import (
	"regexp"
)

const (
	ImageRegexp = `^!\[(.*)\]\(@([a-zA-Z0-9-]*)(\?w=(\d+)&h=(\d*.?\d+))?\)`
)

func FindImageIndex(texts []byte) (Indexes, bool) {
	re, _ := regexp.Compile(ImageRegexp)

	matches := re.FindSubmatchIndex(texts)

	// log.Println(matches)

	if len(matches) >= 6 {
		return Indexes{
			Indexes: matches,
			Matches: matches[2:],
			Type:    Image,
		}, true
	}

	return Indexes{}, false
}

func GetImageCaption(texts []byte, matches []int) []byte {
	return texts[matches[2]:matches[3]]
}

func GetImageId(texts []byte, matches []int) []byte {
	return texts[matches[4]:matches[5]]
}
