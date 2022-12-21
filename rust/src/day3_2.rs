pub fn main() {
    // TODO: use u8 in claculations but parse sum to u16
    let shit_lower = 'a' as u16;
    let shit_upper = 'A' as u16;

    let cool: u16 = std::fs::read_to_string("src/input/day3.prod")
        .expect("cannot read file")
        .lines()
        .array_chunks::<3>()
        .map(|chunk| {
            let a = chunk[0];
            let b = chunk[1];
            let c = chunk[2];

            for char in a.chars() {
                if b.find(char).is_some() && c.find(char).is_some() {
                    if char.is_uppercase() {
                        return 26 + char as u16 - shit_upper + 1;
                    }
                    return char as u16 - shit_lower + 1;
                }
            }
            panic!();
        })
        .sum();

    println!("{:?}", cool);
}
