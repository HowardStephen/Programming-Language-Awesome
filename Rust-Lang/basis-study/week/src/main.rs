#[derive(Debug)]
enum WeekDay {
    Sunday(u8),
    Monday(u8),
    Thuesday(u8),
    Wednesday(u8),
    Thursday(u8),
    Friday(u8),
    Saturday(u8),
}

fn main() {
    let weekdays: [WeekDay; 5] = [
        WeekDay::Monday(1),
        WeekDay::Thuesday(2),
        WeekDay::Wednesday(3),
        WeekDay::Thursday(4),
        WeekDay::Friday(5),
    ];

    for day in weekdays {
//        println!("{:?}", day);
        if let WeekDay::Friday(num) = day {
            println!("{num}");
        }
    }

    let num: Option<u8> = None;
    if let Some(x) = num {
        println!("{}", x);
    } else {
        println!("!!!!");
    }
}
