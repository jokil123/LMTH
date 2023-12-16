use std::fmt::{self, Display, Formatter};

#[derive(Debug, PartialEq)]
pub enum Token {
    Text(String),
    Slash,
    ElementStart,
    ElementEnd,
    DoubleQuote,
    SingleQuote,
    EqualSign,
    Whitespace(i32),
}

impl Display for Token {
    fn fmt(&self, f: &mut Formatter<'_>) -> fmt::Result {
        match self {
            Token::Whitespace(n) => {
                for _ in 0..*n {
                    write!(f, " ")?;
                }
                Ok(())
            }
            Token::Text(text) => write!(f, "{}", text),
            Token::Slash => write!(f, "/"),
            Token::ElementStart => write!(f, "<"),
            Token::ElementEnd => write!(f, ">"),
            Token::DoubleQuote => write!(f, "\""),
            Token::SingleQuote => write!(f, "'"),
            Token::EqualSign => write!(f, "="),
        }
    }
}

pub fn tokenize(html: &str) -> anyhow::Result<Vec<Token>> {
    let mut tokens = vec![];

    let mut chars = html.chars().peekable();

    while let Some(c) = chars.next() {
        match c {
            '\n' | '\t' => continue,
            ' ' => {
                let mut n = 1;

                while let Some(&c) = chars.peek() {
                    match c {
                        ' ' => {
                            n += 1;
                            chars.next();
                        }
                        _ => break,
                    }
                }

                tokens.push(Token::Whitespace(n));
            }
            '<' => tokens.push(Token::ElementStart),
            '>' => tokens.push(Token::ElementEnd),
            '"' => tokens.push(Token::DoubleQuote),
            '\'' => tokens.push(Token::SingleQuote),
            '=' => tokens.push(Token::EqualSign),
            '/' => tokens.push(Token::Slash),
            _ => {
                let mut text = String::new();
                text.push(c);

                while let Some(&c) = chars.peek() {
                    match c {
                        '\n' | '\t' | ' ' | '<' | '>' | '"' | '\'' | '=' | '/' => break,
                        _ => {
                            text.push(c);
                            chars.next();
                        }
                    }
                }

                tokens.push(Token::Text(text));
            }
        }
    }

    Ok(tokens)
}
