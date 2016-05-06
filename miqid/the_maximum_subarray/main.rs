use std::io::{self, BufRead};

fn main() {
    let stdin = io::stdin();
    let mut line = String::new();
    // Parse number of test cases
    stdin.read_line(&mut line).unwrap();
    let cases = line.trim().parse::<usize>().unwrap();
    for _ in 0..cases {
        // Ignore input size
        stdin.read_line(&mut line).unwrap();
        // Parse input values
        line.clear();
        stdin.read_line(&mut line).unwrap();
        let values = line
                     .trim()
                     .split_whitespace()
                     .map(|c| c.parse::<i32>().unwrap())
                     .collect::<Vec<_>>();
        println!("{} {}", max_contiguous(&values), max_noncontiguous(&values));
    }
}

pub fn max_contiguous(values: &[i32]) -> i32 {
    match values.len() {
        0 => return 0,
        1 => return values[0],
        _ => { /* no-op */ },
    }
    let mut max = values[0];
    let mut running_max = max;
    for i in 1..values.len() {
        let v = values[i];
        if running_max + v < v {
            running_max = v;
        } else {
            running_max += v;
        }
        if running_max > max {
            max = running_max;
        }
    }
    max
}

pub fn max_noncontiguous(values: &[i32]) -> i32 {
    match values.len() {
        0 => return 0,
        1 => return values[0],
        _ => { /* no-op */ },
    }
    let mut running_max = values[0];
    for i in 1..values.len() {
        let v = values[i];
        if running_max < v && running_max + v < v {
            running_max = v;
        } else if running_max + v > running_max {
            running_max += v;
        }
    }
    running_max
}

#[cfg(test)]
mod tests {
    use super::*;
    #[test]
    fn test_samples() {
        assert_eq!(10, max_contiguous(&vec![1, 2, 3, 4]));
        assert_eq!(10, max_noncontiguous(&vec![1, 2, 3, 4]));
        assert_eq!(10, max_contiguous(&vec![2, -1, 2, 3, 4, -5]));
        assert_eq!(11, max_noncontiguous(&vec![2, -1, 2, 3, 4, -5]));
    }
}
