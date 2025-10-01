package pokeapi

import (
	"time"

	"github.com/sebasukodo/pokedex-cli/internal/pokecache"
)

const (
	baseURL = "https://pokeapi.co/api/v2/"
)

var cache *pokecache.Cache

func init() {
	cache = pokecache.NewCache(5 * time.Minute)
}
