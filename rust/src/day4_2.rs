use std::fs;

pub fn main() {
    let numpairs: Vec<_> = fs::read_to_string("src/input/day4.prod")
        .expect("Could not read input")
        .lines()
        .map(|line| {
            return line
                .split(",") //
                .map(|x| x.parse::<String>().unwrap())
                .collect::<Vec<String>>();
        })
        .collect();

    // println!("{:?}", numpairs);

    let numranges: Vec<_> = numpairs
        .iter()
        .map(|pair| {
            return pair
                .iter()
                .map(|half| {
                    let nums: Vec<_> = half.split("-").collect();

                    let mut result = vec![];
                    for num in nums[0].parse::<u8>().unwrap()..nums[1].parse::<u8>().unwrap() + 1 {
                        result.push(num);
                    }
                    return result;
                })
                .collect::<Vec<_>>();
        })
        .collect();

    // println!("{:?}", numranges);

    let mut contains = 0;

    for mut pair in numranges {
        pair.sort_by_key(|a| a.len());
        pair.reverse();
        println!("{} : {}", &pair[0].len(), &pair[1].len());
        let to_find = &pair[1].len();
        let mut found: u8 = 0;
        let mut success = false;

        for num in &pair[1] {
            // println!("NUM: {}\nSHORT: {:?}\nLONG: {:?}", num, &pair[1], &pair[0]);

            if !pair[0].contains(&num) {
                // println!("NOT");
                continue;
            }
            found += 1;
            // println!("FOUND");

            if found as usize >= 1 {
                success = true;
                // println!("SUCCESS");
                break;
            }
        }

        if success {
            contains += 1;
            // println!("CONTAINS");
        }
        // println!("\n");
    }

    println!("{:?}", contains);
}
