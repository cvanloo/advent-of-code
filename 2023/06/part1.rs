use std::fs;

#[derive(Debug, PartialEq)]
struct Race {
    length: u32,
    record: u32,
}

fn main() {
    //let input = fs::read_to_string("input.test").unwrap();
    let input = fs::read_to_string("input").unwrap();
    let mut lines = input.lines();
    let race_lengths = lines.next().unwrap();
    let race_records = lines.next().unwrap();

    let lengths = race_lengths
        .trim()
        .split_whitespace()
        .skip(1)
        .map(|num| num.parse::<u32>().unwrap());
    let records = race_records
        .trim()
        .split_whitespace()
        .skip(1)
        .map(|num| num.parse::<u32>().unwrap());

    let races = lengths.zip(records).map(|(l, r)| {
        Race {
            length: l,
            record: r,
        }
    }).collect::<Vec<_>>();

    println!("{:?}", races);

    let mut res = Vec::with_capacity(races.len());
    for race in races {
        let mut win_count = 0;
        for speed in 0..=race.length {
            // {7, 9}
            // x + y = 7   [I]
            // x * y > 9   [II]
            let distance = speed * (race.length - speed);
            if distance > race.record {
                win_count += 1;
            }
        }
        res.push(win_count);
    }
    let csum = res.into_iter().reduce(|acc, n| acc * n).unwrap();
    println!("{csum}");
}
