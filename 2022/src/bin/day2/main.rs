fn main() {
    let input = std::fs::read_to_string("./src/bin/day2/input.txt").unwrap();
    let sum1 = input.split('\n').map(task1::calculate_score).sum::<i32>();
    let sum2 = input.split('\n').map(task2::calculate_score).sum::<i32>();

    println!("1: {}", sum1);
    println!("2: {}", sum2);
}

pub mod task1 {
    pub fn calculate_score(round: &str) -> i32 {
        match_result(round) + played_item(round)
    }

    fn match_result(round: &str) -> i32 {
        if round.contains("A X") || round.contains("B Y") || round.contains("C Z") {
            3
        } else if round.contains("A Y") || round.contains("B Z") || round.contains("C X") {
            6
        } else {
            0
        }
    }

    fn played_item(round: &str) -> i32 {
        if round.ends_with('X') {
            1
        } else if round.ends_with('Y') {
            2
        } else {
            3
        }
    }
}

pub mod task2 {
    pub fn calculate_score(round: &str) -> i32 {
        match_result(round) + item_to_pay(round)
    }

    fn match_result(round: &str) -> i32 {
        if round.contains('X') {
            0
        } else if round.contains('Y') {
            3
        } else if round.contains('Z') {
            6
        } else {
            0
        }
    }

    fn item_to_pay(round: &str) -> i32 {
        if round.starts_with('A') {
            if round.ends_with('X') {
                3
            } else if round.ends_with('Y') {
                1
            } else if round.ends_with('Z') {
                2
            } else {
                0
            }
        } else if round.starts_with('B') {
            if round.ends_with('X') {
                1
            } else if round.ends_with('Y') {
                2
            } else if round.ends_with('Z') {
                3
            } else {
                0
            }
        } else if round.starts_with('C') {
            if round.ends_with('X') {
                2
            } else if round.ends_with('Y') {
                3
            } else if round.ends_with('Z') {
                1
            } else {
                0
            }
        } else {
            0
        }
    }
}
