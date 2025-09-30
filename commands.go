package main

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/sebasukodo/pokedex-cli/internal/pokeapi"
)

// initialize empty map, because CommandHelp() is relying on commands
var commands map[string]cliCommand

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	nextURL     *string
	previousURL *string
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

	locationData, err := pokeapi.ListLocations(cfg.nextURL)
	if err != nil {
		return err
	}

	cfg.nextURL = locationData.Next
	cfg.previousURL = locationData.Previous

	for _, location := range locationData.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandMapB(cfg *config) error {
	if cfg.previousURL == nil {
		return errors.New("you're on the first page")
	}

	locationData, err := pokeapi.ListLocations(cfg.previousURL)
	if err != nil {
		return err
	}

	cfg.nextURL = locationData.Next
	cfg.previousURL = locationData.Previous

	for _, location := range locationData.Results {
		fmt.Println(location.Name)
	}

	return nil

}
