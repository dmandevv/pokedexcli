package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) == 0 || args[0] == "" {
		return errors.New("please provide a pokemon's name. Usage: catch <pokemon-name>")
	}

	pokemon, err := cfg.pokeapiClient.GetPokemonType(args[0])
	if err != nil {
		return err
	}

	species, err := cfg.pokeapiClient.GetPokemonSpecies(args[0])
	if err != nil {
		return err
	}

	println("Throwing a Pokeball at", pokemon.Name+"...")

	chance := rand.Intn(256)
	if chance > species.CaptureRate {
		fmt.Println(pokemon.Name, "escaped!")
		return nil
	}
	fmt.Println(pokemon.Name, "was caught!")

	cfg.pokedex[pokemon.Name] = pokemon
	return nil
}
