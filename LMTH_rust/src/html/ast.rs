pub struct Node {
    pub parent: Option<Box<Node>>,
    pub node_type: NodeType,
}

pub enum NodeType {
    Text(String),
    Element(Element),
    SelfClosingElement(SelfClosingElement),
    Doctype(String),
    Comment(String),
}

pub struct Element {
    pub tag_name: String,
    pub attributes: Vec<Attribute>,
    pub children: Vec<Node>,
}

pub enum Attribute {
    KeyValue(String, String),
    Value(String),
}

pub struct SelfClosingElement {
    pub tag_name: String,
    pub attributes: Vec<Attribute>,
}
