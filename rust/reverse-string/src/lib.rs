pub fn reverse(input: &str) -> String {
    let characters: Vec<char> = input.chars().collect();
    let mut reversed = String::new();

    if characters.len() > 0 {
        for i in (0.. characters.len()).rev() {
            reversed.push(characters[i]);
        }
    }

    return reversed;
}
