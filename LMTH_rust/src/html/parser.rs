use crate::html::ast::{Attribute, Element, Node, SelfClosingElement};

use self::ParseState::*;
use crate::html::tokens::Token as t;

use super::ast::NodeType as nt;

// Recursive function that parses HTML into an AST
fn parse_html(html: &str) -> anyhow::Result<Node> {
    let mut root = Node {
        parent: None,
        node_type: nt::Element(Element {
            tag_name: "html".to_string(),
            attributes: vec![],
            children: vec![],
        }),
    };

    let mut current_node: &mut Node = &mut root;

    let tokens = crate::html::tokens::tokenize(html)?;

    let mut state = ParseState::InText;
    let mut buffer = String::new();
    let mut nesting_level = 0;

    for token in tokens {
        match (&state, &current_node.node_type, token) {
            (InText, nt::Element(e), t::ElementStart) => {
                state = InTagName;
            }
            (InTagName, nt::Element(e), t::Text(text)) => {
                buffer.push_str(&text);
            }
            (InTagName, nt::Element(e), to) => {
                if buffer.is_empty() {
                    return Err(anyhow::anyhow!("Empty tag name"));
                }

                let tag_name = buffer.clone();
                buffer.clear();
            }
            (_, _, _) => {}
        }
    }

    Ok(root)
}

enum ParseState {
    InTagName,
    InTagName,
    InText,
    InComment,
    InDoctype,
    InAttribute,
    InAttributeValue,
}
