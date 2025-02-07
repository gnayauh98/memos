mod parser;
mod token;

fn main() {
    let texts = "
- [DONE] 待办事项

```go
package main

func main()
```

- [DOING] 正在进行中
";

    let nodes = parser::parser(texts);

    for node in nodes.iter() {
        println!(
            "{} {:?}",
            &texts[(node.range.start + node.offset)..(node.range.end + node.offset)],
            node.node_type
        );
    }

    // println!("{:?}", nodes);
}
