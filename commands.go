package main

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/sebasukodo/pokedex-cli/internal/pokeapi"
)

// initialize empty maps
var commands map[string]cliCommand
var dex pokeapi.Pokedex

type cliCommand struct {
	name        string
	description string
	callback    func(*config, []string) error
}

type config struct {
	nextURL     *string
	previousURL *string
}

func init() {
	dex = pokeapi.NewPokedex()

	commands = map[string]cliCommand{
		"catch": {
			name:        "catch <Pokemon>",
			description: "Try to catch a Pokemon",
			callback:    commandCatch,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"explore": {
			name:        "explore <location-name/id>",
			description: "Display all available Pokemon in given area",
			callback:    commandExplore,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"inspect": {
			name:        "inspect <Pokemon>",
			description: "Display stats to one of your caught Pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "List all caught Pokemon",
			callback:    commandPokedex,
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

func commandCatch(cfg *config, input []string) error {
	if len(input) < 2 {
		return fmt.Errorf("usage: catch <Pokemon>")
	}

	pokemon, ok, err := pokeapi.CatchPokemon(input[1])
	if err != nil {
		return err
	}

	fmt.Println("Throwing a Pokeball at", input[1], "...")

	if !ok {
		fmt.Println(input[1], "escaped!")
	} else {
		dex.Add(pokemon)
		fmt.Println(input[1], "was caught!")
	}

	return nil
}

func commandExit(cfg *config, input []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandExplore(cfg *config, input []string) error {
	if len(input) < 2 {
		return fmt.Errorf("usage: explore <location-area>")
	}

	specificLocationData, err := pokeapi.ExploreLocation(input[1])
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %v...\n", specificLocationData.Location.Name)
	fmt.Println("Found Pokemon:")
	var text string
	for _, pokemon := range specificLocationData.PokemonEncounters {
		text += fmt.Sprintf(" - %v\n", pokemon.Pokemon.Name)
	}
	fmt.Println(strings.TrimSuffix(text, "\n"))

	return nil

}

func commandHelp(cfg *config, input []string) error {
	text := "Welcome to the Pokedex!\nUsage:\n\n"
	for key, value := range commands {
		text += fmt.Sprintf("%v: %v\n", key, value.description)
	}
	text = strings.TrimSuffix(text, "\n")
	fmt.Println(text)
	return nil
}

func commandInspect(cfg *config, input []string) error {

	if len(input) < 2 {
		return fmt.Errorf("usage: inspect <Caught Pokemon>")
	}

	if err := dex.InspectStats(input[1]); err != nil {
		return err
	}
	return nil
}

func commandMap(cfg *config, input []string) error {

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

func commandMapB(cfg *config, input []string) error {
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

func commandPokedex(cfg *config, input []string) error {
	dex.List()
	return nil
}
