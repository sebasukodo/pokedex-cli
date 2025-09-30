package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

const (
	baseURL = "https://pokeapi.co/api/v2/location-area/"
)

func getAPIInfo(id int) MapData {
	fullUrl := fmt.Sprintf("%v%v/", baseURL, id)

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

type MapData struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Location struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
}
