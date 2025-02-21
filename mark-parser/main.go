package main

import (
	"fmt"

	"github.com/anqzi/mark-parser/parser"
	"github.com/anqzi/mark-parser/render"
)

func main() {

	text := "下午`UI`op\n\nOP"

	bytes := []byte(text)
	tokens := parser.Parser(bytes)

	fmt.Println(tokens)
	html := render.RenderToHtml(bytes, tokens)

	fmt.Println(html)
}
