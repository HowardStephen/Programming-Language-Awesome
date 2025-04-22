use std::io; // 使用标准输入接受输入

fn main() {
    // 接收选项输入
    let mut input = String::new();
    println!("Make choice(0 or 1): ");
    io::stdin()
        .read_line(&mut input)
        .expect("ERROR: Input!");
    let choice: u8 = match input.trim().parse() {
        Ok(choice) => choice,
        Err(_) => {
            println!("ERROR(1): Transfer!");
            return;
        }
    };

    // 接收内容输入
    let mut input = String::new();
    println!("请输入要转化的值: ");
    io::stdin()
        .read_line(&mut input)
        .expect("ERROR: Input!");
    let num: f64 = match input.trim().parse() {
        Ok(num) => num,
        Err(_) => {
            println!("ERROR(2): Transfer!");
            return;
        }
    };

    if choice == 0 {
        // 摄氏度 -> 华氏度
        println!("C: {} -> F: {}", input.trim(), c_to_f(num));
    } else {
        // 华氏度 -> 摄氏度
        println!("F: {} -> C: {}", input.trim(), f_to_c(num));
    }
}

// 摄氏度 -> 华氏度
fn c_to_f(c: f64) -> f64 {
    return c * (9.0 / 5.0) + 32.0;
}

// 华氏度 -> 摄氏度
fn f_to_c(f: f64) -> f64 {
    return (f - 32.0) * 5.0 / 9.0;
}