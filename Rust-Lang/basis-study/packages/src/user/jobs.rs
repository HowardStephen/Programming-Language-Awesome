pub enum Job {
    Teacher,
    Student,
    Worker,
    Engineer
}

impl Job {
    pub fn print_job(&self) {
        match self {
            Job::Teacher => println!("teacher"),
            Job::Student => println!("student"),
            Job::Worker => println!("worker"),
            Job::Engineer => println!("engineer"),
        }
    }
}