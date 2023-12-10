fn main() {
    task1();
    task2();
}

fn task1() {
    let input =
        std::fs::read_to_string("./src/bin/day10/input.txt").expect("Unable to read the file");

    let result = input.split('\n').fold((0, 1, 0), |acc, line| {
        let (cycles, value, sum) = acc;
        let mut signal_values_sum = sum;

        let (cycles_to_add, value_to_add) = if line == "noop" {
            (1, 0)
        } else {
            let (_command, value) = line.split_once(' ').expect("Canno split");
            let num = value.parse::<i32>().expect("Cannot parse");
            (2, num)
        };

        for i in 1..=cycles_to_add {
            if ((acc.0 + i) - 20) % 40 == 0 && cycles < 221 {
                signal_values_sum += value * (cycles + i)
            }
        }

        (
            cycles + cycles_to_add,
            value + value_to_add,
            signal_values_sum,
        )
    });

    println!("{:?}", result);
}

fn task2() {
    let input =
        std::fs::read_to_string("./src/bin/day10/input.txt").expect("Unable to read the file");

    let (_, _, display) = input
        .split('\n')
        .fold((0, 2, String::from("")), |acc, line| {
            let (cycles, sprite_position, display) = acc;

            let (cycles_to_add, value_to_add) = if line == "noop" {
                (1, 0)
            } else {
                let (_command, value) = line.split_once(' ').expect("Canno split");
                let num = value.parse::<i32>().expect("Cannot parse");
                (2, num)
            };

            let mut to_add = String::from("");
            for i in 1..=cycles_to_add {
                let mut sign = String::from(".");

                for j in -1..=1 {
                    if (cycles + i + j) % 40 == sprite_position {
                        sign = String::from("#");
                    }
                }

                to_add += &sign;
            }

            (
                cycles + cycles_to_add,
                sprite_position + value_to_add,
                display + &to_add,
            )
        });

    let res = display
        .as_bytes()
        .chunks(40)
        .map(|buf| unsafe { std::str::from_utf8_unchecked(buf) })
        .collect::<Vec<&str>>();

    for line in res {
        println!("{:?}", line);
    }
}
