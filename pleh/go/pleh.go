package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

const llamaModelPath = ""
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: pleh <query>")
		os.Exit(1)
	}
	query := strings.Join(os.Args[1:], " ")
	model, err := llama.LoadModel(llamaModelPath)
	if err != nil {
		log.Fatalf("Error loading LLaMA model: %v", err)
	}
	command, err := generateShellCommand(model, query)
	if err != nil {
		fmt.Printf("Error generating shell command: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Generated command: %s\n", command)
	fmt.Print("Press [Enter] to run this command, or [0] to exit.")
	var response string
	fmt.Scanln(&response)
	if response == "" {
		executeCommand(command)
	} else if response == "0" {
		return
	}
}

func generateShellCommand(model *llama.Model, query string) (string, error) {
	prompt := fmt.Sprintf("Generate a shell command that accomplishes the following: %s. Output just the command with no explanation.", query)
	response, err := model.Predict(prompt)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(response), nil
}

func executeCommand(command string) {
	cmd := exec.Command("bash", "-c", command)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error executing command: %v\n", err)
	}
	fmt.Println(string(output))
}

