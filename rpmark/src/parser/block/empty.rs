use crate::{
    parser::{Match, Matcher},
    token::{Node, NodeType},
};

pub struct EmptyLineBlock {}

impl Matcher for EmptyLineBlock {
    const REGEXP: &str = r"^(\s*)[\s]";
}

impl EmptyLineBlock {
    pub fn tokens(&self, _match: &Match, nodes: &mut Vec<Node>) {
        nodes.push(Node {
            range: _match.range.clone(),
            node_type: NodeType::BLOCK,
            offset: _match.offset,
        });
    }
}
