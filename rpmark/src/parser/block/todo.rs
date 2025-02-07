use crate::{
    parser::{Match, Matcher},
    token::{Node, NodeType, Range},
};

pub struct TodoItemBlock {}

impl Matcher for TodoItemBlock {
    const REGEXP: &str = r"^- \[(.*)\] (.*)[\s]";
}

impl TodoItemBlock {
    pub fn tokens(&self, _match: &Match, nodes: &mut Vec<Node>) {
        nodes.push(Node {
            range: _match.range.clone(),
            node_type: NodeType::BLOCK,
            offset: _match.offset,
        });
        nodes.push(Node {
            range: Range {
                start: _match.captures[0],
                end: _match.captures[1],
            },
            offset: _match.offset,
            node_type: NodeType::INLINE,
        });
        nodes.push(Node {
            range: Range {
                start: _match.captures[2],
                end: _match.captures[3],
            },
            offset: _match.offset,
            node_type: NodeType::INLINE,
        });
    }
}
