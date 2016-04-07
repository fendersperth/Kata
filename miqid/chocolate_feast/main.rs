use std::io::{self, BufRead};

fn main() {
    let stdin = io::stdin();
    for line in stdin.lock().lines().skip(1) {
        let params : Vec<u32> = line
                                .unwrap()
                                .split_whitespace()
                                .map(|c| c.parse::<u32>().unwrap())
                                .collect();
        println!("{}", choc_count(params[0], params[1], params[2]));
    }
}

pub fn choc_count(funds: u32, choc_cost: u32, bonus_count: u32) -> u32 {
    let count = funds / choc_cost;
    count + consume(count, bonus_count)
}

fn consume(count: u32, bonus_count: u32) -> u32 {
    let quotient = count / bonus_count;
    let remainder = count % bonus_count;
    if quotient == 0 {
        return 0;
    }
    quotient + consume(quotient + remainder, bonus_count)
}

#[cfg(test)]
mod tests {
    use super::*;
    #[test]
    fn test_samples() {
        assert_eq!(choc_count(10, 2, 5), 6);
        assert_eq!(choc_count(12, 4, 4), 3);
        assert_eq!(choc_count(6, 2, 2), 5);
    }
}
