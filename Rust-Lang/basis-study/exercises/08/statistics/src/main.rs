use std::collections::HashMap;
use std::fs::File;
use std::io::{BufRead, BufReader, ErrorKind};

// 实现内容：
//     1. 计算平均数
//     2. 计算中位数
//     3. 众数

// 补充实现：
//     1. 从文件中读取数据 -> 存储到动态数组中
fn read_file(file_path: &str) -> String {
    let f = match File::open(file_path) {
        Ok(it) => it,
        Err(err) => match err.kind() {
            ErrorKind::NotFound => {
                return "file not found".to_string()
            },
            _ => {
                return "other error".to_string()
            },
        }
    };
    let mut reader = BufReader::new(f);
    let mut buffer = String::new();

    match reader.read_line(&mut buffer) {
        Ok(_) => {
            buffer
        }
        Err(_) => {
            panic!("Error!");
        }
    }
}

//     2. 实现一个排序算法
fn sort_numbers(num_vec: &mut Vec<i32>) {
    let len = num_vec.len();
    for i in 0..len {
        let mut min_index = i;
        for j in i..len {
            if num_vec[j] < num_vec[min_index] {
                min_index = j;
            }
        }
        num_vec.swap(i, min_index);
    }
}

//     3. 使用哈希映射

fn main() {
    let text = read_file("sources/numbers.txt");
    let numbers = text.split_whitespace();

    let mut num_vec: Vec<i32> = Vec::new();
    for item in numbers {
        num_vec.push(item.parse::<i32>().unwrap());
    }

    let sum: i32 = num_vec.iter().sum();
    let len = num_vec.len();

    println!("平均数: {}", sum as f64 / len as f64);

    sort_numbers(&mut num_vec);

    let mid = match len % 2 {
        0 => (num_vec[len / 2] + num_vec[len / 2 - 1]) / 2,
        1 => num_vec[len / 2],
        _ => -1
    };

    println!("中位数: {}", mid);

    let mut map = HashMap::new();
    for num in &num_vec {
        let count = map.entry(*num).or_insert(0);
        *count += 1;
    }
    let mut max_k: i32 = 0;
    let mut max_v: i32 = i32::MIN;
    for (k, v) in &map {
        if *v > max_v {
            max_v = *v;
            max_k = *k;
        }
    }

    println!("众数: {}", max_k);
}
