fn main() {
    task1a();
    task1b();
}

pub fn task1a() {
    let file = std::fs::read_to_string("./src/bin/day1/input.txt");
    let content = file.unwrap_or(String::from(""));
    let result = content.split('\n');
    let mut calories = Vec::new();

    let mut sum = 0;
    for res in result {
        match res.parse::<i32>() {
            Ok(num) => sum += num,
            Err(_) => {
                calories.push(sum);
                sum = 0
            }
        }
    }

    let mut max = 0;
    for calory in calories {
        max = if calory > max { calory } else { max }
    }

    println!("{}", max)
}

fn task1b() {
    let file = std::fs::read_to_string("./src/bin/day1/input.txt");
    let content = file.unwrap_or(String::from(""));
    let result = content.split('\n');
    let mut calories = Vec::new();

    let mut sum = 0;
    for res in result {
        match res.parse::<i32>() {
            Ok(num) => sum += num,
            Err(_) => {
                calories.push(sum);
                sum = 0
            }
        }
    }

    calories.sort();
    calories.reverse();

    let mut sum = 0;
    for calory in &calories[0..=2] {
        sum += calory;
    }
    println!("{}", sum);
}
