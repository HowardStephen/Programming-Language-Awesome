pub mod user;

use crate::user::{User, Job};

fn main() {
    let rand_user = User::new("Henry", 22, Job::Teacher, 2100770220, true);
    println!("{}", rand_user.name);
    println!("{}", rand_user.age);
    println!("{}", rand_user.get_id());
    println!("{}", rand_user.get_vip());
    rand_user.job.print_job();
}
