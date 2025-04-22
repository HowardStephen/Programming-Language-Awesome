use std::io;

fn get_io_num() -> i32 {
    let mut input = String::new();

    println!("please input a number: ");
    io::stdin()
        .read_line(&mut input)
        .expect("ERROR: INPUT!");

    let num: i32 = match input.trim().parse() {
        Ok(num) => num,
        Err(_) => {
            println!("ERROR: TRANSFER!");
            return 0;
        }
    };

    num
}

fn fibonacci(num: i32) -> i32 {
    if num == 1 {
        return 0;
    } else if num == 2 {
        return 1;
    }
    return fibonacci(num - 1) + fibonacci(num - 2);
}

fn main() {
    let num = get_io_num();
    let result = fibonacci(num);
    println!("n -> {}", result);
}
