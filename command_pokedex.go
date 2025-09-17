package main

import (
	"errors"
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {

	if len(cfg.pokedex) == 0 {
		return errors.New("your pokedex is empty. catch some pokemon first")
	}

	fmt.Println("Your Pokedex:")
	for name := range cfg.pokedex {
		fmt.Println(" - " + name)
	}

	return nil
}
