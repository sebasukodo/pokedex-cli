package main

import (
	"fmt"
	"os"
	"strings"
)

// initialize empty map, because CommandHelp() is relying on commands
var commands map[string]cliCommand

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	next     int
	previous int
}

func init() {
	commands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Display next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display previous 20 locations",
			callback:    commandMapB,
		},
	}
}

func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config) error {
	text := "Welcome to the Pokedex!\nUsage:\n\n"
	for key, value := range commands {
		text += fmt.Sprintf("%v: %v\n", key, value.description)
	}
	text = strings.TrimSuffix(text, "\n")
	fmt.Println(text)
	return nil
}

func commandMap(cfg *config) error {
	id := cfg.next

	for i := 1; i <= id; i++ {
		data := getAPIInfo(i)
		fmt.Println(data.Name)
	}

	cfg.next = id + 20
	cfg.previous = id

	return nil
}

func commandMapB(cfg *config) error {
	id := cfg.previous
	if id == 0 {
		fmt.Println("you're on the first page")
		return nil
	}

	for i := 1; i <= id; i++ {
		data := getAPIInfo(i)
		fmt.Println(data.Name)
	}

	cfg.next = id
	cfg.previous = id - 20

	return nil

}
