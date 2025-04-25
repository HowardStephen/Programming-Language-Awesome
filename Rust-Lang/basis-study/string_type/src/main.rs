fn main() {
    let mut words = String::from("Hello World Hello Cargo Hello Rust");

    words.push(' ');
    words.push_str("Hello Henry");

    println!("{}", words);

    // Hello World Hello Cargo Hello Rust Hello Henry
    // 0123456789012345678901234567890123456789012345

    let word = words.as_bytes();
    for (i, &w) in word.iter().enumerate() {
        if w == b' ' {
            println!("{}", &words[..i]);
        }
    }
}
