use orange_trees::{Node, Tree};

fn main() {
    task1();
    task2();
}

fn task1() {
    let input =
        std::fs::read_to_string("./src/bin/day7/input.txt").expect("Unable to read the file");

    let tree: Tree<String, u32> = Tree::new(Node::new(String::from("/"), 0));

    let mut dirs: Vec<String> = vec![String::from("/")];

    let (new_tree, _) = input
        .split('\n')
        .fold((tree, String::from("")), |acc, cur| {
            let (tree, position) = acc;

            if cur.starts_with("$ cd ..") {
                go_up(tree, position, cur)
            } else if cur.starts_with("$ cd /") {
                go_home(tree)
            } else if cur.starts_with("$ cd") {
                go_to(tree, position, cur)
            } else if cur.starts_with("$ ls") {
                (tree, position)
            } else if cur.starts_with("dir") {
                let (res_tree, res_position, res_dir) = create_dir(tree, position, cur);
                dirs.push(res_dir);
                (res_tree, res_position)
            } else {
                create_file(tree, position, cur)
            }
        });

    let dirs_with_size = dirs
        .iter()
        .map(|x| calc_node_value(new_tree.root().query(x).expect("No id")))
        .filter(|size| size < &100000)
        .sum::<u32>();

    println!("{:?}, {:?}", dirs_with_size, dirs);
}

fn task2() {
    let input =
        std::fs::read_to_string("./src/bin/day7/input.txt").expect("Unable to read the file");

    let tree: Tree<String, u32> = Tree::new(Node::new(String::from("/"), 0));

    let total_space = 70000000;
    let needed_space = 30000000;

    let mut dirs: Vec<String> = vec![String::from("/")];

    let (new_tree, _) = input
        .split('\n')
        .fold((tree, String::from("")), |acc, cur| {
            let (tree, position) = acc;

            if cur.starts_with("$ cd ..") {
                go_up(tree, position, cur)
            } else if cur.starts_with("$ cd /") {
                go_home(tree)
            } else if cur.starts_with("$ cd") {
                go_to(tree, position, cur)
            } else if cur.starts_with("$ ls") {
                (tree, position)
            } else if cur.starts_with("dir") {
                let (res_tree, res_position, res_dir) = create_dir(tree, position, cur);
                dirs.push(res_dir);
                (res_tree, res_position)
            } else {
                create_file(tree, position, cur)
            }
        });

    let used_space = calc_node_value(new_tree.root());
    let free_space = total_space - used_space;
    let min_space_to_remove = needed_space - free_space;

    let dirs_with_size = dirs
        .iter()
        .map(|x| calc_node_value(new_tree.root().query(x).expect("No id")))
        .filter(|size| size > &min_space_to_remove)
        .min()
        .expect("No minimum");

    println!("{:?}", dirs_with_size);
}

fn calc_node_value(node: &Node<String, u32>) -> u32 {
    let value = *node.value();

    if value != 0 {
        value
    } else {
        return value + node.children().iter().map(calc_node_value).sum::<u32>();
    }
}

fn go_up(tree: Tree<String, u32>, position: String, _cur: &str) -> (Tree<String, u32>, String) {
    let address = position.split(':').collect::<Vec<_>>();
    let ad = &address[0..address.len() - 1].join(":");

    let mut new_position = String::new();
    new_position.push_str(ad);

    (tree, new_position)
}

fn go_home(tree: Tree<String, u32>) -> (Tree<String, u32>, String) {
    let new_position = String::from("/");
    (tree, new_position)
}

fn go_to(tree: Tree<String, u32>, position: String, cur: &str) -> (Tree<String, u32>, String) {
    let id = cur.split_whitespace().last().expect("No last element");

    let mut new_position = position;
    new_position.push(':');
    new_position.push_str(id);

    (tree, new_position)
}

fn create_dir(
    mut tree: Tree<String, u32>,
    position: String,
    cur: &str,
) -> (Tree<String, u32>, String, String) {
    let name = cur.split(' ').last().expect("No name for dir");

    let current_node = tree
        .root_mut()
        .query_mut(&position)
        .expect("Unable to find current node");

    let mut new_position = position.clone();
    new_position.push(':');
    new_position.push_str(name);
    let res = new_position.clone();

    current_node.add_child(Node::new(new_position, 0));

    (tree, position, res)
}

fn create_file(
    mut tree: Tree<String, u32>,
    position: String,
    cur: &str,
) -> (Tree<String, u32>, String) {
    let (size, name) = cur.split_once(' ').expect("Invalid split");
    let parsed_size = size.parse::<u32>().expect("Error parsing");

    let current_node = tree
        .root_mut()
        .query_mut(&position)
        .expect("Unable to find current node1");

    let mut new_position = position.clone();
    new_position.push(':');
    new_position.push_str(name);

    current_node.add_child(Node::new(new_position, parsed_size));

    (tree, position)
}
