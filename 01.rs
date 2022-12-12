use std::fs::File;
use std::io::{BufRead, BufReader};

fn main() {
    let file = File::open("./inputs/day1.txt").unwrap();
    let reader = BufReader::new(file);
    let lines = reader.lines();

    let mut max = 0;
    let mut max2 = 0;
    let mut max3 = 0;
    let mut current = 0;

    for l in lines {
        let line = l.unwrap();
        if line == "" {
            if current > max {
                max3 = max2;
                max2 = max;
                max = current;
            } else if current > max2 {
                max3 = max2;
                max2 = current;
            } else if current > max3 {
                max3 = current;
            }
            current = 0;
        } else {
            let num: i32 = line.parse().unwrap();
            current += num;
        }
    }
    println!("{}", max);
    println!("{}", max + max2 + max3);
}
