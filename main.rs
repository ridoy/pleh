use clap::{Arg, Command};
use ggml::model::GgmlModel;

fn main() {
    let matches = Command::new("AI Shell Utility")
        .version("1.0")
        .author("Your Name")
        .about("Helps with shell operations using AI")
        .subcommand(Command::new("suggest")
            .about("Suggests a shell command based on input")
            .arg(Arg::new("COMMAND")
                .help("The natural language command")
                .required(true)
                .index(1)))
        .subcommand(Command::new("run")
            .about("Executes a shell command")
            .arg(Arg::new("COMMAND")
                .help("The shell command to run")
                .required(true)
                .index(1)))
        .get_matches();

    match matches.subcommand() {
        Some(("suggest", sub_m)) => {
            let command = sub_m.value_of("COMMAND").unwrap();
            suggest_command(command);
        }
        Some(("run", sub_m)) => {
            let command = sub_m.value_of("COMMAND").unwrap();
            run_command(command);
        }
        _ => println!("Use --help for more information."),
    }
}



fn run_command(command: &str) {
    println!("Running: {}", command);
    std::process::Command::new("sh")
        .arg("-c")
        .arg(command)
        .spawn()
        .expect("Failed to execute command");
}


fn load_model() -> GgmlModel {
    // Load the model from file
    GgmlModel::from_file("path/to/your/model")
        .expect("Failed to load model")
}

fn suggest_command(command: &str) {
    let model = load_model();
    let prompt = format!("Convert this into a shell command: {}", command);
    let response = model.generate_response(&prompt);
    println!("Interpreted command: {}", response);
}
