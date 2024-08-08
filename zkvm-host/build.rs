use std::process::Command;

use sp1_helper::{build_program_with_args, BuildArgs};

fn build_native_program(program: &str) {
    let status = Command::new("cargo")
        .args([
            "build",
            "--workspace",
            "--bin",
            program,
            "--profile",
            "release-client-lto",
        ])
        .status()
        .expect("Failed to execute cargo build command");

    if !status.success() {
        panic!("Failed to build {}", program);
    }

    println!(
        "cargo:warning={} built with release-client-lto profile",
        program
    );
}

fn build_zkvm_program(program: &str) {
    build_program_with_args(
        &format!("../{}", program),
        BuildArgs {
            ignore_rust_version: true,
            elf_name: format!("{}-elf", program),
            ..Default::default()
        },
    );
}

fn main() {
    let programs = vec!["zkvm-client", "validity-client"];

    for program in programs {
        build_native_program(program);
        build_zkvm_program(program);
    }
}