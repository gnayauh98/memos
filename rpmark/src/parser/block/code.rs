use crate::{
    parser::{Match, Matcher},
    token::{self, Node, NodeType},
};

pub struct CodeBlock {}

impl Matcher for CodeBlock {
    const REGEXP: &str = r"^```(.*)\s((?s).*)\s```[\s]";
}

impl CodeBlock {
    pub fn tokens(&self, _match: &Match, nodes: &mut Vec<Node>) {
        nodes.push(Node {
            range: _match.range.clone(),
            node_type: NodeType::BLOCK,
            offset: _match.offset,
        });
        nodes.push(Node {
            range: token::Range {
                start: _match.captures[0],
                end: _match.captures[1],
            },
            offset: _match.offset,
            node_type: NodeType::INLINE,
        });
        nodes.push(Node {
            range: token::Range {
                start: _match.captures[2],
                end: _match.captures[3],
            },
            offset: _match.offset,
            node_type: NodeType::INLINE,
        });
    }
}
