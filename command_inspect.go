package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) == 0 || args[0] == "" {
		return errors.New("please provide a pokemon's name. Usage: inspect <pokemon-name>")
	}

	pokemon, found := cfg.pokedex[args[0]]
	if !found {
		return errors.New("you have not caught that pokemon")
	}

	println("Name:", pokemon.Name)
	println("Height:", pokemon.Height)
	println("Weight:", pokemon.Weight)
	println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("- %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf("- %s\n", t.Type.Name)
	}

	return nil
}
