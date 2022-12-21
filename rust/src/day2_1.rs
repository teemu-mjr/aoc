pub fn main() {
    let cool: usize = include_str!("./input/day2.prod")
        .lines()
        .flat_map(str::parse::<String>)
        .map(|x| {
            return x
                .split(" ") //
                .map(|x| match x {
                    "A" => 1,
                    "B" => 2,
                    "C" => 3,
                    "X" => 1,
                    "Y" => 2,
                    "Z" => 3,
                    _ => 9,
                })
                .collect::<Vec<usize>>();
        })
        .map(|x| {
            if x[0] == x[1] {
                return x[1] + 3;
            }

            if x[0] == 1 {
                if x[1] == 2 {
                    return x[1] + 6;
                }
                else {
                    return x[1] + 0;
                }
            }

            if x[0] == 2 {
                if x[1] == 3 {
                    return x[1] + 6;
                }
                else {
                    return x[1] + 0;
                }
            }

            if x[0] == 3 {
                if x[1] == 1 {
                    return x[1] + 6;
                }
                else {
                    return x[1] + 0;
                }
            }

            else {
                return 0;
            }
        })
        .sum::<usize>();

    println!("{:?}", cool);

    // A Y
    // B X
    // C Z

    // 0
    // A => rock
    // B => paper
    // C => scissors

    // 1
    // X => rock
    // Y => paper
    // Z => scissors

    // score
    // rock     => 1
    // paper    => 2
    // scissors => 3
    // +
    // lost     => 0
    // draw     => 3
    // win      => 6
}
