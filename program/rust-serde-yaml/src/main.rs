use serde_yaml::Value;
use std::error::Error;
use std::fs;
use std::env;

fn main() -> Result<(), Box<dyn Error>> {
    let args: Vec<String> = env::args().collect();
    let path = &args[1];
    let content = fs::read_to_string(path)?;
    let yaml: Value = serde_yaml::from_str(&content)?;
    println!("{:#?}", yaml);

    Ok(())
}
