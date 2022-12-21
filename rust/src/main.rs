#![feature(iter_array_chunks)]
use std::fs;

mod day5_1;

fn main() {
    let input: Vec<String> = fs::read_to_string("src/input/day5.test")
        .expect("Could not read file")
        .lines()
        .map(|line| line.to_string())
        .collect();

    day5_1::main(input);
}
