fn main() {
    // create: String::new()
    let mut name = String::new();
    // create: .to_string()
    let id = "2100770220".to_string();
    // create: String::from()
    let address = String::from("New York");

    // method: push_str()
    name.push_str("Henry");
    // method: push()
    name.push(' ');
    name.push_str("Stephen");

    println!("{name} is {} long", name.len());
    println!("{id} is {} long", id.len());
    println!("{address} is {} long", address.len());

    // UTF-8
    let sentence = String::from("你好👋，世界🌍");
    println!("{sentence} is {} long", sentence.len());
    // -> char
    let words = sentence.chars();
    for word in words {
        print!("{}", word);
    }
    // method: Chars
    let word = match sentence.chars().nth(2) {
        Some(word) => word,
        None => {
            println!("ERROR");
            ' '
        }
    };
    println!();
    println!("{}", word);
    // -> bytes
    let bytes = sentence.bytes();
    for byte in bytes {
        print!("{} ", byte);
    }
    println!();
}
