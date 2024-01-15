package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: govim <filename>")
		os.Exit(1)
	}

	fileName := os.Args[1]

	// Get the current environment variables
	envVars := os.Environ()

	// Use gnome-terminal to open Vim with the specified file in a new Zsh terminal
	cmd := exec.Command("gnome-terminal", "--", "zsh", "-c", fmt.Sprintf("vim %s; exec zsh", fileName))
	cmd.Env = envVars
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println("Error opening file in Vim:", err)
		os.Exit(1)
	}
}


