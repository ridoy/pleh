package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"golang.design/x/clipboard"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: pleh <query>")
		os.Exit(1)
	}
	initClipboard()
	query := strings.Join(os.Args[1:], " ")
	command, err := generateShellCommand(query)
	if err != nil {
		fmt.Printf("Error generating shell command: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Generated command: %s\n", command)
	fmt.Println("Press [Enter] to run this command, [0] to copy to clipboard, or [1] to exit.")
	var response string
	fmt.Scanln(&response)
	if response == "" {
		executeCommand(command)
	} else if response == "0" {
		clipboard.Write(clipboard.FmtText, []byte(command))
		fmt.Println("Command copied to clipboard.")
	} else if response == "1" {
		return 
	}
}

func generateShellCommand(query string) (string, error) {
	prompt := fmt.Sprintf("Generate a shell command that accomplishes the following: %s. Output just the command with no explanation.", query)
	cmd := exec.Command("../../llama-cli", "-m", "../../model/llama-2-7b.Q8_0.gguf", "-p", prompt, "--log-disable")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(output)), nil
}

func executeCommand(command string) {
	cmd := exec.Command("bash", "-c", command)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error executing command: %v\n", err)
	}
	fmt.Println(string(output))
}

func initClipboard() {
	err := clipboard.Init()
	if err != nil {
		panic(err)
	}
}
