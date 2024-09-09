use std::error::Error;
use std::str::FromStr;
use std::num::ParseIntError;
use std::fmt;
use std::fs;
use std::collections::HashMap;

#[derive(Copy, Clone, Debug, PartialEq, Eq, Hash, PartialOrd, Ord)]
enum Card {
    J = 1, N2 = 2, N3, N4, N5, N6, N7, N8, N9, T, Q, K, A
}

type Cards = [Card; 5];

// Note that we give the strength a reference to the cards, so that Ord can do its comparison first
// on the strength, and if the strengths are both equal, then on the card values.
#[derive(Debug, PartialEq, Eq, PartialOrd, Ord)]
enum Strength<'a> {
    HighCard(&'a Cards),
    OnePair(&'a Cards),
    TwoPair(&'a Cards),
    ThreeOfAKind(&'a Cards),
    FullHouse(&'a Cards),
    FourOfAKind(&'a Cards),
    FiveOfAKind(&'a Cards),
}

#[derive(Debug)]
struct Hand {
    cards: Cards,
    bid: u32,
}

#[derive(PartialEq, Debug)]
enum HandParseError {
    InvalidFormat,
    InvalidCard(String),
    ParseInt(ParseIntError),
}

impl Error for HandParseError {}

impl From<ParseIntError> for HandParseError {
    fn from(e: ParseIntError) -> Self {
        Self::ParseInt(e)
    }
}

impl fmt::Display for HandParseError {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        match *self {
            HandParseError::InvalidFormat => write!(f, "invalid format"),
            HandParseError::InvalidCard(ref s) => write!(f, "invalid card {s}"),
            HandParseError::ParseInt(ref e) => write!(f, "{e}"),
        }
    }
}

impl FromStr for Hand {
    type Err = HandParseError;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let mut parts = s.split(' ');
        let (Some(hand), Some(bid), None) = (parts.next(), parts.next(), parts.next()) else {
            return Err(HandParseError::InvalidFormat);
        };
        if hand.len() != 5 {
            return Err(HandParseError::InvalidFormat);
        }
        let mut cards: Cards = [Card::A; 5];
        for i in 0..5 {
            cards[i] = hand[i..i+1].parse::<Card>()?;
        }
        let bid = bid.parse::<u32>()?;
        Ok(Hand{cards, bid})
    }
}

impl FromStr for Card {
    type Err = HandParseError;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        match s {
            "A" => Ok(Card::A),
            "K" => Ok(Card::K),
            "Q" => Ok(Card::Q),
            "J" => Ok(Card::J),
            "T" => Ok(Card::T),
            "9" => Ok(Card::N9),
            "8" => Ok(Card::N8),
            "7" => Ok(Card::N7),
            "6" => Ok(Card::N6),
            "5" => Ok(Card::N5),
            "4" => Ok(Card::N4),
            "3" => Ok(Card::N3),
            "2" => Ok(Card::N2),
            s => Err(HandParseError::InvalidCard(s.to_string()))
        }
    }
}

impl Hand {
    fn strength(&self) -> Strength {
        let mut m: HashMap<Card, u32> = HashMap::with_capacity(5);
        for c in self.cards {
            *m.entry(c).or_default() += 1;
        }
        let mut vals = m.values().collect::<Vec<_>>();
        vals.sort(); // so that we don't have to consider multiple cases like (4, 1) | (1, 4)
        if m.contains_key(&Card::J) {
            let jc = m.get(&Card::J).unwrap();
            match vals.len() {
                1 | 2 => Strength::FiveOfAKind(&self.cards),
                3 => {
                    match (vals[0], vals[1], vals[2]) {
                        (1, 1, 3) => Strength::FourOfAKind(&self.cards),
                        (1, 2, 2) if *jc == 2 => Strength::FourOfAKind(&self.cards),
                        (1, 2, 2) if *jc == 1 => Strength::FullHouse(&self.cards),
                        _ => unreachable!(),
                    }
                },
                4 => Strength::ThreeOfAKind(&self.cards),
                5 => Strength::OnePair(&self.cards),
                _ => unreachable!(),
            }
        } else {
            match vals.len() {
                1 => Strength::FiveOfAKind(&self.cards),
                2 => {
                    match (vals[0], vals[1]) {
                        (1, 4) => Strength::FourOfAKind(&self.cards),
                        (2, 3) => Strength::FullHouse(&self.cards),
                        _ => unreachable!(),
                    }
                },
                3 => {
                    match (vals[0], vals[1], vals[2]) {
                        (1, 1, 3) => Strength::ThreeOfAKind(&self.cards),
                        (1, 2, 2) => Strength::TwoPair(&self.cards),
                        _ => unreachable!(),
                    }
                },
                4 => Strength::OnePair(&self.cards),
                5 => Strength::HighCard(&self.cards),
                _ => unreachable!(),
            }
        }
    }
}

fn main() -> Result<(), Box<dyn Error>> {
    //let input = fs::read_to_string("input.test")?;
    let input = fs::read_to_string("input")?;
    let mut hands = input.lines().map(Hand::from_str).collect::<Result<Vec<Hand>, HandParseError>>()?;

    hands.sort_by(|a, b| a.strength().cmp(&b.strength()));

    let total: u32 = hands.iter().enumerate().map(|(i, hand)| (i + 1) as u32 * hand.bid).sum();

    println!("{total}"); // 250384185

    Ok(())
}
