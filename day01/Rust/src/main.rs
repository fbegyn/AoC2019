use std::io;
const INPUT: &str = include_str!("../input.txt");

fn main() -> io::Result<()> {
    println!("{}",part1(INPUT));
    println!("{}",part2(INPUT));
    Ok(())
}

fn part1(file : &str) -> u32 {
    file
        .lines()
        .map(|input| {
            input.parse::<u32>().ok().and_then(|mass| (mass/3).checked_sub(2)).unwrap_or(0)
        })
        .sum::<u32>()
}

fn part2(file : &str) -> u32 {
    file
        .lines()
        .map(|input| {
            std::iter::successors(
                input.parse::<u32>().ok().and_then(|mass| (mass/3).checked_sub(2)),
                |&mass| (mass/3).checked_sub(2),
            ).sum::<u32>()
        })
        .sum::<u32>()
}

#[cfg(test)]
mod tests {
    use super::*;
    const TEST: &str = include_str!("../test.txt");

    #[test]
    fn test_part1(){
        assert_eq!(part1(TEST), 34241);
    }

    #[test]
    fn test_part2(){
        assert_eq!(part2(TEST), 51316);
    }
}
