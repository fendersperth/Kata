use std::io::{self, BufRead};

fn main() {
    let stdin = io::stdin();
    let lengths : Vec<u32> = stdin.lock().lines().nth(1).unwrap().unwrap()
                             .split_whitespace()
                             .map(|c| c.parse().unwrap())
                             .collect();
    cut(lengths);
}

fn cut(lengths: Vec<u32>) {
    if lengths.len() == 0 {
        return;
    }
    println!("{}", lengths.len());
    let cut_length = lengths.iter().min().unwrap();
    let new_lengths : Vec<u32> = lengths.iter()
                                 .map(|length| length - cut_length)
                                 .filter(|&length| length > 0)
                                 .collect();
    cut(new_lengths);
}
