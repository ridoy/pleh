package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "os"
    "strings"
    "bufio"
    "os/exec"

    "github.com/joho/godotenv"
)

const openAIAPIURL = "https://api.openai.com/v1/chat/completions"

type OpenAIRequest struct {
    Model    string    `json:"model"`
    Messages []Message `json:"messages"`
}

type Message struct {
    Role    string `json:"role"`
    Content string `json:"content"`
}

type OpenAIResponse struct {
    Choices []struct {
        Message Message `json:"message"`
    } `json:"choices"`
}

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    apiKey := os.Getenv("OPENAI_API_KEY")
    if apiKey == "" {
        log.Fatal("OPENAI_API_KEY is not set in the environment")
    }

    if len(os.Args) < 2 || os.Args[0] != "pleh" {
        fmt.Println("Usage: pleh <your query>")
        os.Exit(1)
    }

    command := strings.Join(os.Args[2:], " ")
    for {
        response, err := getOpenAIResponse(apiKey, command)
        if err != nil {
            log.Printf("Error getting response from OpenAI: %v\n", err)
        } else {
            fmt.Println(response)
            fmt.Print("Press [Enter] to use this command now, or [0] to retry")
            scanner := bufio.NewScanner(os.Stdin)
            scanner.Scan()
            input := scanner.Text()
            if input == "" {
                executeCommand(response)
                break
            } else if input == "0" {
                continue
            }
        }
    }
}

func getOpenAIResponse(apiKey, command string) (string, error) {
    // Add an instruction to the prompt to ensure the response is a shell command
    prompt := fmt.Sprintf("Please respond with a valid shell command to the following query:\n\n%s", command)

    requestPayload := OpenAIRequest{
        Model: "gpt-3.5-turbo",
        Messages: []Message{
            {
                Role:    "user",
                Content: prompt,
            },
        },
    }

    payloadBytes, err := json.Marshal(requestPayload)
    if err != nil {
        return "", err
    }

    req, err := http.NewRequest("POST", openAIAPIURL, bytes.NewBuffer(payloadBytes))
    if err != nil {
        return "", err
    }

    req.Header.Set("Authorization", "Bearer "+apiKey)
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        body, _ := ioutil.ReadAll(resp.Body)
        return "", fmt.Errorf("OpenAI API error: %v", string(body))
    }

    var openAIResponse OpenAIResponse
    err = json.NewDecoder(resp.Body).Decode(&openAIResponse)
    if err != nil {
        return "", err
    }

    if len(openAIResponse.Choices) == 0 {
        return "No response from OpenAI", nil
    }

    return openAIResponse.Choices[0].Message.Content, nil
}


func executeCommand(command string) {
    args := strings.Fields(command)
    cmd := exec.Command(args[0], args[1:]...)

    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    err := cmd.Run()
    if err != nil {
        log.Printf("Error executing command: %v\n", err)
    }
}
