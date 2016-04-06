use std::io::{self, BufRead};

fn main() {
    let cavity_map = read_cavity_map();
    process(&cavity_map);
}

fn read_cavity_map() -> Vec<Vec<u32>> {
    let mut cavity_map : Vec<Vec<u32>> = Vec::new();
    let stdin = io::stdin();
    for line in stdin.lock().lines().skip(1) {
        let row : Vec<u32> = line.unwrap()
                             .chars()
                             .map(|c| c.to_digit(10).unwrap())
                             .collect();
        cavity_map.push(row);
    }
    cavity_map
}

fn process(cavity_map: &Vec<Vec<u32>>) {
    let length = cavity_map.len();
    for i in 0..length {
        for j in 0..length {
            if i == 0 || i == length - 1 || j == 0 || j == length - 1 {
                print!("{}", cavity_map[i][j]);
                continue;
            }
            let point = cavity_map[i][j];
            if point > cavity_map[i - 1][j] &&
               point > cavity_map[i + 1][j] &&
               point > cavity_map[i][j - 1] &&
               point > cavity_map[i][j + 1] {
                print!("X");
                continue;
            }
            print!("{}", point);
        }
        println!("");
    }
}
