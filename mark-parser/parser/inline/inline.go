package inline

import (
	"github.com/anqzi/mark-parser/parser/token"
)

type Matcher func([]byte) (Indexes, bool)

type Inline struct {
	Type InlineType
	Tag  string
	Matcher
}

type InlineType int

type Indexes struct {
	Indexes []int
	Type    InlineType
	Matches []int
}

const (
	Code     InlineType = 1
	Blob     InlineType = 2
	Tag      InlineType = 3
	Link     InlineType = 4
	MemoLink InlineType = 5
)

func InlineParse(text []byte, tokens []token.Token, blockStartIndex int, parsers []Inline) []token.Token {

	startIndex := 0

	index := 0

	for index < len(text) {
		flag := false
		for _, parser := range parsers {
			if indexes, ok := parser.Matcher(text[index:]); ok {
				endIndex := index
				// 添加一个开始判断，如果上一个是匹配的
				if startIndex != endIndex {
					tokens = append(tokens, token.Token{
						Type:            "inline",
						Tag:             0,
						BlockStartIndex: blockStartIndex,
						Text:            []int{startIndex, endIndex}, // 在文本内部的偏移
						Children:        []token.Token{},
					})
				}
				tokens = append(tokens, token.Token{
					Type:            "inline",
					Tag:             int(indexes.Type),
					BlockStartIndex: blockStartIndex,
					Text:            []int{index + indexes.Indexes[2], index + indexes.Indexes[3]},
					Children:        []token.Token{},
					Matches:         indexes.Matches,
				})
				startIndex = index + indexes.Indexes[1]
				index += indexes.Indexes[1]
				flag = true
				break
			}
		}
		if !flag {
			index += 1
		}
	}
	if startIndex != len(text) {
		// fmt.Printf("%s \n", text[startIndex:])
		tokens = append(tokens, token.Token{
			Type:            "inline",
			Tag:             0,
			BlockStartIndex: blockStartIndex,
			Text:            []int{startIndex, len(text)},
			Children:        []token.Token{},
		})
	}

	return tokens
}
