use std::fs::File;
use std::io::{BufRead, BufReader};

pub fn main() {
    let filename = "./input/day1.prod";
    let file = File::open(filename).unwrap();
    let reader = BufReader::new(file);

    let mut top_elfs = [[0; 2], [0; 2], [0; 2]];

    let mut calories: Vec<i32> = Vec::new();
    calories.push(0);

    let mut elf_index = 0;

    for line in reader.lines() {
        let line = line.unwrap();
        if line == "" {
            for i in 0..top_elfs.len() {
                if calories[elf_index] > top_elfs[i][1] {
                    for j in (i + 1..top_elfs.len()).rev() {
                        print!("{}: {} -->", i, top_elfs[j][1]);
                        top_elfs[j] = top_elfs[j - 1];
                        println!(" {}", top_elfs[j][1]);
                    }

                    top_elfs[i][0] = elf_index as i32;
                    top_elfs[i][1] = calories[elf_index];
                    break;
                }
            }

            elf_index += 1;
            calories.push(0);
            continue;
        }

        let new_calories: i32 = calories[elf_index] as i32 + line.parse::<i32>().unwrap();
        calories[elf_index] = new_calories;
    }

    // uuf
    for i in 0..top_elfs.len() {
        if calories[elf_index] > top_elfs[i][1] {
            for j in (i + 1..top_elfs.len()).rev() {
                print!("{}: {} -->", i, top_elfs[j][1]);
                top_elfs[j] = top_elfs[j - 1];
                println!(" {}", top_elfs[j][1]);
            }

            top_elfs[i][0] = elf_index as i32;
            top_elfs[i][1] = calories[elf_index];
            break;
        }
    }

    let mut total = 0;
    for elf in top_elfs {
        total += elf[1];
        println!("TOP {}: {}", elf[0], elf[1])
    }
    println!("{}", total);
}
