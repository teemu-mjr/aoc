use std::fs::File;
use std::io::{BufRead, BufReader};

pub fn main() {
    let filename = "";
    let file = File::open(filename).unwrap();
    let reader = BufReader::new(file);

    let mut beast_elf = vec![0, 0];
    let mut elf_index = 0;
    let mut calories: Vec<i32> = Vec::new();
    calories.push(0);

    for line in reader.lines() {
        let line = line.unwrap();
        if line == "" {
            if calories[elf_index] > beast_elf[1] {
                beast_elf[0] = elf_index as i32;
                beast_elf[1] = calories[elf_index];
            }

            elf_index += 1;
            calories.push(0);
            continue;
        }

        let new_calories: i32 = calories[elf_index] as i32 + line.parse::<i32>().unwrap();
        calories[elf_index] = new_calories;
    }

    println!("Elf number: {}, Calories: {}", beast_elf[0] + 1, beast_elf[1]);
}

