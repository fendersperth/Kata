use std::collections::HashMap;
use std::io::{self, BufRead};

/// A key input constraint to note is that `numbers` contains a set of natural
/// numbers starting from one. The first value in `numbers` to swap into place
/// is `max`, with each subsequent value being less by one. Using `indices` to
/// access the index of values in `numbers` is roughly constant time. This
/// is more time efficient than repeatedly searching for the index of the next
/// value in `numbers` to swap.
fn permute(swaps: u32, numbers: &mut [u32], max: u32, indices: HashMap<u32, usize>) {
    let mut indices = indices.clone();
    let mut i : usize = 0;
    let mut j_val = max;
    let mut swaps = swaps;
    while swaps > 0 && i < numbers.len() - 1 {
        let j = indices.get(&j_val).cloned().unwrap();
        // Swap only if necessary
        if j != i {
            let i_val = numbers[i];
            numbers.swap(i, j);
            // Updating the index cache for the number swapped into its final
            // position is unnecessary since it should not be accessed again
            indices.insert(i_val, j);
            swaps -= 1;
        }
        j_val -= 1;
        i += 1;
    }
}

fn main() {
    let stdin = io::stdin();
    // Parse swaps allowed
    let mut line = String::new();
    stdin.read_line(&mut line).unwrap();
    let swaps : u32 = line
                      .split_whitespace()
                      .map(|c| c.parse::<_>().unwrap())
                      .nth(1)
                      .unwrap();
    // Parse permutation
    line.clear();
    stdin.read_line(&mut line).unwrap();
    let mut max = 0;
    let mut indices : HashMap<u32, usize> = HashMap::new();
    let mut numbers : Vec<u32> = line
                                 .split_whitespace()
                                 .enumerate()
                                 .map(|(i, c)| {
                                     let n =  c.parse::<_>().unwrap();
                                     indices.insert(n, i);
                                     if n > max {
                                         max = n;
                                     }
                                     n
                                 })
                                 .collect();
    permute(swaps, &mut numbers, max, indices);
    // Print permutation in the required format
    line = numbers
           .iter()
           .map(|c| c.to_string())
           .collect::<Vec<_>>()
           .join(" ");
    println!("{}", line);
}
