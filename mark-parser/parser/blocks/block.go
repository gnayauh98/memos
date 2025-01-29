package blocks

import (
	"github.com/kehuay/mark-parser/parser/inline"
	"github.com/kehuay/mark-parser/parser/token"
)

type Matcher func([]byte) (Indexes, bool)

type BlockType int

const (
	Header BlockType = iota
	TodoList
	Paragraph
	List
	Code
)

type Block struct {
	Type BlockType `json:"type"`
	Tag  string    `json:"tag"`
	// 块级匹配器
	Matcher Matcher
}

type Indexes struct {
	Indexes []int // 块内容索引
	Matches []int // 捕获组内容索引
	Type    BlockType
}

func (indexes Indexes) GetNextContent(text []byte) []byte {
	if (indexes.Indexes[1] + 1) > len(text) {
		return []byte{}
	}
	return text[indexes.Indexes[1]+1:]
}

func (indexes Indexes) GetIndexesContent(text []byte) []byte {
	switch indexes.Type {
	case Header:
		return GetHeaderContent(text, indexes.Indexes)
	case TodoList:
		return GetTodoContent(text, indexes.Indexes)
	case Paragraph:
		return GetParagraphContent(text, indexes.Indexes)
	case List:
		return GetListContent(text, indexes.Indexes)
	case Code:
		return GetCodeContent(text, indexes.Indexes)
	}
	return []byte{}
}

func (indexes Indexes) GetContentStartIndex() int {
	switch indexes.Type {
	case Header:
		return indexes.Matches[2]
	case TodoList:
		return indexes.Matches[2]
	case Paragraph:
		return indexes.Matches[0]
	case List:
		return indexes.Matches[0]
	case Code:
		return indexes.Matches[2]
	}
	return 0
}

func BlockParse(
	text []byte,
	tokens []token.Token,
	parsers []Block,
	inlineParsers []inline.Inline,
) []token.Token {
	blockStartIndex := 0
	for {
		var indexes Indexes
		var ok bool
		for _, parser := range parsers {
			if indexes, ok = parser.Matcher(text); ok {
				// fmt.Println(indexes)
				inlineContent := indexes.GetIndexesContent(text)
				// fmt.Printf("%s\n", inlineContent)
				tokens = append(tokens, token.Token{
					Type:            "block-start",
					Tag:             int(indexes.Type),
					BlockStartIndex: blockStartIndex,
					Matches:         indexes.Matches,
					Text:            indexes.Indexes[:2], // 从块开始索引开始的下标
					Children:        []token.Token{},
				})
				tokens = inline.InlineParse(
					inlineContent,
					tokens,
					blockStartIndex+indexes.GetContentStartIndex(),
					inlineParsers,
				)
				tokens = append(tokens, token.Token{
					Type:            "block-end",
					Tag:             int(indexes.Type),
					BlockStartIndex: blockStartIndex,
					Matches:         indexes.Matches,
					Text:            indexes.Indexes[:2],
					Children:        []token.Token{},
				})
				break
			}
		}
		if ok {
			blockStartIndex += (indexes.Indexes[1] + 1)
			text = indexes.GetNextContent(text)
		} else {
			// 没有匹配任何内容
		}

		if len(text) == 0 {
			break
		}
	}

	return tokens
}

func GetMatchesContent(text []byte, matches []int) []byte {
	return text[matches[0]:matches[1]]
}

func GetNextContent(text []byte, matches []int) []byte {
	return text[matches[1]+1:]
}
