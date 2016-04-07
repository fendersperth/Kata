use std::io::{self, BufRead};

fn main() {
    let stdin = io::stdin();
    let mut line = String::new();
    stdin.read_line(&mut line).unwrap();
    let cases : u32 = line.trim().parse().unwrap();
    for _ in 0..cases {
        // Parse case parameters
        line.clear();
        stdin.read_line(&mut line).unwrap();
        let params : Vec<u32> = line
                                .split_whitespace()
                                .map(|c| c.parse::<u32>().unwrap())
                                .collect();
        line.clear();
        // Parse arrival times
        line.clear();
        stdin.read_line(&mut line).unwrap();
        let times : Vec<i32> = line
                               .split_whitespace()
                               .map(|c| c.parse::<i32>().unwrap())
                               .collect();
        if is_cancelled(params[1], &times) {
            println!("YES");
        } else {
            println!("NO");
        }
    }
}

pub fn is_cancelled(threshold: u32, times: &Vec<i32>) -> bool {
    threshold > times
                .to_owned()
                .iter()
                .fold(0, |total, &time| total + (time <= 0) as u32)
}

#[cfg(test)]
mod tests {
    use super::*;
    #[test]
    fn test_samples() {
        assert_eq!(is_cancelled(3, &vec![-1, -3, 4, 2]), true);
        assert_eq!(is_cancelled(2, &vec![0, -1, 2, 1]), false);
    }
}
