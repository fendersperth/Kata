use std::io::{self, BufRead};

fn main() {
    let stdin = io::stdin();
    let mut line = String::new();
    // Parse number of test cases
    stdin.read_line(&mut line).unwrap();
    let cases = line.trim().parse::<usize>().unwrap();
    for _ in 0..cases {
        // Parse test case parameters
        line.clear();
        stdin.read_line(&mut line).unwrap();
        let case_params = line
                          .split_whitespace()
                          .map(|c| c.parse::<usize>().unwrap())
                          .collect::<Vec<_>>();
        // Parse edge definitions
        let mut edges : Vec<(usize, usize)> = Vec::new();
        for _ in 0..case_params[1] {
            line.clear();
            stdin.read_line(&mut line).unwrap();
            let mut edge = line
                           .split_whitespace()
                           .map(|c| c.parse::<usize>().unwrap());
            edges.push((edge.next().unwrap(), edge.next().unwrap()));
        }
        // Parse starting node
        line.clear();
        stdin.read_line(&mut line).unwrap();
        let start_node = line.trim().parse::<usize>().unwrap();
        // Calculate output
        let distances = shortest_reach(edges, case_params[0], start_node);
        // Format answer output
        let mut output = String::new();
        for &d in distances.iter() {
            if d == 0 {
                continue;
            }
            if !output.is_empty() {
                output.push_str(" ");
            }
            output.push_str(&d.to_string());
        }
        println!("{}", output);
    }
}

/// Based on the Bellman-Ford algorithm. Does not check for negative weight
/// cycles as input problem set is assumed to not have any. Nodes are also
/// assumed to be indexed from one for identification.
pub fn shortest_reach(edges: Vec<(usize, usize)>,
                      total_nodes: usize,
                      start_node: usize) -> Vec<i32> {
    const EDGE_WEIGHT : i32 = 6;
    const UNREACHABLE : i32 = -1;
    let mut distances =
        (0..).take(total_nodes).map(|_| UNREACHABLE).collect::<Vec<i32>>();
    distances[start_node - 1] = 0;
    for _ in 0..distances.len() {
        for &(u, v) in edges.iter() {
            let u = u - 1;
            let v = v - 1;
            if distances[u] != UNREACHABLE {
                let w = distances[u] + EDGE_WEIGHT;
                if distances[v] == UNREACHABLE || w < distances[v] {
                    distances[v] = w;
                }
            }
            if distances[v] != UNREACHABLE {
                let w = distances[v] + EDGE_WEIGHT;
                if distances[u] == UNREACHABLE || w < distances[u] {
                    distances[u] = w;
                }
            }
        }
    }
    distances
}

#[cfg(test)]
mod tests {
    use super::*;
    #[test]
    fn test_samples() {
        assert_eq!(
            vec![0, 6, 6, -1],
            shortest_reach(vec![(1, 2), (1, 3)], 4, 1)
        );
        assert_eq!(
            vec![-1, 0, 6],
            shortest_reach(vec![(2, 3)], 3, 2)
        );
    }
}
