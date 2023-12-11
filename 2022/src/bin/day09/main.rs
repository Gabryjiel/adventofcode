use std::{collections::HashSet, vec};

fn main() {
    task2();
}

fn task1() {
    let input =
        std::fs::read_to_string("./src/bin/day9/input.txt").expect("Unable to read the file");

    let (_head_positions, _tail_positions, counter) = input
        .split('\n')
        .map(|x| x.split_once(' ').expect("Canno split_once"))
        .map(|(direction, steps)| (direction, steps.parse::<u32>().expect("Cannot parse")))
        .flat_map(|(direction, steps)| {
            let mut vector = vec![];

            for _i in 0..steps {
                vector.push(direction);
            }

            vector
        })
        .fold((vec![(0, 0)], vec![(0, 0)], 1), |all, direction| {
            let mut head_positions = all.0.clone();
            let mut tail_positions = all.1.clone();
            let mut counter = all.2;

            let (last_head_position_x, last_head_position_y) =
                all.0.last().expect("No last head position");
            let (last_tail_position_x, last_tail_position_y) =
                all.1.last().expect("No last tail position");

            let new_head_position = match direction {
                "U" => (*last_head_position_x, *last_head_position_y + 1),
                "D" => (*last_head_position_x, *last_head_position_y - 1),
                "L" => (*last_head_position_x - 1, *last_head_position_y),
                "R" => (*last_head_position_x + 1, *last_head_position_y),
                _ => (*last_head_position_x, *last_head_position_y),
            };

            let distance_between_tail_and_head =
                (((*last_tail_position_x as f32) - (new_head_position.0 as f32)).powi(2)
                    + ((*last_tail_position_y as f32) - (new_head_position.1 as f32)).powi(2))
                .sqrt();

            let new_tail_position = if distance_between_tail_and_head >= 1.9 {
                (*last_head_position_x, *last_head_position_y)
            } else {
                (*last_tail_position_x, *last_tail_position_y)
            };

            let is_unique_tail_position = tail_positions
                .iter()
                .all(|(x, y)| *x != new_tail_position.0 || *y != new_tail_position.1);

            if is_unique_tail_position {
                counter += 1;
            }

            head_positions.push(new_head_position);
            tail_positions.push(new_tail_position);

            (head_positions, tail_positions, counter)
        });

    println!("{:?}", counter);
}

fn task2() {
    let input =
        std::fs::read_to_string("./src/bin/day9/example2.txt").expect("Unable to read the file");

    let (_head_positions, _tail_positions, _) = input
        .split('\n')
        .map(|x| x.split_once(' ').expect("Canno split_once"))
        .map(|(direction, steps)| (direction, steps.parse::<u32>().expect("Cannot parse")))
        .flat_map(|(direction, steps)| {
            let mut vector = vec![];

            for _i in 0..steps {
                vector.push(direction);
            }

            vector
        })
        .fold((vec![(0, 0)], vec![(0, 0)], 1), |all, direction| {
            let mut head_positions = all.0.clone();
            let mut tail_positions = all.1.clone();
            let mut counter = all.2;

            let (last_head_position_x, last_head_position_y) =
                all.0.last().expect("No last head position");
            let (last_tail_position_x, last_tail_position_y) =
                all.1.last().expect("No last tail position");

            let new_head_position = match direction {
                "U" => (*last_head_position_x, *last_head_position_y + 1),
                "D" => (*last_head_position_x, *last_head_position_y - 1),
                "L" => (*last_head_position_x - 1, *last_head_position_y),
                "R" => (*last_head_position_x + 1, *last_head_position_y),
                _ => (*last_head_position_x, *last_head_position_y),
            };

            let distance_between_tail_and_head =
                (((*last_tail_position_x as f32) - (new_head_position.0 as f32)).powi(2)
                    + ((*last_tail_position_y as f32) - (new_head_position.1 as f32)).powi(2))
                .sqrt();

            let new_tail_position = if distance_between_tail_and_head >= 2.1 {
                let x_dist = if last_head_position_x - last_tail_position_x > 0 {
                    1
                } else {
                    -1
                };
                let y_dist = if last_head_position_y - last_tail_position_y > 0 {
                    1
                } else {
                    -1
                };

                (
                    *last_tail_position_x + x_dist,
                    *last_tail_position_y + y_dist,
                )
            } else if distance_between_tail_and_head >= 1.9 {
                (*last_head_position_x, *last_head_position_y)
            } else {
                (*last_tail_position_x, *last_tail_position_y)
            };

            let is_unique_tail_position = tail_positions
                .iter()
                .all(|(x, y)| *x != new_tail_position.0 || *y != new_tail_position.1);

            if is_unique_tail_position {
                counter += 1;
            }

            head_positions.push(new_head_position);
            tail_positions.push(new_tail_position);

            (head_positions, tail_positions, counter)
        });

    let tail1 = get_tail_positions(_head_positions.clone());
    let tail2 = get_tail_positions(tail1.clone());
    let tail3 = get_tail_positions(tail2.clone());
    let tail4 = get_tail_positions(tail3.clone());
    let tail5 = get_tail_positions(tail4.clone());
    let tail6 = get_tail_positions(tail5.clone());
    let tail7 = get_tail_positions(tail6.clone());
    let tail8 = get_tail_positions(tail7.clone());
    let tail9 = get_tail_positions(tail8.clone());

    let mut vector = Vec::new();

    vector.push(_head_positions[5]);
    vector.push(tail1[5]);
    vector.push(tail2[5]);
    vector.push(tail3[5]);
    vector.push(tail4[5]);
    vector.push(tail5[5]);
    vector.push(tail6[5]);
    vector.push(tail7[5]);
    vector.push(tail8[5]);
    vector.push(tail9[5]);

    for y in -16..16 {
        for x in -16..16 {
            if vector.contains(&(x, -y)) {
                print!("#");
            } else {
                print!(".");
            }
        }
        println!();
    }

    println!("{:?}", vector);

    // display_positions(tail1);
    // display_positions(tail2);

    // println!("{:?}", _head_positions);
    // println!("{:?}", tail_positions);
    // println!("{:?}", get_tail_positions(_head_positions));
}

fn get_tail_positions(head_positions: Vec<(i32, i32)>) -> Vec<(i32, i32)> {
    let result = head_positions
        .iter()
        .fold(((0, 0), vec![(0, 0)]), |all, (x, y)| {
            let mut tail_positions = all.1.clone();

            let (last_head_position_x, last_head_position_y) = all.0;
            let (last_tail_position_x, last_tail_position_y) =
                tail_positions.last().expect("No last value for tail");

            let distance_between_tail_and_head = (((*last_tail_position_x as f32) - (*x as f32))
                .powi(2)
                + ((*last_tail_position_y as f32) - (*y as f32)).powi(2))
            .sqrt();

            let new_tail_position = if distance_between_tail_and_head >= 2.1 {
                let x_dist = if x - last_tail_position_x > 0 { 1 } else { -1 };
                let y_dist = if y - last_tail_position_y > 0 { 1 } else { -1 };

                (
                    *last_tail_position_x + x_dist,
                    *last_tail_position_y + y_dist,
                )
            } else if distance_between_tail_and_head >= 1.9 {
                (last_head_position_x, last_head_position_y)
            } else {
                (*last_tail_position_x, *last_tail_position_y)
            };

            tail_positions.push(new_tail_position);

            ((*x, *y), tail_positions)
        });

    result.1
}

fn display_positions(positions: Vec<(i32, i32)>) {
    let mut s: HashSet<(i32, i32)> = HashSet::new();

    for pos in positions {
        s.insert(pos);
    }

    for y in -16..16 {
        for x in -16..16 {
            if s.contains(&(x, -y)) {
                print!("#");
            } else {
                print!(".");
            }
        }
        println!();
    }
}
