package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {

	if len(args) == 0 || args[0] == "" {
		return errors.New("please provide an area name. Usage: explore <area-name>")
	}

	area, err := cfg.pokeapiClient.GetAreaPokemonList(args[0])
	if err != nil {
		return err
	}

	fmt.Println("Exploring", area.Name+"...")
	if len(area.Encounters) == 0 {
		fmt.Println("No Pokémon found in this area.")
		return nil
	}

	fmt.Println("Found Pokémon:")
	for _, encounter := range area.Encounters {
		fmt.Println("- " + encounter.Pokemon.Name)
	}

	return nil
}
