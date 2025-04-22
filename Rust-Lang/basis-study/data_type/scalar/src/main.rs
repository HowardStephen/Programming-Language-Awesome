fn main() {
    // integer - default i32
    // u8, i8, u16, i16, u32, i32, u64, i64, u128, i128, usize, usize
    let age: u8 = 22;
    let area: u16 = 20000;
    let neg_num: i8 = -10;

    // float - default f64
    // f32, f64
    let pi: f32 = 3.1415;

    // boolean
    // bool -> false & true
    let flag: bool = false;

    // character
    // char: 4 bytes
    let rust: char = '🦀';

    println!("{age}, {area}, {neg_num}, {pi}, {flag}, {rust}");
}
