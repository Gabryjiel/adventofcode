fn main() {
    task1();
    task2();
}

fn task1() {
    let input =
        std::fs::read_to_string("./src/bin/day4/input.txt").expect("Unable to read the file");

    let pairs = input
        .split('\n')
        .map(|x| x.split_once(','))
        .map(|x| x.unwrap())
        .map(|(left, right)| {
            let (left_begin, left_end) = left.split_once('-').unwrap();
            let (right_begin, right_end) = right.split_once('-').unwrap();

            let left_begin_number = left_begin.parse::<u32>().unwrap();
            let left_end_number = left_end.parse::<u32>().unwrap();
            let right_begin_number = right_begin.parse::<u32>().unwrap();
            let right_end_number = right_end.parse::<u32>().unwrap();

            if left_begin_number <= right_begin_number && left_end_number >= right_end_number {
                1
            } else if right_begin_number <= left_begin_number && right_end_number >= left_end_number
            {
                1
            } else {
                0
            }
        })
        .sum::<u32>();

    println!("{:?}", pairs);
}

fn task2() {
    let input =
        std::fs::read_to_string("./src/bin/day4/input.txt").expect("Unable to read the file");

    let pairs = input
        .split('\n')
        .map(|x| x.split_once(','))
        .map(|x| x.unwrap())
        // .map(|(left, right)| (left.split_once('-'), right.split_once('-')))
        // .map(|(left, right)| (left.unwrap(), right.unwrap()))
        // .map(|((left_begin, left_end), (right_begin, right_end))| (
        //   (left_begin.parse::<u32>(), left_end.parse::<u32>()),
        //   (right_begin.parse::<u32>(), right_end.parse::<u32>())
        // ))
        // .map(|((left_begin, left_end), (right_begin, right_end))| (
        //   (left_begin.unwrap(), left_end.unwrap()),
        //   (right_begin.unwrap(), right_end.unwrap())
        // ))
        // .map(|((left_begin, left_end), (right_begin, right_end))| (left_begin..=left_end, right_begin..=right_end))
        // .map(|(mut left_range, right_range)| left_range.any(|x| right_range.contains(&x)))
        // .map(|x| if x { 1 } else { 0 })
        // .sum::<usize>();
        .map(|(left, right)| {
            let (left_begin, left_end) = left.split_once('-').unwrap();
            let (right_begin, right_end) = right.split_once('-').unwrap();

            let left_begin_number = left_begin.parse::<u32>().unwrap();
            let left_end_number = left_end.parse::<u32>().unwrap();
            let right_begin_number = right_begin.parse::<u32>().unwrap();
            let right_end_number = right_end.parse::<u32>().unwrap();

            let mut left_range = left_begin_number..=left_end_number;
            let right_range = right_begin_number..=right_end_number;

            if left_range.any(|x| right_range.contains(&x)) {
                1
            } else {
                0
            }
        })
        .sum::<usize>();

    println!("{:?}", pairs);
}
