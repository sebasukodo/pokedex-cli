package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	cfg := &config{}

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()

		inputArray := strings.Fields(strings.ToLower(input))
		if len(inputArray) == 0 {
			continue
		}

		if v, ok := commands[inputArray[0]]; ok {
			if err := v.callback(cfg, inputArray); err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}

	}
}

func cleanInput(text string) []string {
	return strings.Fields(text)
}
