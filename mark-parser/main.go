package main

import (
	"fmt"

	"github.com/kehuay/mark-parser/parser"
	"github.com/kehuay/mark-parser/render"
)

func main() {

	text := `### 个人'SOP'shadj实际~达成共识~结果适当补偿技术差距时都会保持身材不好就很多首班车时间不长的时间

- [ ] 晚上备好隔天的衣服
- [X] 起床先'喝一杯温热水'
- [ ] 睡够8小时
- [ ] 早起列当天Todo List

✅ 喝水喝水

'''go
package main

import "fmt"

func main(){
    fmt.Println("Hello, Huy!")
}
'''

- 1
- 2
- 3

这是一个'Markdown'文本解析器
`

	bytes := []byte(text)
	tokens := parser.Parser(bytes)

	html := render.RenderToHtml(bytes, tokens)

	fmt.Println(html)
}
