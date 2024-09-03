use std::fs;
use std::str::FromStr;

#[derive(Debug, PartialEq)]
struct Race {
    length: u64,
    record: u64,
}

impl FromStr for Race {
    type Err = std::convert::Infallible; // lol

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let mut lines = s.lines();
        let line_length = lines.next().expect("input must have two lines");
        let line_record = lines.next().expect("input must have two lines");

        let length: u64 = line_length
            .trim()
            .split_whitespace()
            .skip(1)
            .fold("".to_string(), |num_str, num_part| num_str + num_part)
            .parse().unwrap_or(0);
        let record: u64 = line_record
            .trim()
            .split_whitespace()
            .skip(1)
            .fold("".to_string(), |num_str, num_part| num_str + num_part)
            .parse().unwrap_or(0);

        Ok(Race { length, record })
    }
}

fn main() {
    //let input = fs::read_to_string("input.test").expect("input is readable");
    let input = fs::read_to_string("input").expect("input is readable");
    let race: Race = input.parse().expect("input is valid race");
    println!("{:?}", race);

    let mut win_count = 0u64;
    for speed in 0..=race.length {
        // {7, 9}
        // x + y = 7   [I]
        // x * y > 9   [II]
        let distance = speed * (race.length - speed);
        if distance > race.record {
            win_count += 1;
        }
    }

    println!("{win_count}");
}
