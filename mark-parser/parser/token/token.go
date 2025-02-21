package token

type Token struct {
	Type            string
	Tag             int
	BlockStartIndex int
	Text            []int
	Matches         []int
	Children        []Token
}

func GetTags(texts []byte, tokens []Token) []string {

	tags := make([]string, 0)
	for _, token := range tokens {
		startIndex := token.BlockStartIndex + token.Text[0]
		if token.Tag == 3 {
			tags = append(tags, string(texts[startIndex+token.Matches[0]:startIndex+token.Matches[1]]))
		}
	}

	return tags
}
