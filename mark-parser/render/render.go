package render

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/alecthomas/chroma/v2/formatters/html"
	"github.com/alecthomas/chroma/v2/lexers"
	"github.com/alecthomas/chroma/v2/styles"
	"github.com/anqzi/mark-parser/parser/blocks"
	"github.com/anqzi/mark-parser/parser/inline"
	"github.com/anqzi/mark-parser/parser/token"
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
	case int(blocks.Quote):
		return "<div class=\"quote\">"
	case int(blocks.List):
		return "<li class=\"list-item\">"
	case int(blocks.TodoList):
		// 提取状态
		status := string(texts[token.BlockStartIndex+token.Matches[0] : token.BlockStartIndex+token.Matches[1]])
		if status == "X" {
			return "<li class=\"todo-item done\">"
		} else if status == "I" {
			return "<li class=\"todo-item doing\">"
		} else {
			return "<li class=\"todo-item\">"
		}
	case int(blocks.Code):
		// 获取语言
		lang := string(texts[token.BlockStartIndex+token.Matches[0] : token.BlockStartIndex+token.Matches[1]])
		text := texts[token.BlockStartIndex+token.Matches[2] : token.BlockStartIndex+token.Matches[3]]
		return fmt.Sprintf("<div><div class=\"top\"><span class=\"lang\">%s</span><span class=\"icon i-lucide:clipboard\"></span></div>%s", lang, RenderCodeHighlight(text, lang))
	case int(blocks.Image):
		// 提取宽度
		style := ""
		var width float64 = 0
		var height float64 = 0
		if token.Matches[6] != -1 && token.Matches[7] != -1 && token.Matches[8] != -1 && token.Matches[9] != -1 {
			widthStr := string(texts[token.BlockStartIndex+token.Matches[6] : token.BlockStartIndex+token.Matches[7]])
			heightStr := string(texts[token.BlockStartIndex+token.Matches[8] : token.BlockStartIndex+token.Matches[9]])
			width, _ = strconv.ParseFloat(widthStr, 64)
			height, _ = strconv.ParseFloat(heightStr, 64)
		}

		if width > 0 {
			style += fmt.Sprint(
				"width: ",
				width,
				"%;",
			)
			height = height * width
		}

		if height > 0 {
			style += fmt.Sprint(
				"aspect-ratio: ",
				fmt.Sprintf("%.0f/%.0f", width, height),
				";",
			)
		}

		return fmt.Sprintf(
			"<div class=\"img\"><img src=\"/api/resource/%s\" style=\"%s\" />",
			texts[token.BlockStartIndex+token.Matches[2]:token.BlockStartIndex+token.Matches[3]],
			style,
		)
	}
	return "<div>"
}

func RenderBlockEnd(texts []byte, token token.Token) string {
	switch token.Tag {
	case int(blocks.Header):
		headerLevel := len(string(texts[token.BlockStartIndex+token.Matches[0] : token.BlockStartIndex+token.Matches[1]]))
		return fmt.Sprintf("</h%d>", headerLevel)
	case int(blocks.Quote):
		return "</div>"
	case int(blocks.List):
		return "</li>"
	case int(blocks.TodoList):
		return "</li>"
	case int(blocks.Code):
		return "</div>"
	case int(blocks.Image):
		return "</div>"
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
	case int(inline.Link):
		link := string(texts[token.BlockStartIndex+token.Matches[2] : token.BlockStartIndex+token.Matches[3]])
		return fmt.Sprintf("<a class=\"outlink\" target=\"_blank\" href=\"%s\">%s</a>", link, text)
	case int(inline.MemoLink):
		link := string(texts[token.BlockStartIndex+token.Matches[0] : token.BlockStartIndex+token.Matches[1]])
		return fmt.Sprintf("<a class=\"memo-link\"><span class=\"memo-link-mark i-lucide:link\"></span>%s</a>", link)
	}

	return text
}

func RenderCodeHighlight(texts []byte, lang string) string {

	style := styles.Get("gruvbox-light")
	if style == nil {
		style = styles.Fallback
	}
	formatter := html.New(
		html.WithLineNumbers(true),
		html.TabWidth(2),
		// html.Standalone(true),
	)

	lexer := lexers.Get(lang)

	iterator, _ := lexer.Tokenise(nil, string(texts))

	var sb strings.Builder

	formatter.Format(&sb, style, iterator)

	return sb.String()
}
