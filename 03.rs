use std::collections::HashSet;
use std::fs::File;
use std::io::{BufRead, BufReader};

fn main() {
    let file = File::open("./inputs/03.txt").unwrap();
    let reader = BufReader::new(file);
    let lines: Vec<String> = reader.lines().map(|l| l.unwrap()).collect();
    let priors = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ";
    let mut total = 0;

    for l in lines.clone() {
        let chars: Vec<char> = l.chars().collect();
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

    total = 0;

    for i in 0..(lines.len() / 3) {
        let chars1: Vec<char> = lines[3 * i].chars().collect();
        let chars2: Vec<char> = lines[3 * i + 1].chars().collect();
        let chars3: Vec<char> = lines[3 * i + 2].chars().collect();
        let mut seen1: HashSet<char> = HashSet::new();
        let mut seen2: HashSet<char> = HashSet::new();
        for c in chars1 {
            seen1.insert(c);
        }
        for c in chars2 {
            seen2.insert(c);
        }
        for c in chars3 {
            if seen1.contains(&c) && seen2.contains(&c) {
                total += priors.find(c).unwrap() + 1;
                break;
            }
        }
    }
    println!("{}", total);
}
