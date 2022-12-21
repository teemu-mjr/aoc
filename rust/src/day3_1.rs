pub fn main() {
    let chars = vec![
        'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r',
        's', 't', 'u', 'v', 'w', 'x', 'y', 'z', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J',
        'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
    ];
    let input: usize = include_str!("./input/day3.prod")
        .lines()
        .map(|x| {
            let (a, b) = x.split_at(x.len() / 2);
            for i in 0..a.len() {
                for j in 0..b.len() {
                    if a.chars().nth(i).unwrap() == b.chars().nth(j).unwrap() {
                        for (k, char) in chars.iter().enumerate() {
                            if a.chars().nth(i).unwrap() == *char {
                                return k + 1;
                            }
                        }
                    }
                }
            }
            panic!("SHIIIS");
        })
        .sum::<usize>();

    println!("{:?}", input);
}
