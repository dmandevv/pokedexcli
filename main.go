package main

import (
	"time"

	"github.com/dmandevv/pokedexcli/internal/pokeapi"
)

func main() {
	pokeclient := pokeapi.NewClient(time.Second*5, time.Minute*5)
	cfg := &config{
		pokeapiClient: pokeclient,
	}
	startRepl(cfg)
}
