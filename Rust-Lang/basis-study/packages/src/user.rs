pub mod jobs;

pub use jobs::Job;

pub struct User {
    pub name: String,
    pub age: u8,
    pub job: Job,
    id: u32,
    vip: bool,
}

impl User {
    pub fn new(name: &str, age: u8, job: Job, id: u32, vip: bool) -> Self {
        User {
            name: name.to_string(),
            age,
            job,
            id,
            vip,
        }
    }

    pub fn get_id(&self) -> u32 {
        self.id
    }

    pub fn get_vip(&self) -> bool {
        self.vip
    }
}
