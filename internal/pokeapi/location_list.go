package pokeapi

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/sebasukodo/pokedex-cli/internal/pokecache"
)

var cache *pokecache.Cache

func init() {
	cache = pokecache.NewCache(5 * time.Minute)
}

func ExploreLocation(id string) (SpecificLocations, error) {
	url := baseURL + "location-area/" + id

	cachedData, ok := cache.Get(url)
	if !ok {
		res, err := http.Get(url)
		if err != nil {
			return SpecificLocations{}, err
		}

		read, err := io.ReadAll(res.Body)
		if err != nil {
			return SpecificLocations{}, err
		}
		if res.StatusCode > 299 {
			log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, read)
		}
		defer res.Body.Close()

		cache.Add(url, read)

		cachedData = read

	}

	var data SpecificLocations
	if err := json.Unmarshal(cachedData, &data); err != nil {
		return SpecificLocations{}, err
	}

	return data, nil
}

func ListLocations(pageURL *string) (ShallowLocations, error) {
	url := baseURL + "location-area"

	if pageURL != nil {
		url = *pageURL
	}

	cachedData, ok := cache.Get(url)
	if !ok {

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

	}

	var data ShallowLocations
	if err := json.Unmarshal(cachedData, &data); err != nil {
		return ShallowLocations{}, err
	}

	return data, nil
}
