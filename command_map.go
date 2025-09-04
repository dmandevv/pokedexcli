package main

import (
	"errors"
	"fmt"

	"github.com/dmandevv/pokedexcli/internal/pokecache"
)

func commandMapf(cfg *config, cache *pokecache.Cache) error {

	locations, err := cfg.pokeapiClient.GetLocationsList(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locations.Next
	cfg.previousLocationsURL = locations.Previous

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandMapb(cfg *config, cache *pokecache.Cache) error {
	if cfg.previousLocationsURL == nil {
		return errors.New("you're on the first page")
	}

	locations, err := cfg.pokeapiClient.GetLocationsList(cfg.previousLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locations.Next
	cfg.previousLocationsURL = locations.Previous

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	return nil
}
