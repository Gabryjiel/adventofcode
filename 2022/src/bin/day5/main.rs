fn main() {
    task1();
    task2();
}

fn task1() {
    let input =
        std::fs::read_to_string("./src/bin/day5/input.txt").expect("Unable to read the file");

    let (data, commands) = input.split_once("\n\n").unwrap();

    let no_of_lines = 36;
    let mut crates_final: Vec<Vec<char>> = vec![];

    for _ in 1..no_of_lines {
        crates_final.push(vec![]);
    }

    data.split('\n')
        .map(|x| x.chars().enumerate())
        .for_each(|x| {
            x.filter(|(_, value)| value.is_alphabetic())
                .for_each(|(index, value)| crates_final[index].push(value))
        });

    let crates_parsed = crates_final
        .iter()
        .filter(|x| !x.is_empty())
        .cloned()
        .collect::<Vec<_>>();

    let commands_parsed = commands
        .split('\n')
        .map(|x| String::from_iter(x.chars().filter(|x| !x.is_alphabetic())))
        .map(|x| String::from_iter(x.get(1..)))
        .map(|x| {
            x.split(' ')
                .map(|x| x.parse::<usize>().unwrap_or(0))
                .filter(|x| x > &0)
                .collect::<Vec<_>>()
        })
        .map(|x| (x[0], x[1] - 1, x[2] - 1))
        .collect::<Vec<_>>();

    let mut state = crates_parsed;
    for (amount, start, end) in commands_parsed {
        let crates_to_move = state[start][0..amount]
            .iter()
            .collect::<String>()
            .chars()
            .collect::<Vec<_>>();

        let start_stack = state[start].clone().get(amount..).unwrap().to_vec();
        let mut end_stack = state[end].clone();

        end_stack.reverse();
        crates_to_move.iter().for_each(|x| end_stack.push(*x));
        end_stack.reverse();

        state[start] = start_stack;
        state[end] = end_stack;
    }

    let word = state.iter().map(|x| x[0]).collect::<String>();

    println!("{:?}", word);
}

fn task2() {
    let input =
        std::fs::read_to_string("./src/bin/day5/input.txt").expect("Unable to read the file");

    let (data, commands) = input.split_once("\n\n").unwrap();

    let no_of_lines = 36;
    let mut crates_final: Vec<Vec<char>> = vec![];

    for _ in 1..no_of_lines {
        crates_final.push(vec![]);
    }

    data.split('\n')
        .map(|x| x.chars().enumerate())
        .for_each(|x| {
            x.filter(|(_, value)| value.is_alphabetic())
                .for_each(|(index, value)| crates_final[index].push(value))
        });

    let crates_parsed = crates_final
        .iter()
        .filter(|x| !x.is_empty())
        .cloned()
        .collect::<Vec<_>>();

    let commands_parsed = commands
        .split('\n')
        .map(|x| String::from_iter(x.chars().filter(|x| !x.is_alphabetic())))
        .map(|x| String::from_iter(x.get(1..)))
        .map(|x| {
            x.split(' ')
                .map(|x| x.parse::<usize>().unwrap_or(0))
                .filter(|x| x > &0)
                .collect::<Vec<_>>()
        })
        .map(|x| (x[0], x[1] - 1, x[2] - 1))
        .collect::<Vec<_>>();

    let mut state = crates_parsed;
    for (amount, start, end) in commands_parsed {
        let mut crates_to_move = state[start][0..amount]
            .iter()
            .collect::<String>()
            .chars()
            .collect::<Vec<_>>();

        let start_stack = state[start].clone().get(amount..).unwrap().to_vec();
        let mut end_stack = state[end].clone();

        crates_to_move.reverse();
        end_stack.reverse();
        end_stack.append(&mut crates_to_move);
        end_stack.reverse();

        state[start] = start_stack;
        state[end] = end_stack;
    }

    let word = state.iter().map(|x| x[0]).collect::<String>();

    println!("{:?}", word);
}
