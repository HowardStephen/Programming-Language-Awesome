#[derive(Debug)]
struct Rectangle {
    width: u32,
    height: u32,
}

// method for Rectangle
impl Rectangle {
    fn print_rec(&self) {
        println!("{:#?}", self);
    }

    fn rec_area(&self) -> u32 {
        self.width * self.height
    }
}

fn main() {
    let rec = Rectangle {
        width: 30,
        height: 30,
    };
//    println!("area is {}", rec_area(&rec)); // function
    rec.print_rec();
    println!("area is {}", rec.rec_area()); // method
}

//fn rec_area(rec: &Rectangle) -> u32 {
//    println!("{:#?}", rec);
//    rec.width * rec.height
//}
