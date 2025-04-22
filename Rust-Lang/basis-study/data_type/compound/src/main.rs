fn main() {
    // tuple
    //    1. contain different types
    //    2. a fixed length
    let tup: (i32, f32, bool) = (1, 1.0, true);
    // deconstruct
    let (x, y, z) = tup;
    println!("{x}, {y}, {z}");
    // get item in tuple - start from 0
    println!("{}, {}, {}", tup.0, tup.1, tup.2);
    // () is important -> unit tuple

    // array
    //    1. same type
    //    2. a fixed length -> vector(unfixed)
    let a: [i32; 5] = [1, 2, 3, 4, 5];
    let b = [1; 5]; // initial
    // get item in array - start from 0
    println!("{}, {}", a[0], b[4]);
}
