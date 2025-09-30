package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type MapData struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Location struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
}

func ListLocations(id int) MapData {
	url := baseURL + "location-area/"
	fullUrl := fmt.Sprintf("%v%v/", url, id)

	res, err := http.Get(fullUrl)
	if err != nil {
		log.Fatal(err)
	}

	read, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, read)
	}

	defer res.Body.Close()

	var data MapData
	if err := json.Unmarshal(read, &data); err != nil {
		log.Fatal(err)
	}

	return data

}
