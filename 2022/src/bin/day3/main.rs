use std::collections::HashSet;

fn main() {
    task1();
    task2();
}

fn task1() {
    let input = std::fs::read_to_string("./src/bin/day3/input.txt").unwrap();
    let rows = input
        .split('\n')
        .map(|rucksack| rucksack.split_at(rucksack.len() / 2))
        .map(|(left, right)| {
            left.chars()
                .collect::<HashSet<_>>()
                .intersection(&right.chars().collect::<HashSet<_>>())
                .filter(|c| c != &&'\0')
                .copied()
                .map(get_priority)
                .sum::<u32>()
        })
        .sum::<u32>();

    println!("{}", rows);
}

fn task2() {
    let input = std::fs::read_to_string("./src/bin/day3/input.txt").unwrap();
    let rows = input.lines();
    let mut grouped_into_3: Vec<Vec<&str>> = vec![];

    for (index, value) in rows.enumerate() {
        let group_index = index / 3;
        let entry_index = index % 3;

        if entry_index == 0 {
            grouped_into_3.push(vec![value]);
        } else {
            grouped_into_3[group_index].push(value);
        }
    }

    let mut sum = 0;
    for group in grouped_into_3 {
        for char in group[0].chars() {
            if group[1].contains(char) && group[2].contains(char) {
                sum += get_priority(char);
                break;
            }
        }
    }

    println!("{}", sum);
}

fn get_priority(input: char) -> u32 {
    if input.is_lowercase() {
        input as u32 - 96
    } else if input.is_uppercase() {
        input as u32 - 65 + 27
    } else {
        0
    }
}
