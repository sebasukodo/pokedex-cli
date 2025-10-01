package pokeapi

import (
	"fmt"
	"strings"
)

type Pokedex struct {
	Pokemon map[string]*Pokemon
}

func (dex *Pokedex) Add(poke Pokemon) error {
	value, ok := dex.Pokemon[poke.Name]
	if ok {
		value.timesCaught++
		dex.Pokemon[poke.Name] = value
		return nil
	}

	poke.timesCaught = 1
	dex.Pokemon[poke.Name] = &poke
	return nil
}

func (dex *Pokedex) InspectStats(name string) error {
	poke, ok := dex.Pokemon[name]
	if !ok {
		return fmt.Errorf("you have not caught that pokemon")
	}

	text := fmt.Sprintf("Name: %v\nHeight: %v\nWeight: %v\nStats:\n", poke.Name, poke.Height, poke.Weight)

	for _, value := range poke.Stats {
		text += fmt.Sprintf("  - %v: %v\n", value.Stat.Name, value.BaseStat)
	}

	text += "Types:\n"

	for _, value := range poke.Types {
		text += fmt.Sprintf("  - %v\n", value.Type.Name)
	}

	text = strings.TrimSuffix(text, "\n")
	fmt.Println(text)
	return nil

}

func (dex *Pokedex) List() {
	text := "Listing all your Pokemons...\n"

	if len(dex.Pokemon) == 0 {
		text += "Empty Pokedex\n"
	} else {
		for _, poke := range dex.Pokemon {
			if poke.timesCaught > 1 {
				text += fmt.Sprintf(" - %v: %v times\n", poke.Name, poke.timesCaught)
			} else {
				text += fmt.Sprintf(" - %v\n", poke.Name)
			}
		}

	}

	text = strings.TrimSuffix(text, "\n")
	fmt.Println(text)

}

func NewPokedex() Pokedex {
	return Pokedex{
		Pokemon: make(map[string]*Pokemon),
	}
}
