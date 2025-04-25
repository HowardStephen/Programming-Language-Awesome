use core::num;

fn main() {
    println!("Hello, world!");
    another_hello();
    function_with_para(22, '🦀');
    println!("10 + 10 = {}", add_two_num(10, 10));
}

// function with no parameter(argument)
fn another_hello() {
    println!("Hello, Rust!");
}

// function with paramater(argument)
fn function_with_para(num: i32, ch: char) {
    println!("{num}: {ch}");
}

// return value
fn add_two_num(x: i32, y: i32) -> i32 {
    x + y
}