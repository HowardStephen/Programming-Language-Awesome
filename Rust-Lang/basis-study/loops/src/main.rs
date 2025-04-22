fn main() {
    let mut num = 0;
    num = loop {
        num += 1;
        if num > 10 {
            // break return a value
            // but with `;` !
            break num - 1;
        }
    };

    println!("{}", num);

    // shadow
    let mut num = 0;
    // label - loop
    'first: loop {
        num += 1;
        print!("{} ", num);

        loop {
            num += 2;
            if num > 20 {
                break 'first;
            }
            print!("{} ", num);
        }
    }
    println!("{}", num);

    // while
    let mut num = 0;
    while num < 10 {
        num += 1;
    }
    println!("{}", num);

    // for
    let nums = [1, 2, 3, 4, 5];
    for num in nums {
        print!("{} ", num);
    }
    println!();
    // range in for
    for num in (1..=4).rev() {
        print!("{num} ");
    }
    println!();
}
