use regex::Regex;

use crate::token::{self, Node};

mod block;
mod inline;

pub trait Matcher {
    const REGEXP: &str;
    fn matcher(&self, texts: &str, start_index: usize) -> Option<Match> {
        let re = Regex::new(Self::REGEXP).unwrap();

        let caps = re.captures(&texts[start_index..]);

        if let Some(caps) = caps {
            let mut ids: Vec<usize> = vec![];

            for cap in caps.iter() {
                let cap = cap.unwrap();
                ids.push(cap.start());
                ids.push(cap.end());
            }

            return Some(Match {
                range: token::Range {
                    start: ids[0],
                    end: ids[1],
                },
                offset: start_index,
                captures: ids[2..].to_vec(),
            });
        }

        None
    }
}

#[derive(Clone)]
pub struct Match {
    range: token::Range,
    captures: Vec<usize>,
    offset: usize,
}

pub enum ImplMatcher {
    Todo(block::todo::TodoItemBlock),
    Code(block::code::CodeBlock),
    EmptyLine(block::empty::EmptyLineBlock),
}

impl ImplMatcher {
    fn matcher(&self, texts: &str, start_index: usize) -> Option<Match> {
        match self {
            ImplMatcher::Code(code) => code.matcher(texts, start_index),
            ImplMatcher::Todo(todo) => todo.matcher(texts, start_index),
            ImplMatcher::EmptyLine(line) => line.matcher(texts, start_index),
        }
    }

    fn tokens(&self, _match: &Match, nodes: &mut Vec<Node>) {
        match self {
            ImplMatcher::Code(code) => code.tokens(_match, nodes),
            ImplMatcher::Todo(todo) => todo.tokens(_match, nodes),
            ImplMatcher::EmptyLine(line) => line.tokens(_match, nodes),
        }
    }
}

pub fn parser(texts: &str) -> Vec<Node> {
    let blocks: Vec<ImplMatcher> = vec![
        ImplMatcher::Todo(block::todo::TodoItemBlock {}),
        ImplMatcher::Code(block::code::CodeBlock {}),
        ImplMatcher::EmptyLine(block::empty::EmptyLineBlock {}),
    ];
    let mut nodes: Vec<Node> = vec![];
    let mut _texts = texts;

    let ids: Vec<usize> = texts.char_indices().map(|(id, _)| id).collect();

    println!("{:?}", ids);

    let mut id: usize = 0;
    while id < ids[ids.len() - 1] {
        // println!("======{}=======", _texts);
        let mut flag = false;
        for block in blocks.iter() {
            if let Some(_match) = block.matcher(_texts, id) {
                block.tokens(&_match, &mut nodes);
                // 更新texts
                let range = _match.range;
                id = range.end + id;
                println!("{:?}", range);
                flag = true;
                break;
            }
        }
        // 一个都没有匹配
        if !flag {
            break;
        }
    }
    nodes
}
