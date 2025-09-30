package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()

		if strings.ToLower(input) == "q" || strings.ToLower(input) == "quit" {
			break
		}

		words := strings.Fields(strings.ToLower(input))
		fmt.Println("Your command was:", words[0])

	}
}

func cleanInput(text string) []string {
	return strings.Fields(text)
}
