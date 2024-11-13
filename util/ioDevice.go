package util

import (
	"bufio"
	"fmt"
	"os"

	"github.com/google/generative-ai-go/genai"
)

func PrintResponse(resp *genai.GenerateContentResponse) {
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				fmt.Println(part)
			}
		}
	}
	//fmt.Println("---")
}

func GetStdIn() (string, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter text: ")

	input, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("error reading input: %w", err) // Return an error if reading fails
	}

	return input, nil
}
