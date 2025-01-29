package render

import (
	"fmt"

	"github.com/kehuay/mark-parser/parser/blocks"
	"github.com/kehuay/mark-parser/parser/inline"
	"github.com/kehuay/mark-parser/parser/token"
)

func RenderToHtml(texts []byte, tokens []token.Token) string {

	result := ""

	for _, token := range tokens {
		switch token.Type {
		case "block-start":
			result += RenderBlockStart(texts, token)
		case "block-end":
			result += RenderBlockEnd(texts, token)
		case "inline":
			result += RenderInline(texts, token)
		}
	}

	return result
}

func RenderBlockStart(texts []byte, token token.Token) string {
	switch token.Tag {
	case int(blocks.Header):
		headerLevel := len(string(texts[token.BlockStartIndex+token.Matches[0] : token.BlockStartIndex+token.Matches[1]]))
		return fmt.Sprintf("<h%d>", headerLevel)
	case int(blocks.List):
		return "<li class=\"list-item\">"
	case int(blocks.TodoList):
		return "<li class=\"todo-item\">"
	case int(blocks.Code):
		// 获取语言
		lang := string(texts[token.BlockStartIndex+token.Matches[0] : token.BlockStartIndex+token.Matches[1]])
		return fmt.Sprintf("<pre class=\"%s\"><code>", lang)
	}
	return "<div>"
}

func RenderBlockEnd(texts []byte, token token.Token) string {
	switch token.Tag {
	case int(blocks.Header):
		headerLevel := len(string(texts[token.BlockStartIndex+token.Matches[0] : token.BlockStartIndex+token.Matches[1]]))
		return fmt.Sprintf("</h%d>", headerLevel)
	case int(blocks.List):
		return "</li>"
	case int(blocks.TodoList):
		return "</li>"
	case int(blocks.Code):
		return "</code></pre>"
	}
	return "</div>"
}

func RenderInline(texts []byte, token token.Token) string {
	text := string(texts[token.BlockStartIndex+token.Text[0] : token.BlockStartIndex+token.Text[1]])
	switch token.Tag {
	case int(inline.Blob):
		return fmt.Sprintf("<strong>%s</strong>", text)
	case int(inline.Code):
		return fmt.Sprintf("<span class=\"inline-code\">%s</span>", text)
	case int(inline.Tag):
		return fmt.Sprintf("<span class=\"tag\">%s</span>", text)
	}

	return text
}
