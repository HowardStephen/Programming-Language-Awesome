use std::io;
use rand::prelude::*;

// 标准
//const NUM: u8 = 20;

// 接收输入函数
fn get_input() -> u8 {
    let mut input = String::new();

    // 提示输入
    println!("请输入一个数字:");
    io::stdin()
        .read_line(&mut input)
        .expect("ERROR: INPUT");

    // 输入处理
    let num: u8 = match input.trim().parse() {
        Ok(num) => num,
        Err(_) => {
            println!("ERROR: TRANSFER");
            0
        },
    };
    num
}

// 逻辑判断
fn range_judge(standard: u8, num: u8) -> bool {
    if num < standard {
        println!("SMALL");
        false
    } else if num > standard {
        println!("BIG");
        false
    } else {
        println!("RIGHT");
        true
    }
}

fn main() {
    // generate a guess_num
    let mut rng = rand::rng();
    let guess_num: u8 = rng.random_range(1..=100);

    let mut num: u8;

    let mut flag = false;

    while !flag {
        num = get_input();
        flag = range_judge(guess_num, num);
    }
}