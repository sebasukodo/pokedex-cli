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
	callback    func() error
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
	}
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	text := "Welcome to the Pokedex!\nUsage:\n\n"
	for key, value := range commands {
		text += fmt.Sprintf("%v: %v\n", key, value.description)
	}
	text = strings.TrimSuffix(text, "\n")
	fmt.Println(text)
	return nil
}
