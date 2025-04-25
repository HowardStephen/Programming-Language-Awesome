fn main() {
    // create: Vec::new() -> Vec<i32>
    let mut nums: Vec<i32> = Vec::new();
    // method: push()
    nums.push(1);
    nums.push(2);
    nums.push(3);
    println!("{:?}", nums);

    // create: vec![]
    let mut other_nums = vec![1, 2, 3];
    println!("{:?}", other_nums);
    // method: pop() -> Option<T>
    let num = match other_nums.pop() {
        Some(num) => num,
        None => {
            println!("ERROR: EMPTY");
            0
        }
    };
    println!("{} has been poped", num);
}