package main

import (
	"fmt"

	"github.com/kehuay/mark-parser/parser"
	"github.com/kehuay/mark-parser/render"
)

func main() {

	text := "[使用Golang实现的代码高亮库](https://github.com/alecthomas/chroma)\n"

	bytes := []byte(text)
	tokens := parser.Parser(bytes)

	html := render.RenderToHtml(bytes, tokens)

	fmt.Println(html)
}
