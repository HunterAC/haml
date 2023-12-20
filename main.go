package main

import (
	"fmt"
	"os"
	"github.com/HunterAC/haml/lexer"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("No file provided. Exiting.")
		return
	}

	filePath := args[0]
	fileBytes, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file at: " + filePath)
		fmt.Printf("Error message: \n%s\n", err.Error())
	}

	fileString := string(fileBytes)
	tokens, _ := lexer.Lex([]rune(fileString))
	for _, t := range tokens {
		fmt.Printf("%v, %v, %v\n", t.Category, t.Value, t.Level)
	}
}