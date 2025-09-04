package main

import (
	"time"

	"github.com/dmandevv/pokedexcli/internal/pokeapi"
	"github.com/dmandevv/pokedexcli/internal/pokecache"
)

func main() {
	pokeclient := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		pokeapiClient: pokeclient,
	}
	cache := pokecache.NewCache(time.Second * 5)
	startRepl(cfg, cache)
}
