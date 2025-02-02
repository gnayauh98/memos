package main

import (
	"fmt"

	"github.com/anqzi/mark-parser/parser"
	"github.com/anqzi/mark-parser/render"
)

func main() {

	text := "```go\npackage main\nfunc main(){\nprintln(\"hello, anqzi!\")\n```"

	bytes := []byte(text)
	tokens := parser.Parser(bytes)

	html := render.RenderToHtml(bytes, tokens)

	fmt.Println(html)
}
