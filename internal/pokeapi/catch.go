package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"math/rand/v2"
	"net/http"
)

func getPokemonInfo(name string) (Pokemon, error) {
	if name == "" {
		return Pokemon{}, fmt.Errorf("cannot catch %v, not a pokemon", name)
	}

	url := baseURL + "pokemon/" + name

	cachedData, ok := cache.Get(url)
	if !ok {

		res, err := http.Get(url)
		if err != nil {
			return Pokemon{}, err
		}
		defer res.Body.Close()

		data, err := io.ReadAll(res.Body)
		if err != nil {
			return Pokemon{}, err
		}

		if res.StatusCode > 299 {
			if res.StatusCode == 404 {
				return Pokemon{}, fmt.Errorf("%v is not a valid pokemon", name)
			}
			return Pokemon{}, fmt.Errorf("response failed with status code: %d and\nbody: %s", res.StatusCode, data)
		}

		cache.Add(url, data)

		cachedData = data

	}

	var poke Pokemon
	if err := json.Unmarshal(cachedData, &poke); err != nil {
		return Pokemon{}, err
	}

	return poke, nil
}

func CatchPokemon(name string) (Pokemon, bool, error) {

	pokemon, err := getPokemonInfo(name)
	if err != nil {
		return Pokemon{}, false, err
	}

	baseChance := 75.0 //the higher the better chances low exp pokemon have
	slope := -0.004    //the lower the better chances high exp pokemon have
	catchChance := (baseChance * math.Exp(slope*float64(pokemon.BaseExperience))) / 100
	random := rand.Float64()

	caught := random < catchChance

	return pokemon, caught, nil
}
