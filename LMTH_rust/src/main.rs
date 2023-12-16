use LMTH_rust::html;

fn main() {
    let fs = std::fs::read_to_string("input.html").unwrap();

    let tokens = html::tokens::tokenize(&fs).unwrap();

    println!("{:?}", tokens);
}
