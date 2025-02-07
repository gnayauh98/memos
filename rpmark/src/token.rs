#[derive(Debug)]
pub enum NodeType {
    BLOCK,
    INLINE,
}

#[derive(Debug, Clone)]
pub struct Range {
    pub start: usize,
    pub end: usize,
}

#[derive(Debug)]
pub struct Node {
    pub offset: usize,
    pub range: Range,
    pub node_type: NodeType,
}
