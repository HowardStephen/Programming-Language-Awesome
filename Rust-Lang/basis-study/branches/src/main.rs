fn main() {
    let num = 10;

    if num < 10 {
        println!("num < 10");
    } else if num > 10 {
        println!("num > 10");
    } else {
        println!("num = 10");
    }

    // if in let: return the same type value -> bool
    let flag = if num == 10 { true } else { false };

    println!("{}", flag);
}
