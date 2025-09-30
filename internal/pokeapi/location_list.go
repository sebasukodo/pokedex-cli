package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/sebasukodo/pokedex-cli/internal/pokecache"
)

var cache *pokecache.Cache

type ShallowLocations struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func init() {
	cache = pokecache.NewCache(5 * time.Minute)
}

func ListLocations(pageURL *string) (ShallowLocations, error) {
	url := baseURL + "location-area/"

	if pageURL != nil {
		url = *pageURL
	}

	fmt.Println("Checking cache...")

	cachedData, ok := cache.Get(url)
	if !ok {
		fmt.Println("Data not found in cache, GET-call...")
		res, err := http.Get(url)
		if err != nil {
			return ShallowLocations{}, err
		}

		read, err := io.ReadAll(res.Body)
		if err != nil {
			return ShallowLocations{}, err
		}
		if res.StatusCode > 299 {
			log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, read)
		}
		defer res.Body.Close()

		cache.Add(url, read)

		cachedData = read

	} else {
		fmt.Println("Data found in cache...")
	}

	var data ShallowLocations
	if err := json.Unmarshal(cachedData, &data); err != nil {
		return ShallowLocations{}, err
	}

	return data, nil

}
