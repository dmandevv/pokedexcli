package main

import (
	"fmt"
)

func commandExplore(cfg *config, area_name string) error {

	area, err := cfg.pokeapiClient.GetAreaPokemonList(area_name)
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
