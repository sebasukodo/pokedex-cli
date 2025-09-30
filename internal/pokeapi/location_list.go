package pokeapi

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type ShallowLocations struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func ListLocations(pageURL *string) (ShallowLocations, error) {
	url := baseURL + "location-area/"

	if pageURL != nil {
		url = *pageURL
	}

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

	var data ShallowLocations
	if err := json.Unmarshal(read, &data); err != nil {
		return ShallowLocations{}, err
	}

	return data, nil

}
