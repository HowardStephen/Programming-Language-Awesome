use std::collections::HashMap;

fn main() {
    let mut schedules = HashMap::new();

    // method: insert(k, v)
    schedules.insert("monday", "apple");
    schedules.insert("tuesday", "juice");

    let fruit = schedules.get("monday").copied().unwrap_or("none");
    println!("monday will eat {}", fruit);

    // traverse
    for (k, v) in &schedules {
        println!("{k}: {v}");
    }

    // update
    schedules.insert("monday", "orange");
    let fruit = schedules.get("monday").copied().unwrap_or("none");
    println!("monday will eat {}", fruit);

    // method: entry()
    schedules.entry("wednesday").or_insert("apple");
    schedules.entry("tuesday").or_insert("apple"); // unchange

    for (k, v) in &schedules {
        println!("{k}: {v}");
    }

    // count words
    let mut counter = HashMap::new();
    let words = "hello world hello cargo hello rust rust cargo rustup";
    for word in words.split_whitespace() {
        let count = counter.entry(word).or_insert(0);
        *count += 1;
    }
    for (k, v) in &counter {
        println!("{k}: {v}");
    }
}
