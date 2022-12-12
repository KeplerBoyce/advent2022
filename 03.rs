use std::collections::HashSet;
use std::fs::File;
use std::io::{BufRead, BufReader};

fn main() {
    let file = File::open("./inputs/03.txt").unwrap();
    let reader = BufReader::new(file);
    let lines = reader.lines();
    let priors = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ";
    let mut total = 0;

    for l in lines {
        let chars: Vec<char> = l.unwrap().chars().collect();
        let left: Vec<char> = chars.clone().drain(..(chars.len() / 2)).collect();
        let right: Vec<char> = chars.clone().drain((chars.len() / 2)..).collect();
        let mut seen: HashSet<char> = HashSet::new();

        for c in left {
            seen.insert(c);
        }
        for c in right {
            if seen.contains(&c) {
                total += priors.find(c).unwrap() + 1;
                break;
            }
        }
    }
    println!("{}", total);
}
