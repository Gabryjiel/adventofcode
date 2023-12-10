use std::collections::HashSet;

fn main() {
    task1();
    task2();
}

fn task1() {
    let input =
        std::fs::read_to_string("./src/bin/day6/input.txt").expect("Unable to read the file");

    let mut index = 0;
    for i in 0..(input.len() - 3) {
        let chunk = input.get(i..i + 4).unwrap().chars();
        let set: HashSet<char> = HashSet::from_iter(chunk);

        if set.len() == 4 {
            index = i + 4;
            break;
        }
    }

    println!("{:?}", index);
}

fn task2() {
    let input =
        std::fs::read_to_string("./src/bin/day6/input.txt").expect("Unable to read the file");

    let mut index = 0;
    for i in 0..(input.len() - 13) {
        let chunk = input.get(i..i + 14).unwrap().chars();
        let set: HashSet<char> = HashSet::from_iter(chunk);

        if set.len() == 14 {
            index = i + 14;
            break;
        }
    }

    println!("{:?}", index);
}
