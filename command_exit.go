package main

import (
	"fmt"
	"os"

	"github.com/dmandevv/pokedexcli/internal/pokecache"
)

func commandExit(cfg *config, cache *pokecache.Cache) error {
	fmt.Print("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}
