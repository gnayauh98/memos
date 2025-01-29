package parser

import (
	"github.com/kehuay/mark-parser/parser/blocks"
	"github.com/kehuay/mark-parser/parser/inline"
	"github.com/kehuay/mark-parser/parser/token"
)

func Parser(text []byte) []token.Token {

	blocksParser := []blocks.Block{
		{
			Type:    blocks.Header,
			Tag:     "h%d",
			Matcher: blocks.FindHeaderIndex,
		},
		{
			Type:    blocks.TodoList,
			Tag:     "todo",
			Matcher: blocks.FindTodoItemIndex,
		},
		{
			Type:    blocks.List,
			Tag:     "list",
			Matcher: blocks.FindListIndex,
		},
		{
			Type:    blocks.Code,
			Tag:     "code",
			Matcher: blocks.FindCodeIndex,
		},
		{
			Type:    blocks.Paragraph,
			Tag:     "p",
			Matcher: blocks.FindParagraphIndex,
		},
	}

	inlineParser := []inline.Inline{
		{
			Type:    inline.Code,
			Tag:     "code",
			Matcher: inline.FindCodeIndex,
		},
		{
			Type:    inline.Blob,
			Tag:     "strong",
			Matcher: inline.FindBlobIndex,
		},
		{
			Type:    inline.Tag,
			Tag:     "tag",
			Matcher: inline.FindTagIndex,
		},
	}

	tokens := []token.Token{}

	tokens = blocks.BlockParse(text, tokens, blocksParser, inlineParser)

	return tokens

	// for _, token := range tokens {
	// 	if token.Type == "block-start" {
	// 		fmt.Printf("2=====: %s\n", token.Tag)
	// 	}
	// 	if token.Type == "block-end" {
	// 		fmt.Printf("=====2: %s\n", token.Tag)
	// 	}
	// 	if token.Type == "inline" {
	// 		fmt.Printf("1: %s\n", textBytes[token.BlockStartIndex+token.Text[0]:token.BlockStartIndex+token.Text[1]])
	// 	}

	// }
}
