fn main() {
    task1();
    task2();
}

fn task1() {
    let input =
        std::fs::read_to_string("./src/bin/day8/input.txt").expect("Unable to read the file");

    let me = input
        .split('\n')
        .map(|x| {
            x.chars()
                .map(|c| c.to_digit(10).expect("Cannot parse"))
                .collect::<Vec<_>>()
        })
        .collect::<Vec<_>>();

    let result = me
        .iter()
        .enumerate()
        .map(|(y, line)| {
            let mut sum = 0;

            for (x, _) in line.iter().enumerate() {
                let is_tree_visible = is_visible(&me, &x, &y);

                if is_tree_visible {
                    sum += 1;
                } else {
                    // println!("{}, {}", x, y);
                }
            }

            sum
        })
        .sum::<u32>();

    println!("{:?}", result);
}

fn task2() {
    let input =
        std::fs::read_to_string("./src/bin/day8/input.txt").expect("Unable to read the file");

    let me = input
        .split('\n')
        .map(|x| {
            x.chars()
                .map(|c| c.to_digit(10).expect("Cannot parse"))
                .collect::<Vec<_>>()
        })
        .collect::<Vec<_>>();

    let result = me
        .iter()
        .enumerate()
        .map(|(y, line)| {
            let mut max = 0;

            for (x, _) in line.iter().enumerate() {
                let tree_score = calculate_viewing_score(&me, &x, &y);

                if tree_score > max {
                    max = tree_score;
                }
            }

            max
        })
        .max();

    println!("{:?}", result);
}

fn calculate_viewing_score(forest: &Vec<Vec<u32>>, x: &usize, y: &usize) -> usize {
    let value = forest[*y][*x];

    if *x == 0 || *y == 0 || *x == forest.len() - 1 || *y == forest[0].len() - 1 {
        return 0;
    }

    let mut left = forest[*y]
        .iter()
        .enumerate()
        .filter(|(index, _)| index < x)
        .map(|(_index, value)| *value)
        .collect::<Vec<_>>();

    left.reverse();

    let right = forest[*y]
        .iter()
        .enumerate()
        .filter(|(index, _)| index > x)
        .map(|(_index, value)| *value)
        .collect::<Vec<_>>();

    let binding = forest
        .iter()
        .enumerate()
        .map(|(_index, value)| value[*x])
        .collect::<Vec<_>>();

    let mut up = binding[0..*y].iter().collect::<Vec<_>>().clone();
    up.reverse();

    let down = &binding[*y + 1..];

    let (left_score, _) = left
        .iter()
        .enumerate()
        .find(|(_index, x)| x >= &&value)
        .unwrap_or((left.len() - 1, &0));
    let (right_score, _) = right
        .iter()
        .enumerate()
        .find(|(_index, x)| x >= &&value)
        .unwrap_or((right.len() - 1, &0));
    let (up_score, _) = up
        .iter()
        .enumerate()
        .find(|(_index, x)| x >= &&&value)
        .unwrap_or((up.len() - 1, &&0));
    let (down_score, _) = down
        .iter()
        .enumerate()
        .find(|(_index, x)| x >= &&value)
        .unwrap_or((down.len() - 1, &0));

    (left_score + 1) * (right_score + 1) * (up_score + 1) * (down_score + 1)
}

fn is_visible(forest: &Vec<Vec<u32>>, x: &usize, y: &usize) -> bool {
    let value = forest[*y][*x];

    if *x == 0 || *y == 0 || *x == forest.len() - 1 || *y == forest[0].len() - 1 {
        return true;
    }

    let left = forest[*y]
        .iter()
        .enumerate()
        .filter(|(index, _)| index < x)
        .map(|(_index, value)| *value)
        .collect::<Vec<_>>();

    let right = forest[*y]
        .iter()
        .enumerate()
        .filter(|(index, _)| index > x)
        .map(|(_index, value)| *value)
        .collect::<Vec<_>>();

    let binding = forest
        .iter()
        .enumerate()
        .map(|(_index, value)| value[*x])
        .collect::<Vec<_>>();
    let up = &binding[0..*y];
    let down = &binding[*y + 1..];

    let value = left.iter().any(|x| *x >= value)
        && right.iter().any(|x| *x >= value)
        && up.iter().any(|x| *x >= value)
        && down.iter().any(|x| *x >= value);

    !value
}
